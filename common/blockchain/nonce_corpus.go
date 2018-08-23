package blockchain

import (
	"math/big"

	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NonceCorpus is the interface to keep track of transaction count of an ethereum account.
type NonceCorpus interface {
	GetAddress() ethereum.Address
	GetNextNonce(ethclient *ethclient.Client) (*big.Int, error)
	MinedNonce(ethclient *ethclient.Client) (*big.Int, error)
}
