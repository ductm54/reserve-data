package core

import (
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/settings"
	ethereum "github.com/ethereum/go-ethereum/common"
)

type Setting interface {
	GetAddress(settings.AddressName) (ethereum.Address, error)
	GetInternalTokens() ([]common.Token, error)
}
