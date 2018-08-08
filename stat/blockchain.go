package stat

import (
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/settings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethereum "github.com/ethereum/go-ethereum/common"
)

// Blockchain is the interface wraps around all stat methods to interact
// with blockchain.
type Blockchain interface {
	CurrentBlock() (uint64, error)
	GetLogs(fromBlock uint64, toBlock uint64) ([]common.KNLog, error)
	GetReserveRates(atBlock, currentBlock uint64, reserveAddress ethereum.Address, tokens []common.Token) (common.ReserveRates, error)
	GetPricingMethod(inputData string) (*abi.Method, error)
	GetAddress(addressType settings.AddressName) (ethereum.Address, error)
	GetAddresses(setType settings.AddressSetName) ([]ethereum.Address, error)
}
