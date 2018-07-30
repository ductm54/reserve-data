package common

import (
	"errors"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum/common"
)

const BinanceTestExchangeID = "binance"

func AddTestExchangeForSetting() {
	SupportedExchanges[BinanceTestExchangeID] = BinanceTestExchange{}
}

type BinanceTestExchange struct {
}

func (self BinanceTestExchange) ID() ExchangeID {
	return "binance"
}
func (self BinanceTestExchange) Address(token Token) (address ethereum.Address, supported bool) {
	return ethereum.Address{}, true
}

func (self BinanceTestExchange) Withdraw(token Token, amount *big.Int, address ethereum.Address, timepoint uint64) (string, error) {
	return "withdrawid", nil
}
func (self BinanceTestExchange) Trade(tradeType string, base Token, quote Token, rate float64, amount float64, timepoint uint64) (id string, done float64, remaining float64, finished bool, err error) {
	return "tradeid", 10, 5, false, nil
}
func (self BinanceTestExchange) CancelOrder(id, base, quote string) error {
	return nil
}
func (self BinanceTestExchange) MarshalText() (text []byte, err error) {
	return []byte("binance"), nil
}
func (self BinanceTestExchange) GetExchangeInfo(pair TokenPairID) (ExchangePrecisionLimit, error) {
	return ExchangePrecisionLimit{}, nil
}
func (self BinanceTestExchange) GetFee() (ExchangeFees, error) {
	return ExchangeFees{}, nil
}
func (self BinanceTestExchange) GetMinDeposit() (ExchangesMinDeposit, error) {
	return ExchangesMinDeposit{}, nil
}
func (self BinanceTestExchange) GetInfo() (ExchangeInfo, error) {
	return ExchangeInfo{}, nil
}
func (self BinanceTestExchange) TokenAddresses() (map[string]ethereum.Address, error) {
	return map[string]ethereum.Address{}, nil
}
func (self BinanceTestExchange) UpdateDepositAddress(token Token, address string) error {
	return nil
}
func (self BinanceTestExchange) GetTradeHistory(fromTime, toTime uint64) (ExchangeTradeHistory, error) {
	return ExchangeTradeHistory{}, nil
}

// GetLiveExchangeInfos of TestExchangeForSetting return a valid result for
func (self BinanceTestExchange) GetLiveExchangeInfos(tokenPairIDs []TokenPairID) (ExchangeInfo, error) {
	ETHKNCpairID := NewTokenPairID("KNC", "ETH")
	result := make(ExchangeInfo)
	for _, pairID := range tokenPairIDs {
		if pairID != ETHKNCpairID {
			return result, errors.New("Token pair ID is not support")
		}
		result[pairID] = ExchangePrecisionLimit{
			AmountLimit: TokenPairAmountLimit{
				Min: 1,
				Max: 900000,
			},
			Precision: TokenPairPrecision{
				Amount: 0,
				Price:  7,
			},
			PriceLimit: TokenPairPriceLimit{
				Min: 0.000192,
				Max: 0.019195,
			},
			MinNotional: 0.01,
		}
	}
	return result, nil
}
