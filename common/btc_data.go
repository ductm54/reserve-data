package common

// BitfinexData is the response of Bitfinix Ticker request.
// Example:
//
// Request: https://api.bitfinex.com/v1/pubticker/btcusd
// Response:
// {
// "mid":"244.755",
// "bid":"244.75",
// "ask":"244.76",
// "last_price":"244.82",
// "low":"244.2",
// "high":"248.19",
// "volume":"7842.11542563",
// "timestamp":"1444253422.348340958"
//}
type BitfinexData struct {
	Valid     bool
	Error     string
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

// BinanceData is the response of Binance Ticker request.
// Example:
// Request: https://api.binance.com/api/v3/ticker/bookTicker?symbol=ETHBTC
// Response: {
//  "symbol": "ETHBTC",
//  "bidPrice": "0.03338700",
//  "bidQty": "3.39200000",
//  "askPrice": "0.03339400",
//  "askQty": "0.08600000"
//}
type BinanceData struct {
	Valid    bool
	Error    string
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

// BTCData is the data returned by /btc-feed API.
type BTCData struct {
	Timestamp uint64
	Bitfinex  BitfinexData `json:"bitfinex"`
	Binance   BinanceData  `json:"binance"`
}
