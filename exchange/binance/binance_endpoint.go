package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/exchange"
	ethereum "github.com/ethereum/go-ethereum/common"
)

// BinanceEndpoint object stand for Binance endpoint
// including signer for api call authentication,
// interf for calling api in different env
// timedelta to make sure calling api in time
type BinanceEndpoint struct {
	signer    Signer
	interf    Interface
	timeDelta int64
}

func (be *BinanceEndpoint) fillRequest(req *http.Request, signNeeded bool, timepoint uint64) {
	if req.Method == "POST" || req.Method == "PUT" || req.Method == "DELETE" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent", "binance/go")
	}
	req.Header.Add("Accept", "application/json")
	log.Printf("Bin Time Delta: %d", be.timeDelta)
	if signNeeded {
		q := req.URL.Query()
		sig := url.Values{}
		req.Header.Set("X-MBX-APIKEY", be.signer.GetKey())
		q.Set("timestamp", fmt.Sprintf("%d", int64(timepoint)+be.timeDelta-1000))
		q.Set("recvWindow", "5000")
		sig.Set("signature", be.signer.Sign(q.Encode()))
		// Using separated values map for signature to ensure it is at the end
		// of the query. This is required for /wapi apis from binance without
		// any damn documentation about it!!!
		req.URL.RawQuery = q.Encode() + "&" + sig.Encode()
	}
}

func (be *BinanceEndpoint) GetResponse(
	method string, url string,
	params map[string]string, signNeeded bool, timepoint uint64) ([]byte, error) {
	var (
		err      error
		respBody []byte
	)
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	be.fillRequest(req, signNeeded, timepoint)

	log.Printf("request to binance: %s\n", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		return respBody, err
	}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			log.Printf("Response body close error: %s", cErr.Error())
		}
	}()
	switch resp.StatusCode {
	case 429:
		err = errors.New("breaking binance request rate limit")
		break
	case 418:
		err = errors.New("ip has been auto-banned by binance for continuing to send requests after receiving 429 codes")
		break
	case 500:
		err = errors.New("500 from Binance, its fault")
		break
	case 401:
		err = errors.New("binance api key not valid")
		break
	case 200:
		respBody, err = ioutil.ReadAll(resp.Body)
		break
	default:
		var response exchange.Binaresp
		if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
			break
		}
		err = fmt.Errorf("Binance return with code: %d - %s", resp.StatusCode, response.Msg)
	}
	if err != nil || len(respBody) == 0 || rand.Int()%10 == 0 {
		log.Printf("request to %s, got response from binance (error or throttled to 10%%): %s, err: %s", req.URL, common.TruncStr(respBody), common.ErrorToString(err))
	}
	return respBody, err
}

func (be *BinanceEndpoint) GetDepthOnePair(pair common.TokenPair) (exchange.Binaresp, error) {

	respBody, err := be.GetResponse(
		"GET", be.interf.PublicEndpoint()+"/api/v1/depth",
		map[string]string{
			"symbol": fmt.Sprintf("%s%s", pair.Base.ID, pair.Quote.ID),
			"limit":  "100",
		},
		false,
		common.GetTimepoint(),
	)

	respData := exchange.Binaresp{}
	if err != nil {
		return respData, err
	}
	if err = json.Unmarshal(respBody, &respData); err != nil {
		return respData, err
	}
	if respData.Code != 0 {
		return respData, fmt.Errorf("Getting depth from Binance failed: %s", respData.Msg)
	}
	return respData, nil
}

// Trade Relevant params:
// symbol ("%s%s", base, quote)
// side (BUY/SELL)
// type (LIMIT/MARKET)
// timeInForce (GTC/IOC)
// quantity
// price
//
// In this version, we only support LIMIT order which means only buy/sell with acceptable price,
// and GTC time in force which means that the order will be active until it's implicitly canceled
func (be *BinanceEndpoint) Trade(tradeType string, base, quote common.Token, rate, amount float64) (exchange.Binatrade, error) {
	result := exchange.Binatrade{}
	symbol := base.ID + quote.ID
	orderType := "LIMIT"
	params := map[string]string{
		"symbol":      symbol,
		"side":        strings.ToUpper(tradeType),
		"type":        orderType,
		"timeInForce": "GTC",
		"quantity":    strconv.FormatFloat(amount, 'f', -1, 64),
	}
	if orderType == "LIMIT" {
		params["price"] = strconv.FormatFloat(rate, 'f', -1, 64)
	}
	respBody, err := be.GetResponse(
		"POST",
		be.interf.AuthenticatedEndpoint()+"/api/v3/order",
		params,
		true,
		common.GetTimepoint(),
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(respBody, &result)
	return result, err
}

func (be *BinanceEndpoint) GetTradeHistory(symbol string) (exchange.BinanceTradeHistory, error) {
	result := exchange.BinanceTradeHistory{}
	timepoint := common.GetTimepoint()
	respBody, err := be.GetResponse(
		"GET",
		be.interf.PublicEndpoint()+"/api/v1/trades",
		map[string]string{
			"symbol": symbol,
			"limit":  "500",
		},
		false,
		timepoint,
	)
	if err == nil {
		err = json.Unmarshal(respBody, &result)
	}
	return result, err
}

func (be *BinanceEndpoint) GetAccountTradeHistory(
	base, quote common.Token,
	fromID string) (exchange.BinaAccountTradeHistory, error) {

	symbol := strings.ToUpper(fmt.Sprintf("%s%s", base.ID, quote.ID))
	result := exchange.BinaAccountTradeHistory{}
	params := map[string]string{
		"symbol": symbol,
		"limit":  "500",
	}
	if fromID != "" {
		params["fromId"] = fromID
	} else {
		params["fromId"] = "0"
	}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/api/v3/myTrades",
		params,
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		err = json.Unmarshal(respBody, &result)
	}
	return result, err
}

func (be *BinanceEndpoint) WithdrawHistory(startTime, endTime uint64) (exchange.Binawithdrawals, error) {
	result := exchange.Binawithdrawals{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/wapi/v3/withdrawHistory.html",
		map[string]string{
			"startTime": fmt.Sprintf("%d", startTime),
			"endTime":   fmt.Sprintf("%d", endTime),
		},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return result, err
		}
		if !result.Success {
			err = errors.New("Getting withdraw history from Binance failed: " + result.Msg)
		}
	}
	return result, err
}

func (be *BinanceEndpoint) DepositHistory(startTime, endTime uint64) (exchange.Binadeposits, error) {
	result := exchange.Binadeposits{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/wapi/v3/depositHistory.html",
		map[string]string{
			"startTime": fmt.Sprintf("%d", startTime),
			"endTime":   fmt.Sprintf("%d", endTime),
		},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return result, err
		}
		if !result.Success {
			err = errors.New("Getting deposit history from Binance failed: " + result.Msg)
		}
	}
	return result, err
}

func (be *BinanceEndpoint) CancelOrder(symbol string, id uint64) (exchange.Binacancel, error) {
	result := exchange.Binacancel{}
	respBody, err := be.GetResponse(
		"DELETE",
		be.interf.AuthenticatedEndpoint()+"/api/v3/order",
		map[string]string{
			"symbol":  symbol,
			"orderId": fmt.Sprintf("%d", id),
		},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return result, err
		}
		if result.Code != 0 {
			err = errors.New("Canceling order from Binance failed: " + result.Msg)
		}
	}
	return result, err
}

func (be *BinanceEndpoint) OrderStatus(symbol string, id uint64) (exchange.Binaorder, error) {
	result := exchange.Binaorder{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/api/v3/order",
		map[string]string{
			"symbol":  symbol,
			"orderId": fmt.Sprintf("%d", id),
		},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return result, err
		}
		if result.Code != 0 {
			err = errors.New(result.Msg)
		}
	}
	return result, err
}

func (be *BinanceEndpoint) Withdraw(token common.Token, amount *big.Int, address ethereum.Address) (string, error) {
	result := exchange.Binawithdraw{}
	respBody, err := be.GetResponse(
		"POST",
		be.interf.AuthenticatedEndpoint()+"/wapi/v3/withdraw.html",
		map[string]string{
			"asset":   token.ID,
			"address": address.Hex(),
			"name":    "reserve",
			"amount":  strconv.FormatFloat(common.BigToFloat(amount, token.Decimals), 'f', -1, 64),
		},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return "", err
		}
		if !result.Success {
			return "", errors.New(result.Msg)
		}
		return result.ID, nil
	}
	return "", fmt.Errorf("withdraw rejected by Binnace: %s", common.ErrorToString(err))
}

func (be *BinanceEndpoint) GetInfo() (exchange.Binainfo, error) {
	result := exchange.Binainfo{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/api/v3/account",
		map[string]string{},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return result, err
		}
	}
	if result.Code != 0 {
		return result, fmt.Errorf("Getting account info from Binance failed: %s", result.Msg)
	}
	return result, err
}

func (be *BinanceEndpoint) OpenOrdersForOnePair(pair common.TokenPair) (exchange.Binaorders, error) {

	result := exchange.Binaorders{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/api/v3/openOrders",
		map[string]string{
			"symbol": pair.Base.ID + pair.Quote.ID,
		},
		true,
		common.GetTimepoint(),
	)
	if err != nil {
		return result, err
	}
	if err = json.Unmarshal(respBody, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (be *BinanceEndpoint) GetDepositAddress(asset string) (exchange.Binadepositaddress, error) {
	result := exchange.Binadepositaddress{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.AuthenticatedEndpoint()+"/wapi/v3/depositAddress.html",
		map[string]string{
			"asset": asset,
		},
		true,
		common.GetTimepoint(),
	)
	if err == nil {
		if err = json.Unmarshal(respBody, &result); err != nil {
			return result, err
		}
		if !result.Success {
			err = errors.New(result.Msg)
		}
	}
	return result, err
}

func (be *BinanceEndpoint) GetExchangeInfo() (exchange.BinanceExchangeInfo, error) {
	result := exchange.BinanceExchangeInfo{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.PublicEndpoint()+"/api/v1/exchangeInfo",
		map[string]string{},
		false,
		common.GetTimepoint(),
	)
	if err == nil {
		err = json.Unmarshal(respBody, &result)
	}
	return result, err
}

func (be *BinanceEndpoint) getServerTime() (uint64, error) {
	result := exchange.BinaServerTime{}
	respBody, err := be.GetResponse(
		"GET",
		be.interf.PublicEndpoint()+"/api/v1/time",
		map[string]string{},
		false,
		common.GetTimepoint(),
	)
	if err == nil {
		err = json.Unmarshal(respBody, &result)
	}
	return result.ServerTime, err
}

func (be *BinanceEndpoint) UpdateTimeDelta() error {
	currentTime := common.GetTimepoint()
	serverTime, err := be.getServerTime()
	responseTime := common.GetTimepoint()
	if err != nil {
		return err
	}
	log.Printf("Binance current time: %d", currentTime)
	log.Printf("Binance server time: %d", serverTime)
	log.Printf("Binance response time: %d", responseTime)
	roundtripTime := (int64(responseTime) - int64(currentTime)) / 2
	be.timeDelta = int64(serverTime) - int64(currentTime) - roundtripTime

	log.Printf("Time delta: %d", be.timeDelta)
	return nil
}

//NewBinanceEndpoint return new endpoint instance for using binance
func NewBinanceEndpoint(signer Signer, interf Interface) *BinanceEndpoint {
	endpoint := &BinanceEndpoint{signer, interf, 0}
	switch interf.(type) {
	case *SimulatedInterface:
		log.Println("Simulate environment, no updateTime called...")
	default:
		err := endpoint.UpdateTimeDelta()
		if err != nil {
			panic(err)
		}
	}
	return endpoint
}
