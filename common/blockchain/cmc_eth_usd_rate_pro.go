package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/KyberNetwork/reserve-data/common"
)

const (
	cmcProURL        = "https://pro-api.coinmarketcap.com"
	ETHRateURL       = "/v1/cryptocurrency/quotes/latest"
	apiKeyHeaderName = "X-CMC_PRO_API_KEY"
)

type CMCProInterface interface {
	GetETHRate() (float64, error)
}

type usdField struct {
	Price float64 `json:"price"`
}

type quoteField struct {
	USD usdField `json:"USD"`
}

type ethField struct {
	Quote quoteField `json:"quote"`
}

type dataField struct {
	ETH ethField `json:"ETH"`
}
type CmcProQuotesReply struct {
	Data dataField `json:"data"`
}
type CMCAPIKey struct {
	Key string `json:"cmc_key"`
}

type CMCProClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewCMCProClient(timeOut time.Duration, path string) (*CMCProClient, error) {
	log.Printf("path is %s", path)
	raw, err := ioutil.ReadFile(path)
	log.Printf("raw is %s", raw)
	if err != nil {
		return nil, err
	}
	key := CMCAPIKey{}
	if err = json.Unmarshal(raw, &key); err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Timeout: timeOut,
	}
	log.Printf("key is %s", key.Key)
	client := CMCProClient{
		httpClient: httpClient,
		apiKey:     key.Key,
	}
	return &client, nil
}

func (cmcPro *CMCProClient) newRequest(url string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// Add header
	req.Header.Add("Accept", "application/json")
	req.Header.Add(apiKeyHeaderName, cmcPro.apiKey)

	// Create raw query, add in params
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (cmcPro *CMCProClient) getResponse(req *http.Request) ([]byte, error) {
	var resbody []byte
	resp, err := cmcPro.httpClient.Do(req)
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

func (cmcPro *CMCProClient) GetETHRate() (float64, error) {
	url := cmcProURL + ETHRateURL
	params := map[string]string{
		"symbol": "ETH",
	}
	req, err := cmcPro.newRequest(url, params)
	//create request
	if err != nil {
		return 0, err
	}

	//do the request and return the reply
	resbody, err := cmcPro.getResponse(req)
	if err != nil {
		return 0, err
	}
	var cmcQuote CmcProQuotesReply
	if err := json.Unmarshal(resbody, &cmcQuote); err != nil {
		return 0, err
	}
	return cmcQuote.Data.ETH.Quote.USD.Price, nil
}
