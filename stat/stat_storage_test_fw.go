package stat

import (
	"errors"
	"fmt"
	"strings"

	"github.com/KyberNetwork/reserve-data/common"
)

type StatStorageTest struct {
	storage StatStorage
}

const (
	TESTWALLETADDR string = "0xdd61803d4a56c597e0fc864f7a20ec7158c6cba5"
	TESTCOUNTRY    string = "unknown"
	TESTASSETADDR  string = "0x2aab2b157a03915c8a73adae735d0cf51c872f31"
	TESTUSERADDR   string = "0x778599Dd7893C8166D313F0F9B5F6cbF7536c293"
)

func NewStatStorageTest(storage StatStorage) *StatStorageTest {
	return &StatStorageTest{storage}
}

func (self *StatStorageTest) TestTradeStatsSummary() error {
	var err error
	mtStat := common.MetricStats{
		10.0,
		4567.8,
		11.1,
		3,
		2,
		4,
		5,
		0,
		0,
	}
	tzmtStat := common.MetricStatsTimeZone{0: {0: mtStat}}
	updates := map[string]common.MetricStatsTimeZone{"trade_summary": tzmtStat}
	if err := self.storage.SetTradeSummary(updates); err != nil {
		return err
	}
	tradeSum, err := self.storage.GetTradeSummary(0, 86400000, 0)
	if err != nil {
		return err
	}
	//log.Println(tradeSum)
	result, ok := (tradeSum[0]).(common.MetricStats)
	if !ok {
		return errors.New("Type mismatched: get trade stat summary return wrong type")
	}
	usdVol := (result.USDVolume)
	if !ok {
		return errors.New("Type mismatched: get trade stat summary return missing field")
	}
	if usdVol != 4567.8 {
		return fmt.Errorf("Wrong usd volume value returned: %v expected 4567.8 ", usdVol)

	}
	return nil
}

func (self *StatStorageTest) TestWalletStats() error {
	var err error
	mtStat := common.MetricStats{
		10.0,
		4567.8,
		11.1,
		3,
		2,
		4,
		5,
		0,
		0,
	}
	tzmtStat := common.MetricStatsTimeZone{0: {0: mtStat}}
	updates := map[string]common.MetricStatsTimeZone{TESTWALLETADDR: tzmtStat}
	err = self.storage.SetWalletStat(updates)
	if err != nil {
		return err
	}
	walletStat, err := self.storage.GetWalletStats(0, 86400000, TESTWALLETADDR, 0)
	result, ok := (walletStat[0]).(common.MetricStats)
	if !ok {
		return errors.New("Type mismatched: get wallet stat return wrong type")
	}
	usdVol := (result.USDVolume)
	if !ok {
		return errors.New("Type mismatched: get wallet stat return missing field")
	}
	if usdVol != 4567.8 {
		return fmt.Errorf("Wrong usd volume value returned: %v expected 4567.8 ", usdVol)
	}
	return nil
}

func (self *StatStorageTest) TestCountryStats() error {
	var err error
	mtStat := common.MetricStats{
		10.0,
		4567.8,
		11.1,
		3,
		2,
		4,
		5,
		0,
		0,
	}
	tzmtStat := common.MetricStatsTimeZone{0: {0: mtStat}}
	updates := map[string]common.MetricStatsTimeZone{TESTCOUNTRY: tzmtStat}
	err = self.storage.SetWalletStat(updates)
	if err != nil {
		return err
	}
	walletStat, err := self.storage.GetCountryStats(0, 86400000, TESTCOUNTRY, 0)
	result, ok := (walletStat[0]).(common.MetricStats)
	if !ok {
		return errors.New("Type mismatched: get country stats return wrong type")
	}
	usdVol := (result.USDVolume)
	if !ok {
		return errors.New("Type mismatched: get country stats return missing field")
	}
	if usdVol != 4567.8 {
		return fmt.Errorf("Wrong usd volume value returned: %v expected 4567.8 ", usdVol)

	}
	return nil
}

func (self *StatStorageTest) TestVolumeStats() error {
	var err error
	vlStat := common.VolumeStats{
		10.0,
		4567.8,
		11.1,
	}
	tzvlStat := common.VolumeStatsTimeZone{"D": {0: vlStat}}
	updates := map[string]common.VolumeStatsTimeZone{TESTASSETADDR: tzvlStat}
	err = self.storage.SetVolumeStat(updates)
	if err != nil {
		return err
	}
	assetVol, err := self.storage.GetAssetVolume(0, 86400000, "D", TESTASSETADDR)
	result, ok := (assetVol[0]).(common.VolumeStats)
	if !ok {
		return errors.New("Type mismatched: get volume stat return wrong type")
	}
	usdVol := (result.USDAmount)
	if !ok {
		return errors.New("Type mismatched: get volume stat return missing field")
	}
	if usdVol != 4567.8 {
		return fmt.Errorf("Wrong usd volume value returned: %v expected 4567.8 ", usdVol)

	}
	assetVol, err = self.storage.GetUserVolume(0, 86400000, "D", TESTASSETADDR)
	result, ok = (assetVol[0]).(common.VolumeStats)
	if !ok {
		return errors.New("Type mismatched: get user volume summary return wrong type")
	}
	usdVol = (result.USDAmount)
	if !ok {
		return errors.New("Type mismatched: get user volume summary return missing field")
	}
	if usdVol != 4567.8 {
		return fmt.Errorf("Wrong usd volume value returned: %v expected 4567.8 ", usdVol)
	}

	return nil

}

func (self *StatStorageTest) TestBurnFee() error {
	var err error
	bfStat := common.BurnFeeStats{
		4567.8,
	}

	tzbfStat := common.BurnFeeStatsTimeZone{"D": {0: bfStat}}
	updates := map[string]common.BurnFeeStatsTimeZone{TESTASSETADDR: tzbfStat}
	err = self.storage.SetBurnFeeStat(updates)
	if err != nil {
		return err
	}
	burnFee, err := self.storage.GetBurnFee(0, 86400000, "D", TESTASSETADDR)
	//Note : This is only temporary, burn fee return needs to be casted to common.BurnFeeStats for consistent in design
	result, ok := (burnFee[0]).(float64)
	if !ok {
		return errors.New(" Type mismatched: get burn fee return wrong type")
	}
	burnVol := (result)
	if !ok {
		return errors.New("Type mismatched: get burn fee return missing field")
	}
	if burnVol != 4567.8 {
		return fmt.Errorf("Wrong burn fee value returned: %v expected 4567.8 ", burnVol)

	}

	updates = map[string]common.BurnFeeStatsTimeZone{fmt.Sprintf("%s_%s", TESTASSETADDR, TESTWALLETADDR): tzbfStat}
	err = self.storage.SetBurnFeeStat(updates)
	if err != nil {
		return err
	}
	burnFee, err = self.storage.GetWalletFee(0, 86400000, "D", TESTASSETADDR, TESTWALLETADDR)
	result, ok = (burnFee[0]).(float64)
	if !ok {
		return errors.New("Type mismatched: get burn fee return wrong type")
	}
	burnVol = (result)
	if !ok {
		return errors.New("Type mismatched: get burn fee return missing field")
	}
	if burnVol != 4567.8 {
		return fmt.Errorf("Wrong wallet fee value returned: %v expected 4567.8 ", burnVol)

	}
	return nil
}

func (self *StatStorageTest) TestWalletAddress() error {
	var err error
	err = self.storage.SetWalletAddress("0xdd61803d4a56c597e0fc864f7a20ec7158c6cba5")
	if err != nil {
		return err
	}
	walletaddresses, err := self.storage.GetWalletAddress()
	if err != nil {
		return err
	}
	if len(walletaddresses) != 1 {
		return fmt.Errorf("expected 1 record, got %d record of wallet addresses returned", len(walletaddresses))
	}
	if walletaddresses[0] != "0xdd61803d4a56c597e0fc864f7a20ec7158c6cba5" {
		return fmt.Errorf("expected address 0xdd61803d4a56c597e0fc864f7a20ec7158c6cba5, got %s instead", walletaddresses[0])
	}
	return err
}

func (self *StatStorageTest) TestLastProcessedTradeLogTimePoint() error {
	var err error
	err = self.storage.SetLastProcessedTradeLogTimepoint(45678)
	if err != nil {
		return err
	}
	lastTimePoint, err := self.storage.GetLastProcessedTradeLogTimepoint()
	if err != nil {
		return err
	}
	if lastTimePoint != 45678 {
		return fmt.Errorf("expected last time point to be 45678, got %d instead", lastTimePoint)
	}
	return err
}

func (self *StatStorageTest) TestCountries() error {
	var err error
	err = self.storage.SetCountry("bunny")
	if err != nil {
		return err
	}
	countries, err := self.storage.GetCountries()
	if err != nil {
		return err
	}
	if len(countries) != 1 {
		return fmt.Errorf("wrong countries len, expect 1, got %d", len(countries))
	}
	if countries[0] != "bunny" {
		return fmt.Errorf("wrong country result, expect bunny, got %s", countries[0])
	}
	return err

}

func (self *StatStorageTest) TestFirstTradeEver() error {
	var err error
	userAddrs := map[string]uint64{(TESTUSERADDR + "_45678"): 45678}
	err = self.storage.SetFirstTradeEver(userAddrs)
	if err != nil {
		return err
	}

	timepoint, err := self.storage.GetFirstTradeEver(TESTUSERADDR)
	if err != nil {
		return err
	}
	if timepoint != 45678 {
		return fmt.Errorf("first trade ever error, expect timepoint 45678, got %d", timepoint)
	}
	timepoint, err = self.storage.GetFirstTradeEver(strings.ToLower(TESTUSERADDR))
	if err != nil {
		return err
	}
	if timepoint != 45678 {
		return (fmt.Errorf("first trade ever error with lower case addr, expect timepoint 45678, got %d", timepoint))
	}
	timepoint, err = self.storage.GetFirstTradeEver(strings.ToUpper(TESTUSERADDR))
	if err != nil {
		return err
	}
	if timepoint != 45678 {
		return (fmt.Errorf("first trade ever error with upper case addr, expect timepoint 45678, got %d", timepoint))
	}
	userAddrs = map[string]uint64{TESTWALLETADDR: 45678}
	err = self.storage.SetFirstTradeEver(userAddrs)
	if err != nil {
		return err
	}
	allFirstTradeEver, err := self.storage.GetAllFirstTradeEver()
	if len(allFirstTradeEver) != 2 {
		return fmt.Errorf("wrong all first trade ever  len, expect 2, got %d", len(allFirstTradeEver))
	}
	return err

}

func (self *StatStorageTest) TestFirstTradeInDay() error {
	var err error
	userAddrs := map[string]uint64{(TESTUSERADDR + "_45678"): 45678}
	err = self.storage.SetFirstTradeInDay(userAddrs)
	if err != nil {
		return err
	}

	timepoint, err := self.storage.GetFirstTradeInDay(TESTUSERADDR, 0, 0)
	if err != nil {
		return err
	}
	if timepoint != 45678 {
		return fmt.Errorf("first trade in day error, expect timepoint 45678, got %d", timepoint)
	}
	return err
	timepoint, err = self.storage.GetFirstTradeInDay(strings.ToLower(TESTUSERADDR), 0, 0)
	if timepoint != 45678 {
		return (fmt.Errorf("first trade in day error with lower case addr, expect timepoint 45678, got %d", timepoint))
	}
	return err
	timepoint, err = self.storage.GetFirstTradeInDay(strings.ToUpper(TESTUSERADDR), 0, 0)
	if timepoint != 45678 {
		return (fmt.Errorf("first trade in day error with upper case addr, expect timepoint 45678, got %d", timepoint))
	}
	return err

}
