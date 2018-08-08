package settings

import (
	"encoding/json"
	"io/ioutil"

	"github.com/KyberNetwork/reserve-data/common"
)

// AddressName is the name of ethereum address used in core.
//go:generate stringer -type=AddressName -linecomment
type AddressName int

const (
	Reserve         AddressName = iota //reserve
	Burner                             //burner
	Bank                               //bank
	Network                            //network
	Wrapper                            //wrapper
	Pricing                            //pricing
	Whitelist                          //whitelist
	InternalNetwork                    //internal_network
)

var addressNameValues = map[string]AddressName{
	"reserve":          Reserve,
	"burner":           Burner,
	"bank":             Bank,
	"network":          Network,
	"wrapper":          Wrapper,
	"pricing":          Pricing,
	"whitelist":        Whitelist,
	"internal_network": InternalNetwork,
}

// AddressNameValues returns the mapping of the string presentation
// of address name and its value.
func AddressNameValues() map[string]AddressName {
	return addressNameValues
}

// AddressSetName is the name of ethereum address set used in core.
//go:generate stringer -type=AddressSetName -linecomment
type AddressSetName int

const (
	ThirdPartyReserves AddressSetName = iota //third_party_reserves
	OldNetWorks                              //old_networks
	OldBurners                               //old_burners
)

var addressSetNameValues = map[string]AddressSetName{
	"third_party_reserves": ThirdPartyReserves,
	"old_networks":         OldNetWorks,
	"old_burners":          OldBurners,
}

// AddressSetNameValues returns the mapping of the string presentation
// of address set name and its value.
func AddressSetNameValues() map[string]AddressSetName {
	return addressSetNameValues
}

// AddressConfig type defines a list of address attribute avaiable in core.
// It is used mainly for
type AddressConfig struct {
	Bank               string   `json:"bank"`
	Reserve            string   `json:"reserve"`
	Network            string   `json:"network"`
	Wrapper            string   `json:"wrapper"`
	Pricing            string   `json:"pricing"`
	FeeBurner          string   `json:"feeburner"`
	Whitelist          string   `json:"whitelist"`
	ThirdPartyReserves []string `json:"third_party_reserves"`
	InternalNetwork    string   `json:"internal network"`
}

// AddressSetting type defines component to handle all address setting in core.
// It contains the storage interface used to query addresses.
type AddressSetting struct {
	Storage AddressStorage
}

func NewAddressSetting(addressStorage AddressStorage) (*AddressSetting, error) {
	return &AddressSetting{addressStorage}, nil
}

func (setting *Settings) loadAddressFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	addrs := AddressConfig{}
	if err = json.Unmarshal(data, &addrs); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Bank, addrs.Bank, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Reserve, addrs.Reserve, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Network, addrs.Network, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Wrapper, addrs.Wrapper, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Pricing, addrs.Pricing, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Burner, addrs.FeeBurner, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(Whitelist, addrs.Whitelist, common.GetTimepoint()); err != nil {
		return err
	}
	if err = setting.Address.Storage.UpdateOneAddress(InternalNetwork, addrs.InternalNetwork, common.GetTimepoint()); err != nil {
		return err
	}
	for _, addr := range addrs.ThirdPartyReserves {
		if err = setting.Address.Storage.AddAddressToSet(ThirdPartyReserves, addr, common.GetTimepoint()); err != nil {
			return err
		}
	}
	return nil
}
