package core

import (
	"github.com/KyberNetwork/reserve-data/common"
)

// ActivityStorage is the interface contains all database operations of core.
type ActivityStorage interface {
	Record(
		action string,
		id common.ActivityID,
		destination string,
		params map[string]interface{},
		result map[string]interface{},
		estatus string,
		mstatus string,
		timepoint uint64) error
	HasPendingDeposit(
		token common.Token, exchange common.Exchange) (bool, error)

	GetActivity(id common.ActivityID) (common.ActivityRecord, error)

	// PendingSetrate return the last pending set rate and number of pending
	// transactions.
	PendingSetrate(minedNonce uint64) (*common.ActivityRecord, uint64, error)
}
