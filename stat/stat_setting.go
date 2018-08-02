package stat

import (
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/settings"
	ethereum "github.com/ethereum/go-ethereum/common"
)

type Setting interface {
	GetInternalTokens() ([]common.Token, error)
	GetActiveTokens() ([]common.Token, error)
	GetTokenByAddress(addr ethereum.Address) (common.Token, error)
	GetActiveTokenByID(id string) (common.Token, error)
	GetAddress(addressType settings.AddressName) (ethereum.Address, error)
	GetAddresses(setType settings.AddressSetName) ([]ethereum.Address, error)
	// ReadyToServe is called prior to running stat functions to make sure core is up
	ReadyToServe() error
}
