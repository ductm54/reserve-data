package http

import "github.com/KyberNetwork/reserve-data/common"

// Huobi is the HTTP server to keep track of immediate transactions between Houbi and Kyber Network.
// The reason is Huobi doesn't allow transaction directly from contract account.
type Huobi interface {
	PendingIntermediateTxs() (map[common.ActivityID]common.TXEntry, error)
}
