# Rebalance activities

## Deposit to exchanges 

**signing required**

```shell
curl -X POST \
  http://localhost:8000/deposit/binance\
  -H 'content-type: multipart/form-data' \
  -F token=EOS \
  -F amount=0xde0b6b3a7640000
```

> Sample response:

```json
{
    "hash": "0x1b0c09f059904f1a9587641f2357c16c1c9fe43dfea161db31607f9221b0cfbb",
    "success": true
}
```

### HTTP Request

**POST**
`<host>:8000/deposit/:exchange_id`

### Form params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
amount | true | string | little endian hex string (must starts with 0x), eg: 0xde0b6b3a7640000
token | true |  string | token id string, eg: ETH, EOS...

## Withdraw from exchanges

**signing required**

```shell
curl -X POST \
  http://localhost:8000/withdraw/binance\
  -H 'content-type: multipart/form-data' \
  -F token=EOS \
  -F amount=0xde0b6b3a7640000
```

> Sample response:

```json
{
    "success": true
}
```

### HTTP Request

**POST**
`<host>:8000/withdraw/:exchange_id`


### Form params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
amount | true | string | little endian hex string (must starts with 0x), eg: 0xde0b6b3a7640000
token | true | string | token id string, eg: ETH, EOS...

## Setting rates 

**signing required**

```shell
curl -X POST \
  http://localhost:8000/setrates \
  -H 'content-type: multipart/form-data' \
  -F tokens=KNC-EOS \
  -F buys=0x5-0x7 \
  -F sells=0x5-0x7 \
  -F afp_mid=0x5-0x7 \
  -F block=2342353
```

### HTTP Request

**POST**
`<host>:8000/setrates`

### Form params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
tokens | true |  string | not including "ETH", represent all base token IDs separated by "-", eg: "ETH-ETH"
buys | true | string | represent all the buy (end users to buy tokens by ether) prices in little endian hex string, rates are separated by "-", eg: "0x5-0x7"
sells | true | string | represent all the sell (end users to sell tokens to ether) prices in little endian hex string, rates are separated by "-", eg: "0x5-0x7"
afp_mid | true | string |  represent all the afp mid (average filled price) in little endian hex string, rates are separated by "-", eg: "0x5-0x7" (this rate only stores in activities for tracking)
block | true | integer | in base 10, the block that prices are calculated on, eg: "3245876" means the prices are calculated from data at the time of block 3245876

## Trade 

**signing required**

```shell
curl -X POST \
  http://localhost:8000/trade/binance\
  -F base=ETH \
  -F quote=KNC \
  -F rate=300 \
  -F type=buy \
  -F amount=0.01
```

> Sample response:

```json
{
    "id": "19234634",
    "success": true,
    "done": 0,
    "remaining": 0.01,
    "finished": false
}
```

### HTTP Request

**POST**
`<host>:8000/trade/:exchange_id`

### Form params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
base | true | string | token id string, eg: ETH, EOS...
quote | true | string | token id string, eg: ETH, EOS...
amount | true | float | amount to trade
rate | true | float | rate for trade
type | true | string |  "buy" or "sell"

## Cancel order 

**signing required**

```shell
curl -X POST \
  http://localhost:8000/trade/binance\
  -F base=ETH \
  -F quote=KNC \
  -F order_id="1231701321"
```

> Sample response:

```json
{
    "reason": "UNKNOWN_ORDER",
    "success": false
}
```

### HTTP Request

**POST**
`<host>:8000/cancelorder/:exchange`

### Form params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
base | true | string | token id string, eg: ETH, EOS...
quote | true | string | token id string, eg: ETH, EOS...
order_id| true | string