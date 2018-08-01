# Settings APIS

## Set Token Update

> Example: This request will list token OMG and NEO. OMG is internal, NEO is external.

```shell
curl -X "POST" "http://localhost:8000/setting/set-token-update" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "data={
    \"OMG\": {
        \"token\": {
            \"id\": \"OMG\",
            \"name\": \"OmisexGO\",
            \"decimals\": 18,
            \"address\": \"0xd26114cd6EE289AccF82350c8d8487fedB8A0C07\",
            \"internal\": true,
            \"active\": true
        },
        \"exchanges\": {
            \"binance\": {
                \"deposit_address\": \"0x22222222222222222222222222222222222\",
                \"fee\": {
                    \"withdraw\": 0.2,
                    \"deposit\": 0.3
                },
                \"min_deposit\": 4
            }
        },
        \"pwis_equation\": {
            \"ask\": {
                \"a\": 800,
                \"b\": 600,
                \"c\": 0,
                \"min_min_spread\": 0,
                \"price_multiply_factor\": 0
            },
            \"bid\": {
                \"a\": 750,
                \"b\": 500,
                \"c\": 0,
                \"min_min_spread\": 0,
                \"price_multiply_factor\": 0
            }
        },
        \"target_qty\": {
            \"set_target\": {
                \"total_target\": 0,
                \"reserve_target\": 0,
                \"rebalance_threshold\": 0,
                \"transfer_threshold\": 0
            }
        },
        \"rebalance_quadratic\": {
            \"rebalance_quadratic\": {
                \"a\": 1,
                \"b\": 2,
                \"c\": 3
            }
        }
    },
    \"NEO\": {
        \"Token\": {
            \"id\": \"NEO\",
            \"name\": \"Request\",
            \"decimals\": 18,
            \"address\": \"0x8f8221afbb33998d8584a2b05749ba73c37a938a\",
            \"internal\": false,
            \"active\": true
        }
    }
}"
```

> Sample response

```json
    {
        "success": true
    }
```

**signing required** 

Prepare token update and store the request as pending

**Note**: 
- The API allow user to update token settings and its status. Hence can be used both for **list** and **delist** a token, as well as  to do minor modification for the token setting. 
To list a token, it active status is set to true. To delist a token, both its internal and active status is set to false.

- This data is in the form of a map tokenID:tokenUpdate which allows mutiple token updates at once

- It also allows mutiple requests, for example, one request update OMG, the other update KNC. Both these requests will be aggregate in to a list of token to be listed. These can be overwritten as well : if there are two requests update KNC, the later will overwite the ealier.  
  
- If a token is marked as internal, it will be required to come with exchange setting( fee, min deposit, exchange precision limit, deposit address) , and metric settings (pwis, targetQty). Since rebalance quadratic data can be zero value, it is optional. 
  
- If exchange precision limit (tokenUpdate.Exchange.Info) is null, It can be queried from exchange and set automatically for the pair (token-ETH). If this data is available in the request,it will be prioritize over the exchange queried data.
  
- In addition, if the update contain any Internal token, that token must be available in Smart contract in order to update its indices. 
  
- The tokenID from the map object will overwrite the token object's ID. Hence this token object ID inside the request is optional.

### HTTP Request

`GET http://example.com/setting/set-token-update`

## Get pending token update

```shell
curl -X "GET" "http://localhost:8000/setting/pending-token-update"
```

> Sample response:

```json
{
    "data": {
        "NEO": {
            "token": {
                "id": "NEO",
                "name": "Request",
                "address": "0x8f8221afbb33998d8584a2b05749ba73c37a938a",
                "decimals": 18,
                "active": true,
                "internal": false,
                "last_activation_change": 0
            },
            "exchanges": null,
            "pwis_equation": null,
            "target_qty": {
                "set_target": {
                    "total_target": 0,
                    "reserve_target": 0,
                    "rebalance_threshold": 0,
                    "transfer_threshold": 0
                }
            },
            "rebalance_quadratic": {
                "rebalance_quadratic": {
                    "a": 0,
                    "b": 0,
                    "c": 0
                }
            }
        },
        "OMG": {
            "token": {
                "id": "OMG",
                "name": "OmisexGO",
                "address": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
                "decimals": 18,
                "active": true,
                "internal": true,
                "last_activation_change": 0
            },
            "exchanges": {
                "binance": {
                    "deposit_address": "",
                    "exchange_info": {
                        "OMG-ETH": {
                            "precision": {
                                "amount": 2,
                                "price": 6
                            },
                            "amount_limit": {
                                "min": 0.01,
                                "max": 90000000
                            },
                            "price_limit": {
                                "min": 0.001611,
                                "max": 0.16103
                            },
                            "min_notional": 0.01
                        }
                    },
                    "fee": {
                        "withdraw": 0.2,
                        "deposit": 0.3
                    },
                    "min_deposit": 0
                }
            },
            "pwis_equation": {
                "ask": {
                    "a": 800,
                    "b": 600,
                    "c": 0,
                    "min_min_spread": 0,
                    "price_multiply_factor": 0
                },
                "bid": {
                    "a": 750,
                    "b": 500,
                    "c": 0,
                    "min_min_spread": 0,
                    "price_multiply_factor": 0
                }
            },
            "target_qty": {
                "set_target": {
                    "total_target": 1,
                    "reserve_target": 2,
                    "rebalance_threshold": 0,
                    "transfer_threshold": 0
                }
            },
            "rebalance_quadratic": {
                "rebalance_quadratic": {
                    "a": 1,
                    "b": 2,
                    "c": 3
                }
            }
        }
    }
}
```

**singing required** 

Return the current pending token updates information

### HTTP Request

`GET http://example.com/setting/pending-token-update`


## Confirm token update

```shell
curl -X "POST" "http://localhost:8000/setting/confirm-token-update" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "data={    
        \"NEO\": {
          \"token\": {
            \"id\": \"NEO\",
            \"name\": \"Request\",
            \"address\": \"0x8f8221afbb33998d8584a2b05749ba73c37a938a\",
            \"decimals\": 18,
            \"active\": true,
            \"internal\": false
          },
          \"exchanges\": null,
          \"pwis_equation\": null,
          \"target_qty\": {
            \"set_target\": {
              \"total_target\": 0,
              \"reserve_target\": 0,
              \"rebalance_threshold\": 0,
              \"transfer_threshold\": 0
            }
          },
          \"rebalance_quadratic\": {
            \"rebalance_quadratic\": {
              \"a\": 0,
              \"b\": 0,
              \"c\": 0
            }
          }
        },
        \"OMG\": {
          \"token\": {
            \"id\": \"OMG\",
            \"name\": \"OmisexGO\",
            \"address\": \"0xd26114cd6EE289AccF82350c8d8487fedB8A0C07\",
            \"decimals\": 18,
            \"active\": true,
            \"internal\": true
          },
          \"exchanges\": {
            \"binance\": {
              \"deposit_address\": \"0x22222222222222222222222222222222222\",
              \"exchange_info\": {
                \"OMG-ETH\": {
                  \"precision\": {
                    \"amount\": 2,
                    \"price\": 6
                  },
                  \"amount_limit\": {
                    \"min\": 0.01,
                    \"max\": 90000000
                  },
                  \"price_limit\": {
                    \"min\": 0.000001,
                    \"max\": 100000
                  },
                  \"min_notional\": 0.01
                }
              },
              \"fee\": {
                \"withdraw\": 0.2,
                \"deposit\": 0.3
              },
              \"min_deposit\": 4
            }
          },
          \"pwis_equation\": {
            \"ask\": {
              \"a\": 800,
              \"b\": 600,
              \"c\": 0,
              \"min_min_spread\": 0,
              \"price_multiply_factor\": 0
            },
            \"bid\": {
              \"a\": 750,
              \"b\": 500,
              \"c\": 0,
              \"min_min_spread\": 0,
              \"price_multiply_factor\": 0
            }
          },
          \"target_qty\": {
            \"set_target\": {
              \"total_target\": 0,
              \"reserve_target\": 0,
              \"rebalance_threshold\": 0,
              \"transfer_threshold\": 0
            }
          },
          \"rebalance_quadratic\": {
            \"rebalance_quadratic\": {
              \"a\": 0,
              \"b\": 0,
              \"c\": 0
            }
          }
        }
    }"
```

> Sample response

```json
{
    "success":true
}
```

**signing required**

Confirm token update and apply all the change to core.

### HTTP Request

`POST <host>:8000/setting/confirm-token-update`


## Reject pending token update

```shell
curl -X "POST" "http://localhost:8000/setting/reject-token-update" \
-H 'Content-Type: application/x-www-form-urlencoded'
```

> Sample response

```json
{
    "success":true
}
```

**signing required**

reject the update and remove the current pending update

### HTTP Request

`POST <host>:8000/setting/reject-token-update`

## Get Token settings

```shell
curl -X "GET" "http://localhost:8000/setting/token-settings"
```

> Sample response

```json
{
  "data": [
    {
      "id": "ABT",
      "name": "",
      "address": "0xb98d4c97425d9908e66e53a6fdf673acca0be986",
      "decimals": 18,
      "active": true,
      "internal": true
    }
  ],
  "success": true
}
```

**signing required**

get current token settings of core.

## Update address

```shell
curl -X "POST" "http://localhost:8000/setting/update-address" \
-H 'Content-Type: application/x-www-form-urlencoded'\
--data-urlencode "name=bank"\
--data-urlencode "address=0x123456789aabbcceeeddff"
```

> Sample response

```json
{
    "success": true
}
```

**signing required**

update a single address

### HTTP Request

`POST <host>:8000/setting/update-address`

### Post form

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
name | true | string | Name of the address (reserve, deposit etc...)
address | true | string | Hex form of the new address

## Add address to set 

```shell
curl -X "POST" "http://localhost:8000/setting/add-address-to-set" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "name=third_party_reserves" \
--data-urlencode "address=0x123456789aabbcceeeddff"
```

> sample response

```json
{
    "success": true
}
```

**signing required**

Add address to a list of address

## Update exchange fee

```shell
curl -X "POST" "http://localhost:8000/setting/update-exchange-fee" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "name=binance" \
--data-urlencode "data= {
      \"Trading\": {
        \"maker\": 0.001,
        \"taker\": 0.001
      },
      \"Funding\": {
        \"Withdraw\": {
          \"ZEC\": 0.005,
          \"ZIL\": 100,
          \"ZRX\": 5.8
        },
        \"Deposit\": {
          \"ZEC\": 0,
          \"ZIL\": 0,
          \"ZRX\": 2
        }
      }
    }"
```

> sample response

```json
{
    "success": true
}
```

**signing required**

Update one exchange fee setting

### HTTP Request

`POST <host>:8000/setting/update-exchange-fee`

### Post form parameters

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
name | true | string | Name of exchange
data | true | string | json format of fee config

<aside class="notice">
UpdateFee will merge the new fee setting to the current fee setting, Any different key will be overwriten from new fee to current fee. This allows update one single token's exchange fee on a destined exchange. UpdateFee will not be mutiplied by any value, so please prepare a big enough number to avoid exchange's fee increasing.
</aside>


## Update exchange mindeposit

```shell
curl -X "POST" "http://localhost:8000/setting/update-exchange-mindeposit" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "name=binance" \
--data-urlencode "data= {
      \"POWR\": 0.1,
      \"MANA\": 0.2
    }"
```

> sample response

```json
{
    "success": true
}
```

**signing required**

Update one exchange min deposit

### HTTP Request

`POST <host>:8000/setting/update-deposit-address`

### Post form parameter

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
name | true | string | Name of exchange
data | true | string | json format of deposit address config

<aside class="notice">
Update Exchange deposit address will merge the new deposit address setting to the current deposit address setting, Any different key will be overwriten from new deposit address to current deposit address. This allows update one single tokenpair's exchange precision limit on a destined exchange.
</aside>

## Update exchange info

```shell
curl -X "POST" "http://localhost:8000/setting/update-exchange-info" \
-H 'Content-Type: application/x-www-form-urlencoded'\
--data-urlencode "name=binance"\
--data-urlencode "data= {
      \"LINK-ETH\": {
        \"precision\": {
          \"amount\": 0,
          \"price\": 8
        },
        \"amount_limit\": {
          \"min\": 1,
          \"max\": 90000000
        },
        \"price_limit\": {
          \"min\": 1e-8,
          \"max\": 120000
        },
        \"min_notional\": 0.01
      }
    }"
```

> sample response

```json
{
    "success": true
}
```

**signing required**

Update one exchange's info

### HTTP Request

`POST <host>:8000/setting/update-exchange-info`

### Post form parameters

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
name | true | string | Name of exchange
data | true | string | json format of exchange info

<aside class="notice">
Update Exchange minDeposit will merge the new exchange info setting to the current exchange info setting, Any different key will be overwriten from new exchange info to current exchange info. This allows update one single token's exchange minDeposit on a destined exchange.
</aside>


## Get all settings

```shell
curl -X "GET" "http://localhost:8000/setting/all-settings"
```

> sample response

```json
{
  "data": {
    "Addresses": {
      "bank": "",
      "burner": "0x07f6e905f2a1559cd9fd43cb92f8a1062a3ca706",
      "network": "0x964f35fae36d75b1e72770e244f6595b68508cf5",
      "old_burners": [
        "0x4e89bc8484b2c454f2f7b25b612b648c45e14a8e"
      ],
      "pricing": "0x798abda6cc246d0edba912092a2a3dbd3d11191b",
      "reserve": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
      "setrate": "",
      "third_party_reserves": [
        "0x2aab2b157a03915c8a73adae735d0cf51c872f31",
        "0x4d864b5b4f866f65f53cbaad32eb9574760865e6",
        "0x6f50e41885fdc44dbdf7797df0393779a9c0a3a6"
      ],
      "whitelist": "0x6e106a75d369d09a9ea1dcc16da844792aa669a3",
      "wrapper": "0x6172afc8c00c46e0d07ce3af203828198194620a"
    },
    "Tokens": [
      {
        "id": "ABT",
        "name": "",
        "address": "0xb98d4c97425d9908e66e53a6fdf673acca0be986",
        "decimals": 18,
        "active": true,
        "internal": true
      },
      {
        "id": "ZIL",
        "name": "",
        "address": "0x05f4a42e251f2d52b8ed15e9fedaacfcef1fad27",
        "decimals": 12,
        "active": true,
        "internal": true
      }
    ],
    "Exchanges": {
      "binance": {
        "echange_addresses": {
          "AE": "0x44d34a119ba21a42167ff8b77a88f0fc7bb2db90",
          "ZIL": "0x44d34a119ba21a42167ff8b77a88f0fc7bb2db90"
        },
        "min_deposit": {
          "YOYO": 0,
          "ZEC": 0,
          "ZIL": 0,
          "ZRX": 0
        },
        "fee": {
          "Trading": {
            "maker": 0.001,
            "taker": 0.001
          },
          "Funding": {
            "Withdraw": {
              "ZEC": 0.005,
              "ZIL": 100,
              "ZRX": 5.8
            },
            "Deposit": {
              "ZIL": 0,
              "ZRX": 2
            }
          }
        },
        "info": {
          "AE-ETH": {
            "precision": {
              "amount": 2,
              "price": 6
            },
            "amount_limit": {
              "min": 0.01,
              "max": 90000000
            },
            "price_limit": {
              "min": 0.000001,
              "max": 100000
            },
            "min_notional": 0.01
          },
          "AION-ETH": {
            "precision": {
              "amount": 2,
              "price": 6
            },
            "amount_limit": {
              "min": 0.01,
              "max": 90000000
            },
            "price_limit": {
              "min": 0.000001,
              "max": 100000
            },
            "min_notional": 0.01
          }
        }
      },
      "huobi": {
        "deposit_address": {
          "ABT": "0x0c8fd73eaf6089ef1b91231d0a07d0d2ca2b9d66",
          "CVC": "0x0c8fd73eaf6089ef1b91231d0a07d0d2ca2b9d66",
          "EDU": "0x0c8fd73eaf6089ef1b91231d0a07d0d2ca2b9d66",
         
        },
        "min_deposit": {
          "ABT": 2,
          "APPC": 0.5,
          "AST": 5,
          "SNT": 50,
          "ZIL": 100
        },
        "fee": {
          "Trading": {
            "maker": 0.002,
            "taker": 0.002
          },
          "Funding": {
            "Withdraw": {
              "ABT": 2,
              "ZRX": 5
            },
            "Deposit": {
              "ABT": 0,
              "ZRX": 0
            }
          }
        },
        "info": {
          "POLY-ETH": {
            "precision": {
              "amount": 4,
              "price": 6
            },
            "amount_limit": {
              "min": 0,
              "max": 0
            },
            "price_limit": {
              "min": 0,
              "max": 0
            },
            "min_notional": 0.02
          }
        }
      }
    }
  },
  "success": true
}
```

**signing required**

Return all current running setting of core

### HTTP Request

`GET <host>:8000/setting/all-settings`