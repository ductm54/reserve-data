package settings

import (
	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

func (setting *Settings) UpdateAddress(name AddressName, address ethereum.Address, timestamp uint64) error {
	// if timestamp is less or equal than 0, that mean no input, to default case of getting Unix timestamp
	if timestamp <= 0 {
		timestamp = common.GetTimepoint()
	}
	return setting.Address.Storage.UpdateOneAddress(name, address.Hex(), timestamp)
}

func (setting *Settings) GetAddress(name AddressName) (ethereum.Address, error) {
	result := ethereum.Address{}
	addr, err := setting.Address.Storage.GetAddress(name)
	if err != nil {
		return result, err
	}
	return ethereum.HexToAddress(addr), err
}

// GetAllAddresses return all the address setting in cores.
func (setting *Settings) GetAllAddresses() (map[string]interface{}, error) {
	return setting.Address.Storage.GetAllAddresses()
}

func (setting *Settings) AddAddressToSet(setName AddressSetName, address ethereum.Address, timestamp uint64) error {
	// if timestamp is less or equal than 0, that mean no input, to default case of getting Unix timestamp
	if timestamp <= 0 {
		timestamp = common.GetTimepoint()
	}
	return setting.Address.Storage.AddAddressToSet(setName, address.Hex(), timestamp)
}

func (setting *Settings) GetAddresses(setName AddressSetName) ([]ethereum.Address, error) {
	result := []ethereum.Address{}
	addrs, err := setting.Address.Storage.GetAddresses(setName)
	if err != nil {
		return result, err
	}
	for _, addr := range addrs {
		result = append(result, ethereum.HexToAddress(addr))
	}
	return result, nil
}

func (setting *Settings) GetAddressVersion() (uint64, error) {
	return setting.Address.Storage.GetAddressVersion()
}
