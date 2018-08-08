package blockchain

import (
	"fmt"
	"log"
	"math/big"
	"path/filepath"
	"time"

	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/common/blockchain"
	"github.com/KyberNetwork/reserve-data/settings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ethDecimals int64  = 18
	ethAddress  string = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

type StatBlockchain struct {
	*blockchain.BaseBlockchain
	wrapper        *blockchain.Contract
	pricing        *blockchain.Contract
	addressSetting addressSetting
}

func (sb *StatBlockchain) GetAddress(addressType settings.AddressName) (ethereum.Address, error) {
	return sb.addressSetting.GetAddress(addressType)
}

func (sb *StatBlockchain) GetAddresses(setType settings.AddressSetName) ([]ethereum.Address, error) {
	return sb.addressSetting.GetAddresses(setType)
}

func (sb *StatBlockchain) AddOldNetwork(addr ethereum.Address) {
	sb.addressSetting.AddAddressToSet(settings.OldNetworks, addr)
}

func (sb *StatBlockchain) AddOldBurners(addr ethereum.Address) {
	sb.addressSetting.AddAddressToSet(settings.OldBurners, addr)
}
func NewStatBlockchain(base *blockchain.BaseBlockchain, addrSetting *settings.AddressSetting) (*StatBlockchain, error) {
	pricingAddr, err := addrSetting.GetAddress(settings.Pricing)
	if err != nil {
		return nil, err
	}
	pricing := blockchain.NewContract(
		pricingAddr,
		filepath.Join(common.CurrentDir(), "pricing.abi"),
	)

	wrapperAddr, err := addrSetting.GetAddress(settings.Wrapper)
	if err != nil {
		return nil, err
	}
	wrapper := blockchain.NewContract(
		wrapperAddr,
		filepath.Join(common.CurrentDir(), "wrapper.abi"),
	)

	return &StatBlockchain{
		BaseBlockchain: base,
		wrapper:        wrapper,
		pricing:        pricing,
		addressSetting: addrSetting,
	}, nil
}

// trackingAddresses return the list of addresses that stat is watching for new events.
func (sb *StatBlockchain) trackingAddresses() ([]ethereum.Address, error) {
	var addresses []ethereum.Address
	networkAddr, err := sb.addressSetting.GetAddress(settings.Network)
	if err != nil {
		return nil, err
	}
	burnerAddr, err := sb.addressSetting.GetAddress(settings.Burner)
	if err != nil {
		return nil, err
	}
	whitelistAddr, err := sb.addressSetting.GetAddress(settings.Whitelist)
	if err != nil {
		return nil, err
	}
	internalAddr, err := sb.addressSetting.GetAddress(settings.InternalNetwork)
	if err != nil {
		return nil, err
	}

	addresses = append(addresses, networkAddr, burnerAddr, whitelistAddr, internalAddr)
	oldNetworks, err := sb.addressSetting.GetAddresses(settings.OldNetworks)
	if err != nil {
		log.Printf("WARNING: can't get old network addresses (%s)", err)
	}
	oldBurners, err := sb.addressSetting.GetAddresses(settings.OldBurners)
	if err != nil {
		log.Printf("WARNING: can't get old burners addresses (%s)", err)
	}
	addresses = append(addresses, oldNetworks...)
	addresses = append(addresses, oldBurners...)
	return addresses, nil
}

func (sb *StatBlockchain) GetRawLogs(fromBlock uint64, toBlock uint64) ([]types.Log, error) {
	addresses, err := sb.trackingAddresses()
	if err != nil {
		return nil, err
	}
	var (
		from = big.NewInt(int64(fromBlock))
		to   = big.NewInt(int64(toBlock))
	)
	param := common.NewFilterQuery(
		from,
		to,
		addresses,
		[][]ethereum.Hash{
			{
				ethereum.HexToHash(tradeEvent),
				ethereum.HexToHash(burnFeeEvent),
				ethereum.HexToHash(feeToWalletEvent),
				ethereum.HexToHash(userCatEvent),
				ethereum.HexToHash(etherReceivalEvent),
			},
		},
	)

	log.Printf("LogFetcher - fetching logs data from block %d, to block %d", fromBlock, to.Uint64())
	return sb.BaseBlockchain.GetLogs(param)
}

// GetLogs gets raw logs from blockchain and process it before returning.
func (sb *StatBlockchain) GetLogs(fromBlock uint64, toBlock uint64) ([]common.KNLog, error) {
	var (
		err      error
		result   []common.KNLog
		noCatLog = 0
	)

	// get all logs from fromBlock to best block
	logs, err := sb.GetRawLogs(fromBlock, toBlock)
	if err != nil {
		return result, err
	}

	for _, logItem := range logs {
		if logItem.Removed {
			log.Printf("LogFetcher - Log is ignored because it is removed due to chain reorg")
			continue
		}

		if len(logItem.Topics) == 0 {
			log.Printf("Getting empty zero topic list. This shouldn't happen and is Ethereum responsibility.")
			continue
		}

		ts, err := sb.InterpretTimestamp(
			logItem.BlockNumber,
			logItem.Index,
		)
		if err != nil {
			return result, err
		}

		topic := logItem.Topics[0]
		switch topic.Hex() {
		case userCatEvent:
			addr, cat := logDataToCatLog(logItem.Data)
			result = append(result, common.SetCatLog{
				Timestamp:       ts,
				BlockNumber:     logItem.BlockNumber,
				TransactionHash: logItem.TxHash,
				Index:           logItem.Index,
				Address:         addr,
				Category:        cat,
			})
			noCatLog++
		case feeToWalletEvent, burnFeeEvent, etherReceivalEvent, tradeEvent:
			if result, err = updateTradeLogs(result, logItem, ts); err != nil {
				return result, err
			}
		default:
			log.Printf("Unknown topic: %s", topic.Hex())
		}
	}

	for i, logItem := range result {
		tradeLog, ok := logItem.(common.TradeLog)
		if !ok {
			continue
		}

		ethRate := sb.GetEthRate(tradeLog.Timestamp / 1000000)
		if ethRate != 0 {
			result[i] = calculateFiatAmount(tradeLog, ethRate)
		}
	}

	log.Printf("LogFetcher - Fetched %d trade logs, %d cat logs", len(result)-noCatLog, noCatLog)
	return result, nil
}

func (sb *StatBlockchain) GetReserveRates(
	atBlock, currentBlock uint64, reserveAddress ethereum.Address,
	tokens []common.Token) (common.ReserveRates, error) {
	result := common.ReserveTokenRateEntry{}
	rates := common.ReserveRates{}
	rates.Timestamp = common.GetTimepoint()

	srcAddresses := []ethereum.Address{}
	destAddresses := []ethereum.Address{}
	for _, token := range tokens {
		srcAddresses = append(srcAddresses, ethereum.HexToAddress(token.Address), ethereum.HexToAddress(ethAddress))
		destAddresses = append(destAddresses, ethereum.HexToAddress(ethAddress), ethereum.HexToAddress(token.Address))
	}

	opts := sb.GetCallOpts(atBlock)
	reserveRate, sanityRate, err := sb.GeneratedGetReserveRates(opts, reserveAddress, srcAddresses, destAddresses)
	if err != nil {
		return rates, err
	}

	rates.BlockNumber = atBlock
	rates.ToBlockNumber = currentBlock
	rates.ReturnTime = common.GetTimepoint()
	for index, token := range tokens {
		rateEntry := common.ReserveRateEntry{}
		rateEntry.BuyReserveRate = common.BigToFloat(reserveRate[index*2+1], 18)
		rateEntry.BuySanityRate = common.BigToFloat(sanityRate[index*2+1], 18)
		rateEntry.SellReserveRate = common.BigToFloat(reserveRate[index*2], 18)
		rateEntry.SellSanityRate = common.BigToFloat(sanityRate[index*2], 18)
		result[fmt.Sprintf("ETH-%s", token.ID)] = rateEntry
	}
	rates.Data = result

	return rates, err
}

func (sb *StatBlockchain) GeneratedGetReserveRates(
	opts blockchain.CallOpts,
	reserveAddress ethereum.Address,
	srcAddresses []ethereum.Address,
	destAddresses []ethereum.Address) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	timeOut := 2 * time.Second
	err := sb.Call(timeOut, opts, sb.wrapper, out, "getReserveRate", reserveAddress, srcAddresses, destAddresses)
	if err != nil {
		log.Println("cannot get reserve rates: ", err.Error())
	}
	return *ret0, *ret1, err
}

func (sb *StatBlockchain) GetPricingMethod(inputData string) (*abi.Method, error) {
	abiPricing := &sb.pricing.ABI
	inputDataByte, err := hexutil.Decode(inputData)
	if err != nil {
		log.Printf("Cannot decode data: %v", err)
		return nil, err
	}
	method, err := abiPricing.MethodById(inputDataByte)
	if err != nil {
		return nil, err
	}
	return method, nil
}
