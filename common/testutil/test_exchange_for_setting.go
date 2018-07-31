package testutil

import (
	"errors"
	"math/big"

	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

const BinanceTestExchangeID = "binance"

func EnsureTestUtilInit() bool {
	return true
}

func init() {
	common.SupportedExchanges[BinanceTestExchangeID] = BinanceTestExchange{}
}

type BinanceTestExchange struct{}

func (self BinanceTestExchange) ID() common.ExchangeID {
	return "binance"
}
func (self BinanceTestExchange) Address(token common.Token) (address ethereum.Address, supported bool) {
	return ethereum.Address{}, true
}

func (self BinanceTestExchange) Withdraw(token common.Token, amount *big.Int, address ethereum.Address, timepoint uint64) (string, error) {
	return "withdrawid", nil
}
func (self BinanceTestExchange) Trade(tradeType string, base common.Token, quote common.Token, rate float64, amount float64, timepoint uint64) (id string, done float64, remaining float64, finished bool, err error) {
	return "tradeid", 10, 5, false, nil
}
func (self BinanceTestExchange) CancelOrder(id, base, quote string) error {
	return nil
}
func (self BinanceTestExchange) MarshalText() (text []byte, err error) {
	return []byte("binance"), nil
}
func (self BinanceTestExchange) GetExchangeInfo(pair common.TokenPairID) (common.ExchangePrecisionLimit, error) {
	return common.ExchangePrecisionLimit{}, nil
}
func (self BinanceTestExchange) GetFee() (common.ExchangeFees, error) {
	return common.ExchangeFees{}, nil
}
func (self BinanceTestExchange) GetMinDeposit() (common.ExchangesMinDeposit, error) {
	return common.ExchangesMinDeposit{}, nil
}
func (self BinanceTestExchange) GetInfo() (common.ExchangeInfo, error) {
	return common.ExchangeInfo{}, nil
}
func (self BinanceTestExchange) TokenAddresses() (map[string]ethereum.Address, error) {
	return map[string]ethereum.Address{}, nil
}
func (self BinanceTestExchange) UpdateDepositAddress(token common.Token, address string) error {
	return nil
}
func (self BinanceTestExchange) GetTradeHistory(fromTime, toTime uint64) (common.ExchangeTradeHistory, error) {
	return common.ExchangeTradeHistory{}, nil
}

// GetLiveExchangeInfos of TestExchangeForSetting return a valid result for
func (self BinanceTestExchange) GetLiveExchangeInfos(tokenPairIDs []common.TokenPairID) (common.ExchangeInfo, error) {
	ETHKNCpairID := common.NewTokenPairID("KNC", "ETH")
	result := make(common.ExchangeInfo)
	for _, pairID := range tokenPairIDs {
		if pairID != ETHKNCpairID {
			return result, errors.New("Token pair ID is not support")
		}
		result[pairID] = common.ExchangePrecisionLimit{
			AmountLimit: common.TokenPairAmountLimit{
				Min: 1,
				Max: 900000,
			},
			Precision: common.TokenPairPrecision{
				Amount: 0,
				Price:  7,
			},
			PriceLimit: common.TokenPairPriceLimit{
				Min: 0.000192,
				Max: 0.019195,
			},
			MinNotional: 0.01,
		}
	}
	return result, nil
}
