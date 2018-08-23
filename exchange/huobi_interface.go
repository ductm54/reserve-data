package exchange

import (
	"math/big"

	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

// HuobiInterface contains the methods to interact with Huobi centralized exchange.
type HuobiInterface interface {
	GetDepthOnePair(pair common.TokenPair) (HuobiDepth, error)

	OpenOrdersForOnePair(pair common.TokenPair) (HuobiOrder, error)

	GetInfo() (HuobiInfo, error)

	GetExchangeInfo() (HuobiExchangeInfo, error)

	GetDepositAddress(token string) (HuobiDepositAddress, error)

	GetAccountTradeHistory(base, quote common.Token) (HuobiTradeHistory, error)

	Withdraw(
		token common.Token,
		amount *big.Int,
		address ethereum.Address) (string, error)

	Trade(
		tradeType string,
		base, quote common.Token,
		rate, amount float64,
		timepoint uint64) (HuobiTrade, error)

	CancelOrder(symbol string, id uint64) (HuobiCancel, error)

	DepositHistory(tokens []common.Token) (HuobiDeposits, error)

	WithdrawHistory(token []common.Token) (HuobiWithdraws, error)

	OrderStatus(symbol string, id uint64) (HuobiOrder, error)
}
