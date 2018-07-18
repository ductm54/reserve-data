# Exchange Info

## Get precision and limit info 

for base-quote pair on an exchange

```shell
  curl -X GET "http://127.0.0.1:8000/exchangeinfo/binance/omg/eth"
```

> response:

```json
  {
    "data": {
        "Precision": {
            "Amount": 8,
            "Price": 8
        },
        "AmountLimit": {
            "Min": 0.01,
            "Max": 90000000
        },
        "PriceLimit": {
            "Min": 0.000001,
            "Max": 100000
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/exchangeinfo/<exchangeid>/<base>/<quote>`

Where **exchangeid** is the id of the exchange, **base** is symbol of the base token and **quote** is symbol of the quote token

## Get precision and limit info

for all base-quote pairs of an exchange

```shell
  curl -X GET "http://127.0.0.1:8000/exchangeinfo?exchangeid=binance"
```

> response:

```json
  {
    "data": {
        "binance": {
            "EOS-ETH": {
                "Precision": {
                    "Amount": 8,
                    "Price": 8
                },
                "AmountLimit": {
                    "Min": 0.01,
                    "Max": 90000000
                },
                "PriceLimit": {
                    "Min": 0.000001,
                    "Max": 100000
                },
                "MinNotional": 0.02
            },
            "KNC-ETH": {
                "Precision": {
                    "Amount": 8,
                    "Price": 8
                },
                "AmountLimit": {
                    "Min": 1,
                    "Max": 90000000
                },
                "PriceLimit": {
                    "Min": 1e-7,
                    "Max": 100000
                },
                "MinNotional": 0.02
            },
            "OMG-ETH": {
                "Precision": {
                    "Amount": 8,
                    "Price": 8
                },
                "AmountLimit": {
                    "Min": 0.01,
                    "Max": 90000000
                },
                "PriceLimit": {
                    "Min": 0.000001,
                    "Max": 100000
                },
                "MinNotional": 0.02
            },
            "SALT-ETH": {
                "Precision": {
                    "Amount": 8,
                    "Price": 8
                },
                "AmountLimit": {
                    "Min": 0.01,
                    "Max": 90000000
                },
                "PriceLimit": {
                    "Min": 0.000001,
                    "Max": 100000
                },
                "MinNotional": 0.02
            },
            "SNT-ETH": {
                "Precision": {
                    "Amount": 8,
                    "Price": 8
                },
                "AmountLimit": {
                    "Min": 1,
                    "Max": 90000000
                },
                "PriceLimit": {
                    "Min": 1e-8,
                    "Max": 100000
                },
                "MinNotional": 0.02
            }
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/exchangeinfo`

### Query paramameter:

Parameter | Required | Type | Description
--------  | -------- | ------- | ----------- 
exchangeid | false | string | id of exchange to get info (optional, if exchangeid is empty then return all exchanges info)

## Get fee for transaction on all exchanges

```shell
curl -X GET "http://127.0.0.1:8000/exchangefees"
```

> response:

```json
  {
    "data": [
        {
            "binance": {
                "Trading": {
                    "maker": 0.001,
                    "taker": 0.001
                },
                "Funding": {
                    "Withdraw": {
                        "EOS": 2,
                        "ETH": 0.005,
                        "FUN": 50,
                        "KNC": 1,
                        "LINK": 5,
                        "MCO": 0.15,
                        "OMG": 0.1
                    },
                    "Deposit": {
                        "EOS": 0,
                        "ETH": 0,
                        "FUN": 0,
                        "KNC": 0,
                        "LINK": 0,
                        "MCO": 0,
                        "OMG": 0
                    }
                }
            }
        },
        {
            "bittrex": {
                "Trading": {
                    "maker": 0.0025,
                    "taker": 0.0025
                },
                "Funding": {
                    "Withdraw": {
                        "BTC": 0.001,
                        "DASH": 0.002,
                        "DOGE": 2,
                        "FTC": 0.2,
                        "LTC": 0.01,
                        "NXT": 2,
                        "POT": 0.002,
                        "PPC": 0.02,
                        "RDD": 2,
                        "VTC": 0.02
                    },
                    "Deposit": {
                        "BTC": 0,
                        "DASH": 0,
                        "DOGE": 0,
                        "FTC": 0,
                        "LTC": 0,
                        "NXT": 0,
                        "POT": 0,
                        "PPC": 0,
                        "RDD": 0,
                        "VTC": 0
                    }
                }
            }
        }
    ],
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/exchangefees`


## Get fee for transaction on an exchange

```shell
curl -X GET "http://127.0.0.1:8000/exchangefees/binance"
```

> response:

```json
  {
    "data": {
        "Trading": {
            "maker": 0.001,
            "taker": 0.001
        },
        "Funding": {
            "Withdraw": {
                "EOS": 2,
                "ETH": 0.005,
                "FUN": 50,
                "KNC": 1,
                "LINK": 5,
                "MCO": 0.15,
                "OMG": 0.1
            },
            "Deposit": {
                "EOS": 0,
                "ETH": 0,
                "FUN": 0,
                "KNC": 0,
                "LINK": 0,
                "MCO": 0,
                "OMG": 0
            }
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/exchangefees/<exchangeid>`

Where **exchangeid** is the id of the exchange