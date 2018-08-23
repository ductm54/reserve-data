package stat

import (
	"github.com/KyberNetwork/reserve-data/common"
)

// FeeSetRateStorage it the storage interface to keep track of SetRate fee.
type FeeSetRateStorage interface {
	GetLastBlockChecked() (uint64, error)
	StoreTransaction(tx []common.SetRateTxInfo) error
	GetFeeSetRateByDay(fromTime, toTime uint64) ([]common.FeeSetRate, error)
}
