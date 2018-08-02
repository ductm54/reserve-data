package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

const (
	nameKey                    = "name"
	nonceParamKey              = "nonce"
	addrParamKey               = "address"
	idParamKey                 = "ID"
	getInternalTokensEndpoint  = "setting/internal-tokens"
	getActiveTokensEndpoint    = "setting/active-tokens"
	getTokenByAddressEndpoint  = "setting/token-by-address"
	getActiveTokenByIDEndpoint = "setting/active-token-by-id"
	getAddressEndpoint         = "setting/address"
	getAddressesEndpoint       = "setting/addresses"
	readyToServeEndpoint       = "setting/ping"
)

type clientAuthentication interface {
	KNSign(message string) string
}

// SettingClient is a http Client used to query setting from core APIs
type SettingClient struct {
	authEngine clientAuthentication
	client     *http.Client
	coreURL    string
}

func NewSettingClient(authEng clientAuthentication,
	timeOut time.Duration, url string) *SettingClient {
	httpClient := &http.Client{
		Timeout: timeOut,
	}
	return &SettingClient{authEngine: authEng,
		client:  httpClient,
		coreURL: url}
}

// SortByKey sort all the params by key in string order
// This is required for the request to be signed correctly
func sortByKey(params map[string]string) map[string]string {
	newParams := make(map[string]string, len(params))
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		newParams[key] = params[key]
	}
	return newParams
}

func (sc *SettingClient) sign(req *http.Request, message string, nonce string) {
	signed := sc.authEngine.KNSign(message)
	req.Header.Add("nonce", nonce)
	req.Header.Add("signed", signed)
}

func (sc *SettingClient) newRequest(method, url string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	// Add header
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Create raw query
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	//sign
	nonce, ok := params["nonce"]
	if !ok {
		log.Printf("there was no nonce")
	} else {
		sc.sign(req, q.Encode(), nonce)
	}

	return req, nil
}

func (sc *SettingClient) getReponse(method, url string, params map[string]string) ([]byte, error) {
	params = sortByKey(params)
	req, err := sc.newRequest(method, url, params)
	//create request
	if err != nil {
		return nil, err
	}

	//do the request and return the reply
	var resbody []byte
	resp, err := sc.client.Do(req)
	if err != nil {
		return resbody, err
	}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			log.Printf("Response body close error: %s", cErr.Error())
		}
	}()
	if resp.StatusCode == http.StatusOK {
		resbody, err = ioutil.ReadAll(resp.Body)
	} else {
		log.Printf("The reply code %v was unexpected", resp.StatusCode)
		resbody, err = ioutil.ReadAll(resp.Body)
	}
	log.Printf("request to %s, got response: \n %s \n\n", req.URL, common.TruncStr(resbody))
	return resbody, err
}

type TokensReply struct {
	Data    []common.Token
	Success bool
}

func (sc *SettingClient) GetInternalTokens() ([]common.Token, error) {
	url := fmt.Sprintf("%s/%s", sc.coreURL, getInternalTokensEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	params[nonceParamKey] = nonce
	response, err := sc.getReponse(http.MethodGet, url, params)
	if err != nil {
		return nil, err
	}
	var internalTokenReply TokensReply
	err = json.Unmarshal(response, &internalTokenReply)
	if err != nil {
		return nil, err
	}
	return internalTokenReply.Data, nil
}

func (sc *SettingClient) GetActiveTokens() ([]common.Token, error) {
	url := fmt.Sprintf("%s/%s", sc.coreURL, getActiveTokensEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	params[nonceParamKey] = nonce
	response, err := sc.getReponse(http.MethodGet, url, params)
	if err != nil {
		return nil, err
	}
	var activeTokensReply TokensReply
	err = json.Unmarshal(response, &activeTokensReply)
	if err != nil {
		return nil, err
	}
	return activeTokensReply.Data, nil
}

type TokenReply struct {
	Data    common.Token
	Success bool
}

func (sc *SettingClient) GetTokenByAddress(addr ethereum.Address) (common.Token, error) {
	url := fmt.Sprintf("%s/%s", sc.coreURL, getTokenByAddressEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	addrStr := addr.Hex()
	params[nonceParamKey] = nonce
	params[addrParamKey] = addrStr
	var tokenReply TokenReply
	response, err := sc.getReponse(http.MethodGet, url, params)
	if err != nil {
		return common.Token{}, err
	}
	err = json.Unmarshal(response, &tokenReply)
	if err != nil {
		return common.Token{}, err
	}
	return tokenReply.Data, nil
}

func (sc *SettingClient) GetActiveTokenByID(id string) (common.Token, error) {
	url := fmt.Sprintf("%s/%s", sc.coreURL, getActiveTokenByIDEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	params[nonceParamKey] = nonce
	params[idParamKey] = id
	var tokenReply TokenReply
	response, err := sc.getReponse(http.MethodGet, url, params)
	if err != nil {
		return common.Token{}, err
	}
	err = json.Unmarshal(response, &tokenReply)
	if err != nil {
		return common.Token{}, err
	}
	return tokenReply.Data, nil
}

type AddressReply struct {
	Data    ethereum.Address
	Success bool
}

func (sc *SettingClient) GetAddress(addressType AddressName) (ethereum.Address, error) {
	url := fmt.Sprintf("%s/%s", sc.coreURL, getAddressEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	params[nonceParamKey] = nonce
	params[nameKey] = addressType.String()
	var addressReply AddressReply
	response, err := sc.getReponse(http.MethodGet, url, params)
	if err != nil {
		return ethereum.Address{}, err
	}
	err = json.Unmarshal(response, &addressReply)
	if err != nil {
		return ethereum.Address{}, err
	}
	return addressReply.Data, nil
}

type AddressesReply struct {
	Data    []ethereum.Address
	Success bool
}

func (sc *SettingClient) GetAddresses(setType AddressSetName) ([]ethereum.Address, error) {
	url := fmt.Sprintf("%s/%s", sc.coreURL, getAddressesEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	params[nonceParamKey] = nonce
	params[nameKey] = setType.String()
	var addressesReply AddressesReply
	response, err := sc.getReponse(http.MethodGet, url, params)
	if err != nil {
		return []ethereum.Address{}, err
	}
	err = json.Unmarshal(response, &addressesReply)
	if err != nil {
		return []ethereum.Address{}, err
	}
	return addressesReply.Data, nil
}

// ReadyToServe is called prior to running stat functions to make sure core is up
func (sc *SettingClient) ReadyToServe() error {
	url := fmt.Sprintf("%s/%s", sc.coreURL, readyToServeEndpoint)
	params := make(map[string]string)
	nonce := strconv.FormatUint(common.GetTimepoint(), 10)
	params[nonceParamKey] = nonce
	_, err := sc.getReponse(http.MethodGet, url, params)
	return err
}
