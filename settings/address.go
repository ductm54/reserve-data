package settings

import (
	"errors"

	ethereum "github.com/ethereum/go-ethereum/common"
)

var ErrNoAddr = errors.New("cannot find the address")

func (setting *Settings) GetAddresses(setName AddressSetName) ([]ethereum.Address, error) {
	return setting.Address.GetAddresses(setName)
}
func (setting *Settings) GetAddress(name AddressName) (ethereum.Address, error) {
	return setting.Address.GetAddress(name)
}

func (addrsetting *AddressSetting) GetAddress(name AddressName) (ethereum.Address, error) {
	addr, ok := addrsetting.Addresses[name]
	if !ok {
		return ethereum.Address{}, ErrNoAddr
	}
	return addr, nil
}

func (addrsetting *AddressSetting) GetAddresses(setName AddressSetName) ([]ethereum.Address, error) {
	addrs, ok := addrsetting.AddressSets[setName]
	if !ok {
		return nil, ErrNoAddr
	}
	return addrs, nil
}

// GetAllAddresses return all the address setting in cores.
func (setting *Settings) GetAllAddresses() (map[string]interface{}, error) {
	allAddress := make(map[string]interface{})
	for name, addr := range setting.Address.Addresses {
		allAddress[name.String()] = addr
	}
	for setName, addrs := range setting.Address.AddressSets {
		allAddress[setName.String()] = addrs
	}
	return allAddress, nil
}
func (addrsetting *AddressSetting) AddAddressToSet(setName AddressSetName, address ethereum.Address) {
	//no need to handle availability here, if the record is not exist yet it's still be able to update new address-
	addrs, _ := addrsetting.AddressSets[setName]
	addrs = append(addrs, address)
	addrsetting.AddressSets[setName] = addrs
}
