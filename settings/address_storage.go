package settings

//AddressStorage abstracts the storage of addresses.
//Any implementation of AddressStorage must fulfilled these funcs
//All params should be in lower case
type AddressStorage interface {
	//UpdateOneAddress stores an address, with its functional name(network, burner etc.. ) as key
	//This will overwrite the value if the key is duplicated and return error if occurs
	UpdateOneAddress(name AddressName, address string, timestamp uint64) error

	//GetAddress return a single address based on its functional name(network, burner etc... ) as key
	//It will return a string as address and error if occurs.
	GetAddress(name AddressName) (string, error)

	//AddAddressToSet add a single address to a set of address, based on setName(for example, third_party_reserver)
	//It must not allow duplication of addresses(i.e, use address as key) and return error if occurs.
	AddAddressToSet(setName AddressSetName, address string, timestamp uint64) error
	//GetAddresses return all addresses in a set of address, based on setName(for example, third_party_reserver)
	//It will return error if occurs
	GetAddresses(setName AddressSetName) ([]string, error)

	//CountAddress return number of address in address setting
	//If there is error, it will return 0
	CountAddress() (uint64, error)

	//GetAllAddresses return all the address setting currently in core
	//Return error if occur.
	GetAllAddresses() (map[string]interface{}, error)

	//GetAddressVersion return the current version of address setting in core
	//Retrun error if occur.
	GetAddressVersion() (uint64, error)
}
