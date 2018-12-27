package reserve

import (
	"math/big"

	"github.com/KyberNetwork/reserve-data/common"
)

// ReserveData is the interface of of all data query methods.
// All methods' implementations must support concurrency.
type ReserveData interface {
	CurrentPriceVersion(timestamp uint64) (common.Version, error)
	GetAllPrices(timestamp uint64) (common.AllPriceResponse, error)
	GetOnePrice(id common.TokenPairID, timestamp uint64) (common.OnePriceResponse, error)

	CurrentAuthDataVersion(timestamp uint64) (common.Version, error)
	GetAuthData(timestamp uint64) (common.AuthDataResponse, error)

	// GetRate returns latest valid rates for all tokens that is before timestamp.
	GetRate(timestamp uint64) (common.AllRateResponse, error)
	// GetRates returns list of valid rates for all tokens that is collected between [fromTime, toTime).
	GetRates(fromTime, toTime uint64) ([]common.AllRateResponse, error)

	GetRecords(fromTime, toTime uint64) ([]common.ActivityRecord, error)
	GetPendingActivities() ([]common.ActivityRecord, error)

	GetGoldData(timepoint uint64) (common.GoldData, error)

	GetBTCData(timepoint uint64) (common.BTCData, error)

	UpdateFeedConfiguration(string, bool) error
	GetFeedConfiguration() ([]common.FeedConfiguration, error)

	GetExchangeStatus() (common.ExchangesStatus, error)
	UpdateExchangeStatus(exchange string, status bool, timestamp uint64) error

	UpdateExchangeNotification(exchange, action, tokenPair string, from, to uint64, isWarning bool, msg string) error
	GetNotifications() (common.ExchangeNotifications, error)

	GetTradeHistory(fromTime, toTime uint64) (common.AllTradeHistory, error)

	Run() error
	RunStorageController() error
	Stop() error
}

// ReserveCore is the interface that wrap around all interactions
// with exchanges and blockchain.
type ReserveCore interface {
	// place order
	Trade(
		exchange common.Exchange,
		tradeType string,
		base common.Token,
		quote common.Token,
		rate float64,
		amount float64,
		timestamp uint64) (id common.ActivityID, done float64, remaining float64, finished bool, err error)

	Deposit(
		exchange common.Exchange,
		token common.Token,
		amount *big.Int,
		timestamp uint64) (common.ActivityID, error)

	Withdraw(
		exchange common.Exchange,
		token common.Token,
		amount *big.Int,
		timestamp uint64) (common.ActivityID, error)

	CancelOrder(id common.ActivityID, exchange common.Exchange) error

	// blockchain related action
	SetRates(tokens []common.Token, buys, sells []*big.Int, block *big.Int, afpMid []*big.Int, msgs []string) (common.ActivityID, error)
}
