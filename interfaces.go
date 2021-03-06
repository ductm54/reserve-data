package reserve

import (
	"math/big"

	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

// ReserveStats is the interface of all statistic methods.
type ReserveStats interface {
	GetTradeLogs(fromTime uint64, toTime uint64) ([]common.TradeLog, error)
	GetCatLogs(fromTime uint64, toTime uint64) ([]common.SetCatLog, error)
	GetAssetVolume(fromTime, toTime uint64, freq, asset string) (common.StatTicks, error)
	GetBurnFee(fromTime, toTime uint64, freq, reserveAddr string) (common.StatTicks, error)
	GetWalletFee(fromTime, toTime uint64, freq, reserveAddr, walletAddr string) (common.StatTicks, error)
	GetWalletAddresses() ([]string, error)
	GetUserVolume(fromTime, toTime uint64, freq, userAddr string) (common.StatTicks, error)
	GetUsersVolume(fromTime, toTime uint64, freq string, userAddrs []string) (common.UsersVolume, error)
	GetReserveVolume(fromTime, toTime uint64, freq, reserveAddr, tokenID string) (common.StatTicks, error)
	GetTradeSummary(fromTime, toTime uint64, timezone int64) (common.StatTicks, error)

	GetCapByUser(userID string) (*common.UserCap, error)
	GetCapByAddress(addr ethereum.Address) (*common.UserCap, error)
	GetTxCapByAddress(addr ethereum.Address) (*big.Int, bool, error)
	ExceedDailyLimit(addr ethereum.Address) (bool, error)
	GetPendingAddresses() ([]string, error)
	GetWalletStats(fromTime uint64, toTime uint64, walletAddr string, timezone int64) (common.StatTicks, error)
	GetReserveRates(fromTime, toTime uint64, reserveAddr ethereum.Address) ([]common.ReserveRates, error)
	UpdateUserAddresses(userID string, addresses []ethereum.Address, timestamps []uint64) error
	UpdatePriceAnalyticData(timestamp uint64, value []byte) error
	GetPriceAnalyticData(fromTime uint64, toTime uint64) ([]common.AnalyticPriceResponse, error)

	GetGeoData(fromTime, toTime uint64, country string, tzparam int64) (common.StatTicks, error)
	GetHeatMap(fromTime, toTime uint64, tzparam int64) (common.HeatmapResponse, error)
	GetTokenHeatmap(fromTime, toTime uint64, token, freq string) (common.TokenHeatmapResponse, error)
	GetCountries() ([]string, error)

	GetUserList(fromTime, toTime uint64, timezone int64) (common.UserListResponse, error)
	GetFeeSetRateByDay(fromTime uint64, toTime uint64) ([]common.FeeSetRate, error)

	RunStorageController() error
	Run() error
	Stop() error
}

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
