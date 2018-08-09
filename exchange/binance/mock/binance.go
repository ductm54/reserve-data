package mock

import (
	"errors"
	"math/big"

	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

const binanceTestExchangeID = "binance"

func init() {
	common.SupportedExchanges[binanceTestExchangeID] = &BinanceTestExchange{}
}

// BinanceTestExchange is the mock implementation of binance exchange, for testing purpose.
type BinanceTestExchange struct{}

func (bte *BinanceTestExchange) ID() common.ExchangeID {
	return "binance"
}
func (bte *BinanceTestExchange) Address(token common.Token) (address ethereum.Address, supported bool) {
	return ethereum.Address{}, true
}

func (bte *BinanceTestExchange) Withdraw(token common.Token, amount *big.Int, address ethereum.Address, timepoint uint64) (string, error) {
	return "withdrawid", nil
}
func (bte *BinanceTestExchange) Trade(tradeType string, base common.Token, quote common.Token, rate float64, amount float64, timepoint uint64) (id string, done float64, remaining float64, finished bool, err error) {
	return "tradeid", 10, 5, false, nil
}
func (bte *BinanceTestExchange) CancelOrder(id, base, quote string) error {
	return nil
}
func (bte *BinanceTestExchange) MarshalText() (text []byte, err error) {
	return []byte("binance"), nil
}
func (bte *BinanceTestExchange) GetExchangeInfo(pair common.TokenPairID) (common.ExchangePrecisionLimit, error) {
	return common.ExchangePrecisionLimit{}, nil
}
func (bte *BinanceTestExchange) GetFee() (common.ExchangeFees, error) {
	return common.ExchangeFees{}, nil
}
func (bte *BinanceTestExchange) GetMinDeposit() (common.ExchangesMinDeposit, error) {
	return common.ExchangesMinDeposit{}, nil
}
func (bte *BinanceTestExchange) GetInfo() (common.ExchangeInfo, error) {
	return common.ExchangeInfo{}, nil
}
func (bte *BinanceTestExchange) TokenAddresses() (map[string]ethereum.Address, error) {
	return map[string]ethereum.Address{}, nil
}
func (bte *BinanceTestExchange) UpdateDepositAddress(token common.Token, address string) error {
	return nil
}
func (bte *BinanceTestExchange) GetTradeHistory(fromTime, toTime uint64) (common.ExchangeTradeHistory, error) {
	return common.ExchangeTradeHistory{}, nil
}

// GetLiveExchangeInfos of TestExchangeForSetting return a valid result for
func (bte *BinanceTestExchange) GetLiveExchangeInfos(tokenPairIDs []common.TokenPairID) (common.ExchangeInfo, error) {
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
