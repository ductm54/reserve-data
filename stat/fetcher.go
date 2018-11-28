package stat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/settings"
	statutil "github.com/KyberNetwork/reserve-data/stat/util"
	ethereum "github.com/ethereum/go-ethereum/common"
)

const (
	reorgBlockSafe       uint64 = 7
	TimezoneBucketPrefix string = "utc"
	StartTimezone        int64  = -11
	EndTimezone          int64  = 14
	blockRange           uint64 = 200
	success              string = "OK"
	noTxsFound           string = "No transactions found"

	ethDecimals int64  = 18
	ethAddress  string = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"

	TradeSummaryAggregation string = "trade_summary_aggregation"
	WalletAggregation       string = "wallet_aggregation"
	CountryAggregation      string = "country_aggregation"
	VolumeStatAggregation   string = "volume_stat_aggregation"
	BurnfeeAggregation      string = "burn_fee_aggregation"
	UserInfoAggregation     string = "user_info_aggregation"
	TradeSummaryKey         string = "trade_summary"

	etherScanAPIEndpoint      = "http://api.etherscan.io/api"
	broadcastKyberAPIEndpoint = "https://broadcast.kyber.network"
)

type Fetcher struct {
	statStorage            StatStorage
	userStorage            UserStorage
	logStorage             LogStorage
	rateStorage            RateStorage
	feeSetRateStorage      FeeSetRateStorage
	blockchain             Blockchain
	runner                 FetcherRunner
	currentBlock           uint64
	currentBlockUpdateTime uint64
	deployBlock            uint64
	apiKey                 string
	sleepTime              time.Duration
	blockNumMarker         uint64
	setting                Setting
	ipLocator              *statutil.IPLocator
	addressLookup          map[ethereum.Address]common.Token
	mu                     sync.RWMutex
}

func NewFetcher(
	statStorage StatStorage,
	logStorage LogStorage,
	rateStorage RateStorage,
	userStorage UserStorage,
	feeSetRateStorage FeeSetRateStorage,
	runner FetcherRunner,
	deployBlock uint64,
	beginBlockSetRate uint64,
	apiKey string,
	setting Setting,
	iploc *statutil.IPLocator) *Fetcher {
	sleepTime := time.Second
	fetcher := &Fetcher{
		statStorage:       statStorage,
		logStorage:        logStorage,
		rateStorage:       rateStorage,
		userStorage:       userStorage,
		feeSetRateStorage: feeSetRateStorage,
		blockchain:        nil,
		runner:            runner,
		deployBlock:       deployBlock,
		apiKey:            apiKey,
		sleepTime:         sleepTime,
		setting:           setting,
		ipLocator:         iploc,
		addressLookup:     make(map[ethereum.Address]common.Token),
	}
	lastBlockChecked, err := fetcher.feeSetRateStorage.GetLastBlockChecked()
	if err != nil {
		log.Printf("can't get last block checked from db: %s", err)
		panic(err)
	}
	if lastBlockChecked == 0 {
		fetcher.blockNumMarker = beginBlockSetRate
	} else {
		fetcher.blockNumMarker = lastBlockChecked + 1
	}
	return fetcher
}

func (self *Fetcher) Stop() error {
	return self.runner.Stop()
}

func (self *Fetcher) SetBlockchain(blockchain Blockchain) {
	self.blockchain = blockchain
	self.FetchCurrentBlock()
}

func (self *Fetcher) WaitForCoreAndRun() {
	const waitTime = 5 * time.Second
	//wait till core is ready to serve
	for {
		err := self.setting.ReadyToServe()
		log.Printf("STAT: waiting for core.... try to ping core, got err %v. ", err)
		if err == nil {
			break
		}
		time.Sleep(waitTime)
	}

	go self.RunBlockFetcher()
	go self.RunLogFetcher()
	go self.RunReserveRatesFetcher()
	go self.RunTradeLogProcessor()
	go self.RunCatLogProcessor()
	go self.RunFeeSetrateFetcher()
}

func (self *Fetcher) Run() error {
	log.Printf("Fetcher runner is starting...")
	if err := self.runner.Start(); err != nil {
		return err
	}
	go self.WaitForCoreAndRun()

	log.Printf("Fetcher runner is running...")
	return nil
}

func (self *Fetcher) RunFeeSetrateFetcher() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	for {
		err := self.FetchTxs(client)
		if err != nil {
			log.Printf("failed to fetch data from etherescan: %s", err)
		}
		time.Sleep(self.sleepTime)
	}
}

type APIResponse struct {
	Message string                 `json:"message"`
	Result  []common.SetRateTxInfo `json:"result"`
}

func (self *Fetcher) FetchTxs(client http.Client) error {
	fromBlock := self.blockNumMarker
	toBlock := self.GetToBlock()
	if toBlock == 0 {
		return errors.New("Can't get latest block nummber")
	}
	pricingAddress, err := self.blockchain.GetAddress(settings.Pricing)
	if err != nil {
		return err
	}
	api := fmt.Sprintf("%s?module=account&action=txlist&address=%s&startblock=%d&endblock=%d&apikey=%s", etherScanAPIEndpoint, pricingAddress.String(), fromBlock, toBlock, self.apiKey)
	log.Println("api get txs of setrate: ", api)
	resp, err := client.Get(api)
	if err != nil {
		return err
	}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			log.Printf("cannot close response body: %s", cErr.Error())
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	apiResponse := APIResponse{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Printf("can't unmarshal data from etherscan: %s", err)
		return err
	}

	if apiResponse.Message == success || apiResponse.Message == noTxsFound {
		sameBlockBucket := []common.SetRateTxInfo{}
		setRateTxsInfo := apiResponse.Result
		numberEle := len(setRateTxsInfo)
		var blockNumber string
		for index, transaction := range setRateTxsInfo {
			if self.isPricingMethod(transaction.Input) {
				blockNumber = transaction.BlockNumber
				sameBlockBucket = append(sameBlockBucket, transaction)
				if index < numberEle-1 && setRateTxsInfo[index+1].BlockNumber == blockNumber {
					continue
				}
				err = self.feeSetRateStorage.StoreTransaction(sameBlockBucket)
				if err != nil {
					log.Printf("failed to store pricing's txs: %s", err)
					return err
				}
				sameBlockBucket = []common.SetRateTxInfo{}
			}
		}
		log.Println("fetch and store pricing's txs done!")
		if toBlock == self.currentBlock {
			self.blockNumMarker = toBlock
		} else {
			self.blockNumMarker = toBlock + 1
		}
	}
	return nil
}

func (self *Fetcher) isPricingMethod(inputData string) bool {
	if inputData == "0x" {
		return false
	}
	method, err := self.blockchain.GetPricingMethod(inputData)
	if err != nil {
		log.Printf("Cannot find method from input data: %v", err)
		return false
	}
	methodName := method.Name
	if methodName == "setCompactData" || methodName == "setBaseRate" {
		return true
	}
	return false
}

func (self *Fetcher) GetToBlock() uint64 {
	currentBlock := self.currentBlock
	blockNumMarker := self.blockNumMarker
	if currentBlock == 0 {
		return 0
	}
	if currentBlock <= blockNumMarker+blockRange {
		self.sleepTime = 5 * time.Minute
		return currentBlock
	}
	toBlock := blockNumMarker + blockRange
	self.sleepTime = time.Second
	return toBlock
}

func (self *Fetcher) RunCatLogProcessor() {
	for {
		t := <-self.runner.GetCatLogProcessorTicker()
		// get trade log from db
		fromTime, err := self.userStorage.GetLastProcessedCatLogTimepoint()
		if err != nil {
			log.Printf("get last processor state from db failed: %v", err)
			continue
		}
		fromTime++
		if fromTime == 1 {
			// there is no cat log being processed before
			// load the first log we have and set the fromTime to it's timestamp
			var l common.SetCatLog
			l, err = self.logStorage.GetFirstCatLog()
			if err != nil {
				log.Printf("can't get first cat log: err(%s)", err)
				continue
			} else {
				fromTime = l.Timestamp - 1
			}
		}
		toTime := common.TimeToTimepoint(t) * 1000000
		maxRange := self.logStorage.MaxRange()
		if toTime-fromTime > maxRange {
			toTime = fromTime + maxRange
		}
		catLogs, err := self.logStorage.GetCatLogs(fromTime, toTime)
		if err != nil {
			log.Printf("get cat log from db failed: %v", err)
			continue
		}
		log.Printf("PROCESS %d cat logs from %d to %d", len(catLogs), fromTime, toTime)
		if len(catLogs) > 0 {
			var last uint64
			for _, l := range catLogs {
				err := self.userStorage.UpdateAddressCategory(
					l.Address,
					l.Category,
				)
				if err != nil {
					log.Printf("updating address and category failed: err(%s)", err)
				} else {
					if l.Timestamp > last {
						last = l.Timestamp
					}
				}
			}
			if err := self.userStorage.SetLastProcessedCatLogTimepoint(last); err != nil {
				log.Printf("Set last process cat log timepoint error: %s", err.Error())
			}
		} else {
			l, err := self.logStorage.GetLastCatLog()
			if err != nil {
				log.Printf("LogFetcher - can't get last cat log: err(%s)", err)
			} else {
				// log.Printf("LogFetcher - got last cat log: %+v", l)
				if toTime < l.Timestamp {
					// if we are querying on past logs, store toTime as the last
					// processed trade log timepoint
					if err := self.userStorage.SetLastProcessedCatLogTimepoint(toTime); err != nil {
						log.Printf("Set last process cat log timepoint error: %s", err.Error())
					}
				}
			}
		}

		log.Println("processed cat logs")
	}
}

func (self *Fetcher) GetTradeLogTimeRange(fromTime uint64, t time.Time) (uint64, uint64) {
	fromTime++
	if fromTime == 1 {
		// there is no trade log being processed before
		// load the first log we have and set the fromTime to it's timestamp
		l, err := self.logStorage.GetFirstTradeLog()
		if err != nil {
			log.Printf("can't get first trade log: err(%s)", err)
			// continue
		} else {
			log.Printf("got first trade: %+v", l)
			fromTime = l.Timestamp - 1
		}
	}
	toTime := common.TimeToTimepoint(t) * 1000000
	maxRange := self.logStorage.MaxRange()
	if toTime-fromTime > maxRange {
		toTime = fromTime + maxRange
	}
	return fromTime, toTime
}

func (self *Fetcher) RunCountryStatAggregation(t time.Time) {
	// get trade log from db
	fromTime, err := self.statStorage.GetLastProcessedTradeLogTimepoint(CountryAggregation)
	if err != nil {
		log.Printf("get trade log processor state from db failed: %v", err)
		return
	}
	fromTime, toTime := self.GetTradeLogTimeRange(fromTime, t)
	tradeLogs, err := self.logStorage.GetTradeLogs(fromTime, toTime)
	if err != nil {
		log.Printf("get trade log from db failed: %v", err)
	}
	if len(tradeLogs) > 0 {
		if err := self.statStorage.SetFirstTradeEver(&tradeLogs); err != nil {
			log.Printf("Set first trade ever error: %s", err.Error())
		}
		if err := self.statStorage.SetFirstTradeInDay(&tradeLogs); err != nil {
			log.Printf("Set first trade ever error: %s", err.Error())
		}
		var last uint64
		countryStats := map[string]common.MetricStatsTimeZone{}
		allFirstTradeEver, _ := self.statStorage.GetAllFirstTradeEver()
		kycEdUsers, _ := self.userStorage.GetKycUsers()
		for _, trade := range tradeLogs {
			if err := self.aggregateCountryStats(trade, countryStats, allFirstTradeEver, kycEdUsers); err != nil {
				log.Printf("STAT: aggregate country stat got err %s, return now and wait till next ticker", err)
				return
			}
			if trade.Timestamp > last {
				last = trade.Timestamp
			}
		}
		if err := self.statStorage.SetCountryStat(countryStats, last); err != nil {
			log.Printf("Set country stat error: %s", err.Error())
			return
		}
	} else {
		l, err := self.logStorage.GetLastTradeLog()
		if err != nil {
			log.Printf("can't get last trade log: err(%s)", err)
			return
		}
		if toTime < l.Timestamp {
			// if we are querying on past logs, store toTime as the last
			// processed trade log timepoint
			if err := self.statStorage.SetLastProcessedTradeLogTimepoint(CountryAggregation, toTime); err != nil {
				log.Printf("Set last processed tradelog timepoint error: %s", err.Error())
			}
		}
	}
}

func (self *Fetcher) RunTradeSummaryAggregation(t time.Time) {
	// get trade log from db
	fromTime, err := self.statStorage.GetLastProcessedTradeLogTimepoint(TradeSummaryAggregation)
	if err != nil {
		log.Printf("get trade log processor state from db failed: %v", err)
		return
	}
	fromTime, toTime := self.GetTradeLogTimeRange(fromTime, t)
	tradeLogs, err := self.logStorage.GetTradeLogs(fromTime, toTime)
	if err != nil {
		log.Printf("get trade log from db failed: %v", err)
		return
	}
	if len(tradeLogs) > 0 {
		log.Printf("STAT: Trade summary aggregation got %d logs from %d to %d", len(tradeLogs), fromTime, toTime)
		if err := self.statStorage.SetFirstTradeEver(&tradeLogs); err != nil {
			log.Printf("Set first trade ever error: %s", err.Error())
		}
		if err := self.statStorage.SetFirstTradeInDay(&tradeLogs); err != nil {
			log.Printf("Set first trade in day: %s", err.Error())
		}
		var last uint64

		tradeSummary := map[string]common.MetricStatsTimeZone{}
		allFirstTradeEver, _ := self.statStorage.GetAllFirstTradeEver()
		kycEdUsers, _ := self.userStorage.GetKycUsers()
		for _, trade := range tradeLogs {
			if err := self.aggregateTradeSumary(trade, tradeSummary, allFirstTradeEver, kycEdUsers); err != nil {
				log.Printf("STAT: aggregate trade summary got err %s, return now and wait till next ticker", err)
				return
			}
			if trade.Timestamp > last {
				last = trade.Timestamp
			}
		}
		if err := self.statStorage.SetTradeSummary(tradeSummary, last); err != nil {
			log.Printf("Set trade summary error: %s", err.Error())
			return
		}
		// self.statStorage.SetLastProcessedTradeLogTimepoint(tradeSummaryAggregation, last)
	} else {
		l, err := self.logStorage.GetLastTradeLog()
		if err != nil {
			log.Printf("can't get last trade log: err(%s)", err)
			return
		}
		if toTime < l.Timestamp {
			// if we are querying on past logs, store toTime as the last
			// processed trade log timepoint
			if err := self.statStorage.SetLastProcessedTradeLogTimepoint(TradeSummaryAggregation, toTime); err != nil {
				log.Printf("Set last processed tradelog timepoint error: %s", err.Error())
			}
		}
	}
}

func (self *Fetcher) RunWalletStatAggregation(t time.Time) {
	// get trade log from db
	fromTime, err := self.statStorage.GetLastProcessedTradeLogTimepoint(WalletAggregation)
	if err != nil {
		log.Printf("get trade log processor state from db failed: %v", err)
		return
	}
	fromTime, toTime := self.GetTradeLogTimeRange(fromTime, t)
	tradeLogs, err := self.logStorage.GetTradeLogs(fromTime, toTime)
	if err != nil {
		log.Printf("get trade log from db failed: %v", err)
		return
	}
	if len(tradeLogs) > 0 {
		if err := self.statStorage.SetFirstTradeEver(&tradeLogs); err != nil {
			log.Printf("Set first trade ever error: %s", err.Error())
		}
		if err := self.statStorage.SetFirstTradeInDay(&tradeLogs); err != nil {
			log.Printf("Set first trade in day error: %s", err.Error())
		}
		var last uint64

		walletStats := map[string]common.MetricStatsTimeZone{}
		allFirstTradeEver, _ := self.statStorage.GetAllFirstTradeEver()
		kycEdUsers, _ := self.userStorage.GetKycUsers()
		for _, trade := range tradeLogs {
			if err := self.aggregateWalletStats(trade, walletStats, allFirstTradeEver, kycEdUsers); err != nil {
				log.Printf("STAT: aggregate wallet stat got err %s, return now and wait till next ticker", err)
				return
			}
			if trade.Timestamp > last {
				last = trade.Timestamp
			}
		}
		if err := self.statStorage.SetWalletStat(walletStats, last); err != nil {
			log.Printf("Set wallet stats error: %s", err.Error())
			return
		}
		// self.statStorage.SetLastProcessedTradeLogTimepoint(walletAggregation, last)
	} else {
		l, err := self.logStorage.GetLastTradeLog()
		if err != nil {
			log.Printf("can't get last trade log: err(%s)", err)
			return
		}
		if toTime < l.Timestamp {
			// if we are querying on past logs, store toTime as the last
			// processed trade log timepoint
			if err := self.statStorage.SetLastProcessedTradeLogTimepoint(WalletAggregation, toTime); err != nil {
				log.Printf("Set last processed tradelog timepoint: %s", err.Error())
			}
		}
	}
}

func (self *Fetcher) RunBurnFeeAggregation(t time.Time) {
	// get trade log from db
	fromTime, err := self.statStorage.GetLastProcessedTradeLogTimepoint(BurnfeeAggregation)
	if err != nil {
		log.Printf("get trade log processor state from db failed: %v", err)
		return
	}
	fromTime, toTime := self.GetTradeLogTimeRange(fromTime, t)
	tradeLogs, err := self.logStorage.GetTradeLogs(fromTime, toTime)
	if err != nil {
		log.Printf("get trade log from db failed: %v", err)
		return
	}
	if len(tradeLogs) > 0 {
		var last uint64

		burnFeeStats := map[string]common.BurnFeeStatsTimeZone{}
		for _, trade := range tradeLogs {
			if err := self.aggregateBurnFeeStats(trade, burnFeeStats); err != nil {
				log.Printf("STAT: aggregate burnFee stat got err %s, return now and wait till next ticker", err)
				return
			}
			if trade.Timestamp > last {
				last = trade.Timestamp
			}
		}
		if err := self.statStorage.SetBurnFeeStat(burnFeeStats, last); err != nil {
			log.Printf("Set burn fee error: %s", err.Error())
			return
		}
		// self.statStorage.SetLastProcessedTradeLogTimepoint(burnfeeAggregation, last)
	} else {
		l, err := self.logStorage.GetLastTradeLog()
		if err != nil {
			log.Printf("can't get last trade log: err(%s)", err)
			return
		}
		if toTime < l.Timestamp {
			if err := self.statStorage.SetLastProcessedTradeLogTimepoint(BurnfeeAggregation, toTime); err != nil {
				log.Printf("Set last processed tradelog timepoint error: %s", err.Error())
			}
		}
	}
}

func (self *Fetcher) RunVolumeStatAggregation(t time.Time) {
	// get trade log from db
	fromTime, err := self.statStorage.GetLastProcessedTradeLogTimepoint(VolumeStatAggregation)
	if err != nil {
		log.Printf("get trade log processor state from db failed: %v", err)
		return
	}
	fromTime, toTime := self.GetTradeLogTimeRange(fromTime, t)
	tradeLogs, err := self.logStorage.GetTradeLogs(fromTime, toTime)
	if err != nil {
		log.Printf("get trade log from db failed: %v", err)
		return
	}
	if len(tradeLogs) > 0 {
		var last uint64

		volumeStats := map[string]common.VolumeStatsTimeZone{}
		for _, trade := range tradeLogs {
			if err := self.aggregateVolumeStats(trade, volumeStats); err != nil {
				log.Printf("STAT: aggregate Volume stat got err %s, return now and wait till next ticker", err)
				return
			}
			if trade.Timestamp > last {
				last = trade.Timestamp
			}
		}
		if err := self.statStorage.SetVolumeStat(volumeStats, last); err != nil {
			log.Printf("Set volume stat error: %s", err.Error())
			return
		}
	} else {
		l, err := self.logStorage.GetLastTradeLog()
		if err != nil {
			log.Printf("can't get last trade log: err(%s)", err)
			return
		}
		if toTime < l.Timestamp {
			if err := self.statStorage.SetLastProcessedTradeLogTimepoint(VolumeStatAggregation, toTime); err != nil {
				log.Printf("Set last processed tradelog timepoint error: %s", err.Error())
			}
		}
	}
	return
}

func (self *Fetcher) RunUserInfoAggregation(t time.Time) {
	// get trade log from db
	fromTime, err := self.statStorage.GetLastProcessedTradeLogTimepoint(UserInfoAggregation)
	if err != nil {
		log.Printf("get trade log processor state from db failed: %v", err)
		return
	}
	fromTime, toTime := self.GetTradeLogTimeRange(fromTime, t)
	tradeLogs, err := self.logStorage.GetTradeLogs(fromTime, toTime)
	if err != nil {
		log.Printf("get trade log from db failed: %v", err)
		return
	}
	if len(tradeLogs) > 0 {
		var last uint64
		userInfos := map[string]common.UserInfoTimezone{}
		for _, trade := range tradeLogs {
			if err := self.aggregateUserInfo(trade, userInfos); err != nil {
				log.Printf("STAT: aggregate user info got err %s, return now and wait till next ticker", err)
				return
			}
			if trade.Timestamp > last {
				last = trade.Timestamp
			}
		}
		if err := self.statStorage.SetUserList(userInfos, last); err != nil {
			log.Printf("Set user list: %s", err.Error())
			return
		}
	} else {
		l, err := self.logStorage.GetLastTradeLog()
		if err != nil {
			log.Printf("can't get last trade log: err(%s)", err)
			return
		}
		if toTime < l.Timestamp {
			if err := self.statStorage.SetLastProcessedTradeLogTimepoint(UserInfoAggregation, toTime); err != nil {
				log.Printf("Set last processed tradelog timepoint error: %s", err.Error())
			}
		}
	}
}

func runAggregationInParallel(wg *sync.WaitGroup, t time.Time, f func(t time.Time)) {
	defer wg.Done()
	f(t)
}

func (self *Fetcher) RunTradeLogProcessor() {
	for {
		t := <-self.runner.GetTradeLogProcessorTicker()
		// self.RunUserAggregation(t)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go runAggregationInParallel(&wg, t, self.RunBurnFeeAggregation)
		wg.Add(1)
		go runAggregationInParallel(&wg, t, self.RunVolumeStatAggregation)
		wg.Add(1)
		go runAggregationInParallel(&wg, t, self.RunTradeSummaryAggregation)
		wg.Add(1)
		go runAggregationInParallel(&wg, t, self.RunWalletStatAggregation)
		wg.Add(1)
		go runAggregationInParallel(&wg, t, self.RunCountryStatAggregation)
		wg.Add(1)
		go runAggregationInParallel(&wg, t, self.RunUserInfoAggregation)
		wg.Wait()
	}
}

func (self *Fetcher) RunReserveRatesFetcher() {
	for {
		log.Printf("waiting for signal from reserve rate channel")
		t := <-self.runner.GetReserveRatesTicker()
		log.Printf("got signal in reserve rate channel with timstamp %d", common.GetTimepoint())
		timepoint := common.TimeToTimepoint(t)
		self.FetchReserveRates(timepoint)
		log.Printf("fetched reserve rate from blockchain")
	}
}

func (self *Fetcher) GetReserveRates(
	currentBlock uint64, reserveAddr ethereum.Address,
	tokens []common.Token, data *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	rates, err := self.blockchain.GetReserveRates(currentBlock-1, currentBlock, reserveAddr, tokens)
	if err != nil {
		log.Println(err.Error())
	}
	data.Store(reserveAddr, rates)
}

func (self *Fetcher) ReserveSupportedTokens(reserve ethereum.Address) ([]common.Token, error) {
	tokens := []common.Token{}
	reserveAddr, err := self.blockchain.GetAddress(settings.Reserve)
	if err != nil {
		return tokens, err
	}
	if reserve == reserveAddr {
		internalTokens, err := self.setting.GetInternalTokens()
		if err != nil {
			log.Printf("ERROR: Can not get internal tokens: %s", err)
			return tokens, err
		}
		for _, token := range internalTokens {
			if token.ID != "ETH" {
				tokens = append(tokens, token)
			}
		}
	} else {
		activeTokens, err := self.setting.GetActiveTokens()
		if err != nil {
			log.Printf("ERROR: Can not get internal tokens: %s", err)
			return tokens, err
		}
		for _, token := range activeTokens {
			if token.ID != "ETH" {
				tokens = append(tokens, token)
			}
		}
	}
	return tokens, nil
}

func (self *Fetcher) FetchReserveRates(timepoint uint64) {
	log.Printf("Fetching reserve and sanity rate from blockchain")
	thirdPartyReserves, err := self.blockchain.GetAddresses(settings.ThirdPartyReserves)
	if err != nil {
		log.Printf("ERROR: Can not get reserve rates %s", err)
		return
	}
	reserveAddr, err := self.blockchain.GetAddress(settings.Reserve)
	if err != nil {
		log.Printf("ERROR: Can not get reserve rates %s", err)
		return
	}
	supportedReserves := append(thirdPartyReserves, reserveAddr)
	data := sync.Map{}
	wg := sync.WaitGroup{}
	// get current block to use to fetch all reserve rates.
	// dont use self.currentBlock directly with self.GetReserveRates
	// because otherwise, rates from different reserves will not
	// be synced with block no
	block := self.currentBlock
	for _, reserve := range supportedReserves {
		tokens, err := self.ReserveSupportedTokens(reserve)
		if err == nil {
			wg.Add(1)
			go self.GetReserveRates(block, reserve, tokens, &data, &wg)
		} else {
			log.Printf("ERROR: Can not get reserve rates %s", err)
		}
	}
	wg.Wait()
	data.Range(func(key, value interface{}) bool {
		reserveAddr, ok := key.(ethereum.Address)
		//if there is conversion error, continue to next key,val
		if !ok {
			log.Printf("key (%v) cannot be asserted to ethereum.Address", key)
			return true
		}
		rates, ok := value.(common.ReserveRates)
		if !ok {
			log.Printf("valuve (%v) cannot be asserted to reserveRates", value)
			return true
		}
		log.Printf("Storing reserve rates to db...")
		if err := self.rateStorage.StoreReserveRates(reserveAddr, rates, common.GetTimepoint()); err != nil {
			log.Printf("Store reserve rates error: %s", err.Error())
		}
		return true
	})
}

func (self *Fetcher) RunLogFetcher() {
	for {
		log.Printf("LogFetcher - waiting for signal from log channel")
		t := <-self.runner.GetLogTicker()
		timepoint := common.TimeToTimepoint(t)
		log.Printf("LogFetcher - got signal in log channel with timestamp %d", timepoint)
		lastBlock, err := self.logStorage.LastBlock()
		if lastBlock == 0 {
			lastBlock = self.deployBlock
		}
		if err == nil {
			toBlock := lastBlock + 1 + 1440 // 1440 is considered as 6 hours
			if toBlock > self.currentBlock-reorgBlockSafe {
				toBlock = self.currentBlock - reorgBlockSafe
			}
			if lastBlock+1 > toBlock {
				continue
			}
			nextBlock, fErr := self.FetchLogs(lastBlock+1, toBlock, timepoint)
			if fErr != nil {
				// in case there is error, we roll back and try it again.
				// dont have to do anything here. just continute with the loop.
				log.Printf("LogFetcher - continue with the loop to try it again: %s", fErr)
			} else {
				if nextBlock == lastBlock && toBlock != 0 {
					// in case that we are querying old blocks (6 hours in the past)
					// and got no logs. we will still continue with next block
					// It is not the case if toBlock == 0, means we are querying
					// best window, we should keep querying it in order not to
					// miss any logs due to node inconsistency
					nextBlock = toBlock
				}
				log.Printf("LogFetcher - update log block: %d", nextBlock)
				if err = self.logStorage.UpdateLogBlock(nextBlock, timepoint); err != nil {
					log.Printf("Update log block: %s", err.Error())
				}
			}
		} else {
			log.Printf("LogFetcher - failed to get last fetched log block, err: %+v", err)
		}
	}
}

func (self *Fetcher) RunBlockFetcher() {
	for {
		log.Printf("waiting for signal from block channel")
		t := <-self.runner.GetBlockTicker()
		timepoint := common.TimeToTimepoint(t)
		log.Printf("got signal in block channel with timestamp %d", timepoint)
		self.FetchCurrentBlock()
		log.Printf("fetched block from blockchain")
	}
}

//GetTradeGeo get geo from trade log
func GetTradeGeo(ipLocator *statutil.IPLocator, txHash string) (string, string, error) {
	url := fmt.Sprintf("%s/get-tx-info/%s", broadcastKyberAPIEndpoint, txHash)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	response := common.TradeLogGeoInfoResp{}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			log.Printf("Response body close error: %s", cErr.Error())
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", "", err
	}
	if response.Success {
		var country string
		if response.Data.Country != "" {
			return response.Data.IP, response.Data.Country, err
		}
		country, err = ipLocator.IPToCountry(response.Data.IP)
		if err != nil {
			return "", "", err
		}
		return response.Data.IP, country, err
	}
	return "", "", err
}

func enforceFromBlock(fromBlock uint64) uint64 {
	if fromBlock == 0 {
		return 0
	}
	return fromBlock - 1

}

// func (self *Fetcher) isDuplicateLog(blockNum, index uint64) bool{
// 	block, index,
// }

//SetCountryField set country field for tradelog
func SetcountryFields(ipLocator *statutil.IPLocator, l *common.TradeLog) {
	txHash := l.TxHash()
	ip, country, err := GetTradeGeo(ipLocator, txHash.Hex())
	if err != nil {
		log.Printf("LogFetcher - Getting country failed: %s", err.Error())
	}
	l.IP = ip
	l.Country = country
}

// CheckDupAndStoreTradeLog Check if the tradelog is duplicated, if it is not, manage to store it into DB
// return error if db operation is not successful
func (self *Fetcher) CheckDupAndStoreTradeLog(l common.TradeLog, timepoint uint64) error {
	var err error
	block, index, uErr := self.logStorage.LoadLastTradeLogIndex()
	if uErr == nil && (block > l.BlockNumber || (block == l.BlockNumber && index >= l.Index)) {
		log.Printf("LogFetcher - Duplicated trade log %+v (new block number %d is smaller or equal to latest block number %d and tx index %d is smaller or equal to last log tx index %d)", l, block, l.BlockNumber, index, l.Index)
	} else {
		if uErr != nil {
			log.Printf("Can not check duplicated status of current trade log, process to store it (overwrite the log if duplicated)")
		}
		err = self.logStorage.StoreTradeLog(l, timepoint)
		if err != nil {
			return err
		}
	}
	return err
}

// CheckDupAndStoreCatLog Check if the catlog is duplicated, if it is not, manage to store it into DB
// return error if db operation is not successful
func (self *Fetcher) CheckDupAndStoreCatLog(l common.SetCatLog, timepoint uint64) error {
	var err error
	block, index, uErr := self.logStorage.LoadLastCatLogIndex()
	if uErr == nil && (block > l.BlockNumber || (block == l.BlockNumber && index >= l.Index)) {
		log.Printf("LogFetcher - Duplicated trade log %+v (new block number %d is smaller or equal to latest block number %d and tx index %d is smaller or equal to last log tx index %d)", l, block, l.BlockNumber, index, l.Index)
	} else {
		if uErr != nil {
			log.Printf("Can not check duplicated status of current cat log, process to store it(overwrite the log if duplicated)")
		}
		err = self.logStorage.StoreCatLog(l)
		if err != nil {
			return err
		}
	}
	return err
}

// FetchLogs return block number that we just fetched the logs
func (self *Fetcher) FetchLogs(fromBlock uint64, toBlock uint64, timepoint uint64) (uint64, error) {

	logs, err := self.blockchain.GetLogs(fromBlock, toBlock)
	if err != nil {
		log.Printf("LogFetcher - fetching logs data from block %d failed, error: %v", fromBlock, err)
		return enforceFromBlock(fromBlock), err
	}
	if len(logs) > 0 {
		var maxBlock = enforceFromBlock(fromBlock)
		for _, il := range logs {
			// If there is log conversion error, print error and continue to the next log
			if il.Type() == "TradeLog" {
				l, ok := il.(common.TradeLog)
				if !ok {
					log.Printf("LogFetcher: ERROR cannot convert log (%v) to tradelog", il)
					continue
				}
				SetcountryFields(self.ipLocator, &l)
				if dbErr := self.CheckDupAndStoreTradeLog(l, timepoint); dbErr != nil {
					log.Printf("LogFetcher - at block %d, storing trade log failed, stop at current block and wait till next ticker, err: %+v", l.BlockNo(), dbErr)
					return maxBlock, dbErr
				}
			} else if il.Type() == "SetCatLog" {
				l, ok := il.(common.SetCatLog)
				if !ok {
					log.Printf("LogFetcher: ERROR cannot convert log (%v) to catlog", il)
					continue
				}
				if dbErr := self.CheckDupAndStoreCatLog(l, timepoint); dbErr != nil {
					log.Printf("LogFetcher - at block %d, storing cat log failed, stop at current block and wait till next ticker, err: %+v", l.BlockNo(), dbErr)
					return maxBlock, dbErr
				}
			}
			if il.BlockNo() > maxBlock {
				maxBlock = il.BlockNo()
				if err := self.logStorage.UpdateLogBlock(maxBlock, common.GetTimepoint()); err != nil {
					log.Printf("Update log block error: %s", err.Error())
				}
			}
		}
		return maxBlock, nil
	}
	return enforceFromBlock(fromBlock), nil
}

func checkWalletAddress(walletAddr ethereum.Address) bool {
	cap := big.NewInt(0)
	cap.Exp(big.NewInt(2), big.NewInt(128), big.NewInt(0))
	if walletAddr.Big().Cmp(cap) < 0 {
		return false
	}
	return true
}

func getTimestampFromTimeZone(t uint64, freq string) uint64 {
	result := uint64(0)
	ui64Day := uint64(time.Hour * 24)
	switch freq {
	case "m", "M":
		result = t / uint64(time.Minute) * uint64(time.Minute)
	case "h", "H":
		result = t / uint64(time.Hour) * uint64(time.Hour)
	case "d", "D":
		result = t / ui64Day * ui64Day
	default:
		offset, _ := strconv.ParseInt(strings.TrimPrefix(freq, "utc"), 10, 64)
		ui64offset := uint64(int64(time.Hour) * offset)
		if offset > 0 {
			result = (t+ui64offset)/ui64Day*ui64Day + ui64offset
		} else {
			offset = 0 - offset
			result = (t-ui64offset)/ui64Day*ui64Day - ui64offset
		}
	}
	return result
}

// getTokenFromAddress do a in-mem lookup for a token matched with the input address
// if not found, it will attempt to get it from core.
// if both measure still return no token, it will return error
func (self *Fetcher) getTokenFromAddress(addr ethereum.Address) (common.Token, error) {
	self.mu.RLock()
	var err error
	token, ok := self.addressLookup[addr]
	self.mu.RUnlock()
	if !ok {
		token, err = self.setting.GetTokenByAddress(addr)
		if err != nil {
			return token, err
		}
		self.mu.Lock()
		//update in-mem lookup
		self.addressLookup[addr] = token
		self.mu.Unlock()
	}
	return token, nil
}

//This function return srcAmount, destAmount, ethAmount and burnFee information of a trade log respectively
func (self *Fetcher) getTradeInfo(trade common.TradeLog) (float64, float64, float64, float64, error) {
	var srcAmount, destAmount, ethAmount, burnFee float64
	srcAddr := common.AddrToString(trade.SrcAddress)
	srcToken, err := self.getTokenFromAddress(ethereum.HexToAddress(srcAddr))
	if err != nil {
		log.Printf("get token from address: %s - %s", err.Error(), srcAddr)
		return srcAmount, destAmount, ethAmount, burnFee, err
	}
	srcAmount = common.BigToFloat(trade.SrcAmount, srcToken.Decimals)
	if srcToken.IsETH() {
		ethAmount = srcAmount
	}

	dstAddr := common.AddrToString(trade.DestAddress)
	destToken, err := self.getTokenFromAddress(ethereum.HexToAddress(dstAddr))
	if err != nil {
		log.Printf("get dest token from address: %s - %s", err.Error(), dstAddr)
		return srcAmount, destAmount, ethAmount, burnFee, err
	}
	destAmount = common.BigToFloat(trade.DestAmount, destToken.Decimals)
	if destToken.IsETH() {
		ethAmount = destAmount
	} else if trade.EtherReceivalAmount != nil {
		// Token-Token
		receivalAmount := common.BigToFloat(trade.EtherReceivalAmount, ethDecimals)
		ethAmount = receivalAmount
	}

	if trade.BurnFee != nil {
		burnFee = common.BigToFloat(trade.BurnFee, ethDecimals)
	}

	return srcAmount, destAmount, ethAmount, burnFee, nil
}

func (self *Fetcher) aggregateCountryStats(trade common.TradeLog,
	countryStats map[string]common.MetricStatsTimeZone, allFirstTradeEver map[ethereum.Address]uint64,
	kycEdUsers map[string]uint64) error {
	userAddr := common.AddrToString(trade.UserAddress)

	// ensure backward compatible.
	if trade.Country == "" {
		trade.Country = statutil.UnknownCountry
	}

	err := self.statStorage.SetCountry(trade.Country)
	if err != nil {
		log.Printf("Cannot store country: %s", err.Error())
		return err
	}
	_, _, ethAmount, burnFee, err := self.getTradeInfo(trade)
	if err != nil {
		log.Printf("get trade info error: %s", err.Error())
		return err
	}
	var kycEd bool
	regTime, exist := kycEdUsers[userAddr]
	if exist && regTime < trade.Timestamp {
		kycEd = true
	}
	self.aggregateMetricStat(trade, trade.Country, ethAmount, burnFee, countryStats, kycEd, allFirstTradeEver)
	return err
}

func (self *Fetcher) aggregateWalletStats(trade common.TradeLog,
	walletStats map[string]common.MetricStatsTimeZone, allFirstTradeEver map[ethereum.Address]uint64, kycEdUsers map[string]uint64) error {
	userAddr := common.AddrToString(trade.UserAddress)
	if checkWalletAddress(trade.WalletAddress) {
		if err := self.statStorage.SetWalletAddress(trade.WalletAddress); err != nil {
			log.Printf("Set wallet address error: %s", err.Error())
		}
	}
	_, _, ethAmount, burnFee, err := self.getTradeInfo(trade)
	if err != nil {
		return err
	}
	var kycEd bool
	regTime, exist := kycEdUsers[userAddr]
	if exist && regTime < trade.Timestamp {
		kycEd = true
	}
	self.aggregateMetricStat(trade, common.AddrToString(trade.WalletAddress), ethAmount, burnFee, walletStats, kycEd, allFirstTradeEver)
	return nil
}

func (self *Fetcher) aggregateTradeSumary(trade common.TradeLog,
	tradeSummary map[string]common.MetricStatsTimeZone, allFirstTradeEver map[ethereum.Address]uint64, kycEdUsers map[string]uint64) error {

	userAddr := common.AddrToString(trade.UserAddress)
	_, _, ethAmount, burnFee, err := self.getTradeInfo(trade)
	if err != nil {
		return err
	}
	var kycEd bool
	regTime, exist := kycEdUsers[userAddr]
	if exist && regTime < trade.Timestamp {
		kycEd = true
	}
	self.aggregateMetricStat(trade, TradeSummaryKey, ethAmount, burnFee, tradeSummary, kycEd, allFirstTradeEver)
	return nil
}

func (self *Fetcher) aggregateVolumeStats(trade common.TradeLog, volumeStats map[string]common.VolumeStatsTimeZone) error {

	srcAddr := common.AddrToString(trade.SrcAddress)
	dstAddr := common.AddrToString(trade.DestAddress)
	userAddr := common.AddrToString(trade.UserAddress)
	reserveAddr := common.AddrToString(trade.ReserveAddress)

	srcAmount, destAmount, ethAmount, _, err := self.getTradeInfo(trade)
	if err != nil {
		return err
	}
	// token volume
	self.aggregateVolumeStat(trade, srcAddr, srcAmount, ethAmount, trade.FiatAmount, volumeStats)
	self.aggregateVolumeStat(trade, dstAddr, destAmount, ethAmount, trade.FiatAmount, volumeStats)

	//user volume
	self.aggregateVolumeStat(trade, userAddr, srcAmount, ethAmount, trade.FiatAmount, volumeStats)

	// reserve volume
	var assetAddr string
	var assetAmount float64
	if srcAddr != ethAddress {
		assetAddr = srcAddr
		assetAmount = srcAmount
	} else {
		assetAddr = dstAddr
		assetAmount = destAmount
	}

	// token volume
	key := fmt.Sprintf("%s_%s", reserveAddr, assetAddr)
	self.aggregateVolumeStat(trade, key, assetAmount, ethAmount, trade.FiatAmount, volumeStats)

	// eth volume
	key = fmt.Sprintf("%s_%s", reserveAddr, ethAddress)
	self.aggregateVolumeStat(trade, key, ethAmount, ethAmount, trade.FiatAmount, volumeStats)

	// country token volume
	key = fmt.Sprintf("%s_%s", trade.Country, assetAddr)
	//log.Printf("aggegate volume: %s", key)
	self.aggregateVolumeStat(trade, key, assetAmount, ethAmount, trade.FiatAmount, volumeStats)

	return nil
}

func (self *Fetcher) aggregateBurnFeeStats(trade common.TradeLog, burnFeeStats map[string]common.BurnFeeStatsTimeZone) error {

	reserveAddr := common.AddrToString(trade.ReserveAddress)
	walletAddr := common.AddrToString(trade.WalletAddress)
	_, _, _, burnFee, err := self.getTradeInfo(trade)
	if err != nil {
		return err
	}
	// reserve burn fee
	self.aggregateBurnfee(reserveAddr, burnFee, trade, burnFeeStats)

	// wallet fee
	var walletFee float64
	if trade.WalletFee != nil {
		walletFee = common.BigToFloat(trade.WalletFee, ethDecimals)
	}
	self.aggregateBurnfee(fmt.Sprintf("%s_%s", reserveAddr, walletAddr), walletFee, trade, burnFeeStats)
	return nil
}

func (self *Fetcher) aggregateUserInfo(trade common.TradeLog, userInfos map[string]common.UserInfoTimezone) error {
	userAddr := common.AddrToString(trade.UserAddress)

	_, _, ethAmount, _, err := self.getTradeInfo(trade)
	if err != nil {
		return err
	}

	email, _, err := self.userStorage.GetUserOfAddress(trade.UserAddress)
	if err != nil {
		return err
	}
	userAddrInfo, exist := userInfos[userAddr]
	if !exist {
		userAddrInfo = common.UserInfoTimezone{}
	}
	for timezone := StartTimezone; timezone <= EndTimezone; timezone++ {
		freq := fmt.Sprintf("%s%d", TimezoneBucketPrefix, timezone)
		timestamp := getTimestampFromTimeZone(trade.Timestamp, freq)
		timezoneInfo, exist := userAddrInfo[timezone]
		if !exist {
			timezoneInfo = map[uint64]common.UserInfo{}
		}
		currentUserInfo, exist := timezoneInfo[timestamp]
		if !exist {
			currentUserInfo = common.UserInfo{
				Email: email,
				Addr:  userAddr,
			}
		}
		currentUserInfo.ETHVolume += ethAmount
		currentUserInfo.USDVolume += trade.FiatAmount
		timezoneInfo[timestamp] = currentUserInfo
		userAddrInfo[timezone] = timezoneInfo
		userInfos[userAddr] = userAddrInfo
	}
	return nil
}

func (self *Fetcher) aggregateBurnfee(key string, fee float64, trade common.TradeLog, burnFeeStats map[string]common.BurnFeeStatsTimeZone) {
	for _, freq := range []string{"M", "H", "D"} {
		timestamp := getTimestampFromTimeZone(trade.Timestamp, freq)

		currentVolume, exist := burnFeeStats[key]
		if !exist {
			currentVolume = common.BurnFeeStatsTimeZone{}
		}
		dataTimeZone, exist := currentVolume[freq]
		if !exist {
			dataTimeZone = map[uint64]common.BurnFeeStats{}
		}
		data, exist := dataTimeZone[timestamp]
		if !exist {
			data = common.BurnFeeStats{}
		}
		data.TotalBurnFee += fee
		dataTimeZone[timestamp] = data
		currentVolume[freq] = dataTimeZone
		burnFeeStats[key] = currentVolume
	}
}

func (self *Fetcher) aggregateVolumeStat(
	trade common.TradeLog,
	key string,
	assetAmount, ethAmount, fiatAmount float64,
	assetVolumetStats map[string]common.VolumeStatsTimeZone) {
	for _, freq := range []string{"M", "H", "D"} {
		timestamp := getTimestampFromTimeZone(trade.Timestamp, freq)

		currentVolume, exist := assetVolumetStats[key]
		if !exist {
			currentVolume = common.VolumeStatsTimeZone{}
		}
		dataTimeZone, exist := currentVolume[freq]
		if !exist {
			dataTimeZone = map[uint64]common.VolumeStats{}
		}
		data, exist := dataTimeZone[timestamp]
		if !exist {
			data = common.VolumeStats{}
		}
		data.ETHVolume += ethAmount
		data.USDAmount += fiatAmount
		data.Volume += assetAmount
		dataTimeZone[timestamp] = data
		currentVolume[freq] = dataTimeZone
		assetVolumetStats[key] = currentVolume
	}
}

func (self *Fetcher) aggregateMetricStat(trade common.TradeLog, statKey string, ethAmount, burnFee float64,
	metricStats map[string]common.MetricStatsTimeZone,
	kycEd bool,
	allFirstTradeEver map[ethereum.Address]uint64) {
	userAddr := trade.UserAddress

	for i := StartTimezone; i <= EndTimezone; i++ {
		freq := fmt.Sprintf("%s%d", TimezoneBucketPrefix, i)
		timestamp := getTimestampFromTimeZone(trade.Timestamp, freq)
		currentMetricData, exist := metricStats[statKey]
		if !exist {
			currentMetricData = common.MetricStatsTimeZone{}
		}
		dataTimeZone, exist := currentMetricData[i]
		if !exist {
			dataTimeZone = map[uint64]common.MetricStats{}
		}
		data, exist := dataTimeZone[timestamp]
		if !exist {
			data = common.MetricStats{}
		}
		timeFirstTrade := allFirstTradeEver[trade.UserAddress]
		if timeFirstTrade == trade.Timestamp {
			data.NewUniqueAddresses++
			data.UniqueAddr++
			if kycEd {
				data.KYCEd++
			}
		} else {
			firstTradeInday, err := self.statStorage.GetFirstTradeInDay(userAddr, trade.Timestamp, i)
			if err != nil {
				log.Printf("ERROR: get first trade in day failed. %v", err)
			}
			if firstTradeInday == trade.Timestamp {
				data.UniqueAddr++
				if kycEd {
					data.KYCEd++
				}
			}
		}

		data.ETHVolume += ethAmount
		data.BurnFee += burnFee
		data.TradeCount++
		data.USDVolume += trade.FiatAmount
		dataTimeZone[timestamp] = data
		currentMetricData[i] = dataTimeZone
		metricStats[statKey] = currentMetricData
	}
	return
}

func (self *Fetcher) FetchCurrentBlock() {
	block, err := self.blockchain.CurrentBlock()
	if err != nil {
		log.Printf("Fetching current block failed: %v. Ignored.", err)
	} else {
		// update currentBlockUpdateTime first to avoid race condition
		// where fetcher is trying to fetch new rate
		self.currentBlockUpdateTime = common.GetTimepoint()
		self.currentBlock = block
	}
}
