package blockchain

import (
	"github.com/KyberNetwork/reserve-data/settings"
	ethereum "github.com/ethereum/go-ethereum/common"
)

type addressSetting interface {
	GetAddress(settings.AddressName) (ethereum.Address, error)
	GetAddresses(settings.AddressSetName) ([]ethereum.Address, error)
	AddAddressToSet(setName settings.AddressSetName, address ethereum.Address)
}
