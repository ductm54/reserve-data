# Authorized data

## Get trade history for an account

**signing required**

```shell
curl -X GET "http://localhost:8000/tradehistory?\
fromTime=1530403200000&\
toTime=1530489600000"
```

> sample response:

```json
{
    "data": {
        "Version": 1517298257114,
        "Valid": true,
        "Timestamp": "1517298257115",
        "Data": {
            "binance": {
                "EOS-ETH": [],
                "KNC-ETH": [
                    {
                        "ID": "548002",
                        "Price": 0.003038,
                        "Qty": 50,
                        "Type": "buy",
                        "Timestamp": 1516116380102
                    },
                    {
                        "ID": "548003",
                        "Price": 0.0030384,
                        "Qty": 7,
                        "Type": "buy",
                        "Timestamp": 1516116380102
                    },
                    {
                        "ID": "548004",
                        "Price": 0.003043,
                        "Qty": 16,
                        "Type": "buy",
                        "Timestamp": 1516116380102
                    },
                    {
                        "ID": "548005",
                        "Price": 0.0030604,
                        "Qty": 29,
                        "Type": "buy",
                        "Timestamp": 1516116380102
                    },
                    {
                        "ID": "548006",
                        "Price": 0.003065,
                        "Qty": 29,
                        "Type": "buy",
                        "Timestamp": 1516116380102
                    },
                    {
                        "ID": "548007",
                        "Price": 0.003065,
                        "Qty": 130,
                        "Type": "buy",
                        "Timestamp": 1516116380102
                    }
                ],
                "OMG-ETH": [
                    {
                        "ID": "123980",
                        "Price": 0.020473,
                        "Qty": 48,
                        "Type": "buy",
                        "Timestamp": 1512395498231
                    },
                    {
                        "ID": "130518",
                        "Price": 0.021022,
                        "Qty": 13.49,
                        "Type": "buy",
                        "Timestamp": 1512564108827
                    },
                    {
                        "ID": "130706",
                        "Price": 0.020202,
                        "Qty": 9.93,
                        "Type": "sell",
                        "Timestamp": 1512569059460
                    },
                    {
                        "ID": "140078",
                        "Price": 0.019098,
                        "Qty": 11.07,
                        "Type": "buy",
                        "Timestamp": 1512714826339
                    },
                    {
                        "ID": "140157",
                        "Price": 0.019053,
                        "Qty": 7.68,
                        "Type": "sell",
                        "Timestamp": 1512716338997
                    },
                    {
                        "ID": "295923",
                        "Price": 0.020446,
                        "Qty": 4,
                        "Type": "buy",
                        "Timestamp": 1514360742162
                    }
                ],
                "SALT-ETH": [],
                "SNT-ETH": []
            },
            "bittrex": {
                "OMG-ETH": [
                    {
                        "ID": "eb948865-6261-4991-8615-b36c8ccd1256",
                        "Price": 0.01822057,
                        "Qty": 1,
                        "Type": "buy",
                        "Timestamp": 18446737278344972745
                    }
                ],
                "SALT-ETH": [],
                "SNT-ETH": []
            }
        }
    },
    "success": true
}
```

### HTTP Request

`<host>:8000/tradehistory`

### URL Params

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | millisecond 
toTime | true | integer | millisecond 

<aside class="notice">Restriction: toTime - fromTime <= 3 days (in millisecond)</aside>



## Get exchange balances, reserve balances, pending activities at once

**signing required**

```shell
curl -X GET "http://localhost:8000/authdata"
```

> sample response:

```json
{
    "data": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114408227",
        "ReturnTime": "1514114408810",
        "ExchangeBalances": {
            "bittrex": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408226",
                "ReturnTime": "1514114408461",
                "AvailableBalance": {
                    "ETH": 0.10704306,
                    "OMG": 2.97381136
                },
                "LockedBalance": {
                    "ETH": 0,
                    "OMG": 0
                },
                "DepositBalance": {
                    "ETH": 0,
                    "OMG": 0
                }
            }
        },
        "ReserveBalances": {
            "ADX": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "BAT": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "CVC": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "DGD": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "EOS": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "ETH": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 360169992138038352
            },
            "FUN": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "GNT": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "KNC": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "LINK": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "MCO": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            },
            "OMG": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 23818094310417195708
            },
            "PAY": {
                "Valid": true,
                "Error": "",
                "Timestamp": "1514114408461",
                "ReturnTime": "1514114408799",
                "Balance": 0
            }
        },
        "PendingActivities": []
    },
    "block": 2345678,
    "success": true,
    "timestamp": "1514114409088",
    "version": 39
}
```

### HTTP Request

**GET**
`<host>:8000/authdata`

## Get all activities

```shell
curl -X GET "http://localhost:8000/activities?\
fromTime=1530489600000&\
toTime=1530576000000"
```

> sample response:

```json
{
    "data": {

    },
    "success": true
}
```

**signing required**

### HTTP Request

**GET**
`<host>:8000/activities`

### URL params: 

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | false | integer | from timepoint, unix millisecond
toTime | false | integer | to timepoint - uint64, unix millisecond

<aside class="notice">Restriction: toTime - fromTime <= 1 day (in millisecond)</aside>

## Get immediate pending activities

```shell
curl -X GET "http://localhost:8000/immediate-pending-activities"
```

> sample response:

```json
{
    "data": {},
    "success": true
}
```

**signing required**

### HTTP Request

**GET**
`<host>:8000/immediate-pending-activities`