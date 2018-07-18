# Stats

## Get trade logs

```shell
    curl -X GET "http://127.0.0.1:8000/tradelogs"
```

> sample response

```json
{
    "data": [
        {
            "Timestamp": 1531872104000000000,
            "BlockNumber": 5983404,
            "TransactionHash": "0x0c2e554e51dc83b9242ac23474e51383b6b256a503ea8c4d35a8e17060e35739",
            "Index": 18,
            "EtherReceivalSender": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
            "EtherReceivalAmount": 5000050000000000000,
            "UserAddress": "0x85c5c26dc2af5546341fc1988b9d178148b4838b",
            "SrcAddress": "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
            "DestAddress": "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
            "SrcAmount": 7.000192905074535e+21,
            "DestAmount": 5000050000000000000,
            "FiatAmount": 2505.05005025,
            "ReserveAddress": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
            "WalletAddress": "0x0000000000000000000000000000000000000000",
            "WalletFee": null,
            "BurnFee": 5550055500000000000,
            "IP": "",
            "Country": "unknown"
        },
        {
            "Timestamp": 1531872126000000000,
            "BlockNumber": 5983406,
            "TransactionHash": "0x58ee79a11a28145191441e31dd6e8c6e9d777ca239af41591cc2b73d9a71e9f7",
            "Index": 39,
            "EtherReceivalSender": "0x21433dec9cb634a23c6a4bbcce08c83f5ac2ec18",
            "EtherReceivalAmount": 2000020000000000000,
            "UserAddress": "0x85c5c26dc2af5546341fc1988b9d178148b4838b",
            "SrcAddress": "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
            "DestAddress": "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
            "SrcAmount": 2.797241514941635e+21,
            "DestAmount": 2000020000000000000,
            "FiatAmount": 1002.0200201,
            "ReserveAddress": "0x21433dec9cb634a23c6a4bbcce08c83f5ac2ec18",
            "WalletAddress": "0x0000000000000000000000000000000000000000",
            "WalletFee": null,
            "BurnFee": 1065610656000000000,
            "IP": "",
            "Country": "unknown"
        },
        {
            "Timestamp": 1531872126000000000,
            "BlockNumber": 5983406,
            "TransactionHash": "0x1060bacda8800c3fc67999131a41a9c89ce12031269bd77a764e8591973ffe59",
            "Index": 50,
            "EtherReceivalSender": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
            "EtherReceivalAmount": 3500318029638611500,
            "UserAddress": "0x85c5c26dc2af5546341fc1988b9d178148b4838b",
            "SrcAddress": "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
            "DestAddress": "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
            "SrcAmount": 4.913845263147823e+21,
            "DestAmount": 3500318029638611500,
            "FiatAmount": 1753.6768344390925,
            "ReserveAddress": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
            "WalletAddress": "0x0000000000000000000000000000000000000000",
            "WalletFee": null,
            "BurnFee": 3885353012898858500,
            "IP": "",
            "Country": "unknown"
        },
        {
            "Timestamp": 1531872126000000000,
            "BlockNumber": 5983406,
            "TransactionHash": "0x0ec26ce906a1718d6163231967da11c36cfb68f3c72faef276a970f952b52dbe",
            "Index": 61,
            "EtherReceivalSender": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
            "EtherReceivalAmount": 3496091443166923300,
            "UserAddress": "0x85c5c26dc2af5546341fc1988b9d178148b4838b",
            "SrcAddress": "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
            "DestAddress": "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
            "SrcAmount": 4.913845263147823e+21,
            "DestAmount": 3496091443166923300,
            "FiatAmount": 1751.5592934838444,
            "ReserveAddress": "0x63825c174ab367968ec60f061753d3bbd36a0d8f",
            "WalletAddress": "0x0000000000000000000000000000000000000000",
            "WalletFee": null,
            "BurnFee": 3880661501915284500,
            "IP": "",
            "Country": "unknown"
        }
    ]
}
```

### HTTP Request

**GET**
`<host>:8000/tradelogs`


## Get asset volume 

for aggregate time (hour, day, month)

```shell
curl -x GET http://localhost:8000/get-asset-volume?\
    fromTime=1520640035000&\
    toTime=1520722835000&\
    asset=eth&\
    freq=M
```

> sample response

```json
{
    "data": {
        "1520652360000": {
            "usd_amount": 0.734518,
            "volume": 0.001
        },
        "1520654280000": {
            "usd_amount": 0.7297319999999999,
            "volume": 0.001
        },
        "1520654820000": {
            "usd_amount": 1.4581552500230603,
            "volume": 0.001998206533389053
        },
        "1520656440000": {
            "usd_amount": 0.7297319999999999,
            "volume": 0.001
        },
        "1520656500000": {
            "usd_amount": 0.7297319999999999,
            "volume": 0.001
        },
        "1520656560000": {
            "usd_amount": 0.7297319999999999,
            "volume": 0.001
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-asset-volume`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from time stamp (millisecond)
toTime | true | integer | required: to time stamp (millisecond)
freq | true | string | frequency to get data (H/D/M)
asset | true | string | asset name (eg: ETH)

## Get reserve volume

```shell
curl -x GET \
    http://localhost:8000/get-reserve-volume?\
    fromTime=1522540800000&\
    toTime=1522627200000&\
    freq=D&token=KNC&\
    reserveAddr=0x63825c174ab367968EC60f061753D3bbD36A0D8F
```

> sample response:

```json
{
    "data": {
        "1522540800000": {
            "eth_amount": 9.971150530912206,
            "usd_amount": 3838.6105908493496,
            "volume": 3945.5899585215247
        },
        "1522627200000": {
            "eth_amount": 14.749439804645423,
            "usd_amount": 5766.650333669346,
            "volume": 5884.90733954939
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/exchange-notifications`

### URL Params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | millisecond
toTime | true | integer | millisecond
token | true | string | name of token to get volume (eg: ETH)
reserveAddr | true | string | reserve address to get volume of token
freq | true | string | frequency to get volume ("M", "H", "D" - Minute, Hour, Day)


## Get burn fee

follow aggregate time (hour, day, month)

```shell
    curl -X GET http://localhost:8000/get-burn-fee?\
    fromTime=1520640035000&\
    toTime=1520722835000&\
    reserveAddr=0x2c5a182d280eeb5824377b98cd74871f78d6b8bc&\
    freq=H
```

> sample response

```json
{
    "data": {
        "1520650800000": 0.00225,
        "1520654400000": 0.005622982350062684
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-burn-fee`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from time stamp (millisecond)
toTime | true | integer | to time stamp (millisecond)
freq | true | string | frequency to get data (H/D/M)
reserveAddr | true | string | reserve address to get burn fee


## Get wallet fee

follow aggregate time (hour, day, month)

```shell
    curl -X GET http://localhost:8000/get-burn-fee?\
    fromTime=1520640035000&\
    toTime=1520722835000&\
    reserveAddr=0x2c5a182d280eeb5824377b98cd74871f78d6b8bc&\
    freq=H
```

> sample response

```json
{
    "data": {
        "1520650800000": 0.00225,
        "1520654400000": 0.005622982350062684
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-wallet-fee`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from time stamp (millisecond)
toTime | true | integer | to time stamp (millisecond)
freq | true | integer | frequency to get data (H/D/M)
reserveAddr | true | integer | reserve address
walletAddr | true | integer | wallet address to get fee


## Get user volume

follow aggregate time (hour, day, month)

```shell
    curl -X GET http://localhost:8000/get-burn-fee?\
    fromTime=1520640035000&\
    toTime=1520722835000&\
    reserveAddr=0x2c5a182d280eeb5824377b98cd74871f78d6b8bc&\
    freq=H
```

> sample response

```json
{
    "data": {
        "1520650800000": 0.00225,
        "1520654400000": 0.005622982350062684
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-user-volume`

### URL params:
Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from time stamp (millisecond)
toTime | true | integer | to time stamp (millisecond)
freq | true | string | frequency to get data (H/D/M)
userAddr | true | string | user address to get volume

## Get rate from blockchain 

follow reserve (including sanity rate)

```shell
    curl -x GET http://localhost:8000/get-reserve-rate?\
    fromTime=1520650426000&\
    reserveAddr=0x2C5a182d280EeB5824377B98CD74871f78d6b8BC
```

> sample response

```json
{
    "data": [
        {
            "Timestamp": 0,
            "ReturnTime": 1520655211398,
            "BlockNumber": 5228238,
            "Data": {
                "APPC-ETH": {
                    "ReserveRate": 0.008393501685222925,
                    "SanityRate": 0.009476954807692308
                },
                "BAT-ETH": {
                    "ReserveRate": 0.004239837479770336,
                    "SanityRate": 0.004723026
                },
                "BQX-ETH": {
                    "ReserveRate": 0.000584106942517358,
                    "SanityRate": 0.000652623583333333
                },
                "ELF-ETH": {
                    "ReserveRate": 0.000111035861616385,
                    "SanityRate": 0.000123576933333333
                },
                "ENG-ETH": {
                    "ReserveRate": 0.000596961062855617,
                    "SanityRate": 0.000671348333333333
                },
                "EOS-ETH": {
                    "ReserveRate": 0.002752586323625439,
                    "SanityRate": 0.0029518775
                },
                "ETH-APPC": {
                    "ReserveRate": 117.09352189952426,
                    "SanityRate": 127.67814393478591
                },
                "ETH-BAT": {
                    "ReserveRate": 229.70817443088293,
                    "SanityRate": 256.19168727845243
                },
                "ETH-BQX": {
                    "ReserveRate": 1674.1030165099337,
                    "SanityRate": 1854.0549727299406
                },
                "ETH-ELF": {
                    "ReserveRate": 8741.854159397268,
                    "SanityRate": 9791.471331758756
                },
                "ETH-ENG": {
                    "ReserveRate": 1615.634600168794,
                    "SanityRate": 1802.3430459597466
                },
                "ETH-EOS": {
                    "ReserveRate": 356.63560217559376,
                    "SanityRate": 409.9086090123997
                },
                "ETH-GTO": {
                    "ReserveRate": 377.41338205276884,
                    "SanityRate": 432.7020247561754
                },
                "ETH-KNC": {
                    "ReserveRate": 3343.445798727388,
                    "SanityRate": 3740.5264791019335
                },
                "ETH-MANA": {
                    "ReserveRate": 2653.110602592891,
                    "SanityRate": 2961.048749629869
                },
                "ETH-OMG": {
                    "ReserveRate": 221.03211631662654,
                    "SanityRate": 247.71892159696137
                },
                "ETH-POWR": {
                    "ReserveRate": 2625.724091042635,
                    "SanityRate": 2849.3656923419003
                },
                "ETH-RDN": {
                    "ReserveRate": 49.46371879742714,
                    "SanityRate": 54.91076347236177
                },
                "ETH-REQ": {
                    "ReserveRate": 5123.294220111987,
                    "SanityRate": 5665.576472406068
                },
                "ETH-SALT": {
                    "ReserveRate": 532.7984920557698,
                    "SanityRate": 611.1450636146453
                },
                "ETH-SNT": {
                    "ReserveRate": 924.3533982883454,
                    "SanityRate": 1052.547224086108
                },
                "GTO-ETH": {
                    "ReserveRate": 0.002590177384056749,
                    "SanityRate": 0.002796381645502645
                },
                "KNC-ETH": {
                    "ReserveRate": 0.0002894500017402,
                    "SanityRate": 0.000323483875
                },
                "MANA-ETH": {
                    "ReserveRate": 0.00036813437957934,
                    "SanityRate": 0.000408639
                },
                "OMG-ETH": {
                    "ReserveRate": 0.004383746019560721,
                    "SanityRate": 0.004884568333333332
                },
                "POWR-ETH": {
                    "ReserveRate": 0.000369936605210205,
                    "SanityRate": 0.000424655916666666
                },
                "RDN-ETH": {
                    "ReserveRate": 0.01987031936393942,
                    "SanityRate": 0.02203575261904761
                },
                "REQ-ETH": {
                    "ReserveRate": 0.000191920526182855,
                    "SanityRate": 0.0002135705
                },
                "SALT-ETH": {
                    "ReserveRate": 0.001821407074188612,
                    "SanityRate": 0.00197989
                },
                "SNT-ETH": {
                    "ReserveRate": 0.001042608668464954,
                    "SanityRate": 0.001149592125
                }
            }
        },
        {
            "Timestamp": 0,
            "ReturnTime": 1520655227886,
            "BlockNumber": 5228239,
            "Data": {
                "APPC-ETH": {
                    "ReserveRate": 0.000369936605210205,
                    "SanityRate": 0.000424655916666666
                },
                "BAT-ETH": {
                    "ReserveRate": 0.0002894500017402,
                    "SanityRate": 0.000323483875
                },
                "BQX-ETH": {
                    "ReserveRate": 0.002590177384056749,
                    "SanityRate": 0.002796381645502645
                },
                "ELF-ETH": {
                    "ReserveRate": 0.01987031936393942,
                    "SanityRate": 0.02203575261904761
                },
                "ENG-ETH": {
                    "ReserveRate": 0.000584106942517358,
                    "SanityRate": 0.000652623583333333
                },
                "EOS-ETH": {
                    "ReserveRate": 0.000191920526182855,
                    "SanityRate": 0.0002135705
                },
                "ETH-APPC": {
                    "ReserveRate": 2625.724091042635,
                    "SanityRate": 2849.3656923419003
                },
                "ETH-BAT": {
                    "ReserveRate": 3343.445798727388,
                    "SanityRate": 3740.5264791019335
                },
                "ETH-BQX": {
                    "ReserveRate": 377.41338205276884,
                    "SanityRate": 432.7020247561754
                },
                "ETH-ELF": {
                    "ReserveRate": 49.46371879742714,
                    "SanityRate": 54.91076347236177
                },
                "ETH-ENG": {
                    "ReserveRate": 1674.1030165099337,
                    "SanityRate": 1854.0549727299406
                },
                "ETH-EOS": {
                    "ReserveRate": 5123.294220111987,
                    "SanityRate": 5665.576472406068
                },
                "ETH-GTO": {
                    "ReserveRate": 229.70817443088293,
                    "SanityRate": 256.19168727845243
                },
                "ETH-KNC": {
                    "ReserveRate": 1615.634600168794,
                    "SanityRate": 1802.3430459597466
                },
                "ETH-MANA": {
                    "ReserveRate": 221.03211631662654,
                    "SanityRate": 247.71892159696137
                },
                "ETH-OMG": {
                    "ReserveRate": 8741.854159397268,
                    "SanityRate": 9791.471331758756
                },
                "ETH-POWR": {
                    "ReserveRate": 924.3533982883454,
                    "SanityRate": 1052.547224086108
                },
                "ETH-RDN": {
                    "ReserveRate": 532.7984920557698,
                    "SanityRate": 611.1450636146453
                },
                "ETH-REQ": {
                    "ReserveRate": 117.09352189952426,
                    "SanityRate": 127.67814393478591
                },
                "ETH-SALT": {
                    "ReserveRate": 2653.110602592891,
                    "SanityRate": 2961.048749629869
                },
                "ETH-SNT": {
                    "ReserveRate": 356.63560217559376,
                    "SanityRate": 409.9086090123997
                },
                "GTO-ETH": {
                    "ReserveRate": 0.004239837479770336,
                    "SanityRate": 0.004723026
                },
                "KNC-ETH": {
                    "ReserveRate": 0.000596961062855617,
                    "SanityRate": 0.000671348333333333
                },
                "MANA-ETH": {
                    "ReserveRate": 0.004383746019560721,
                    "SanityRate": 0.004884568333333332
                },
                "OMG-ETH": {
                    "ReserveRate": 0.000111035861616385,
                    "SanityRate": 0.000123576933333333
                },
                "POWR-ETH": {
                    "ReserveRate": 0.001042608668464954,
                    "SanityRate": 0.001149592125
                },
                "RDN-ETH": {
                    "ReserveRate": 0.001821407074188612,
                    "SanityRate": 0.00197989
                },
                "REQ-ETH": {
                    "ReserveRate": 0.008393501685222925,
                    "SanityRate": 0.009476954807692308
                },
                "SALT-ETH": {
                    "ReserveRate": 0.00036813437957934,
                    "SanityRate": 0.000408639
                },
                "SNT-ETH": {
                    "ReserveRate": 0.002752586323625439,
                    "SanityRate": 0.0029518775
                }
            }
        }
    ],
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-reserve-rate`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from timestamp (millisecond)
toTime | true | integer | to timestamp (millisecond)
reserveAddr | true | string | Address of the reserve to get rate from


## Get trade summary

follow timeframe (day)

```shell
    curl -x GET http://localhost:8000/get-trade-summary?\
    fromTime=1519297149000&\
    toTime=1519815549000
```

> sample response

```json
{
    "data": {
        "1519344000000": {
            "eth_per_trade": 0.55402703087424,
            "kyced_addresses": 0,
            "new_unique_addresses": 35,
            "total_burn_fee": 0,
            "total_eth_volume": 44.3221624699392,
            "total_trade": 80,
            "total_usd_amount": 30981.281202536768,
            "unique_addresses": 50,
            "usd_per_trade": 387.26601503170957
        },
        "1519430400000": {
            "eth_per_trade": 0.17008867987348247,
            "kyced_addresses": 0,
            "new_unique_addresses": 17,
            "total_burn_fee": 0,
            "total_eth_volume": 8.674522673547607,
            "total_trade": 51,
            "total_usd_amount": 6060.828270348999,
            "unique_addresses": 29,
            "usd_per_trade": 118.83977000684311
        },
        "1519516800000": {
            "eth_per_trade": 0.14234886960871,
            "kyced_addresses": 0,
            "new_unique_addresses": 9,
            "total_burn_fee": 1.1025,
            "total_eth_volume": 5.40925704513098,
            "total_trade": 38,
            "total_usd_amount": 3779.4100326337,
            "unique_addresses": 18,
            "usd_per_trade": 99.45815875351843
        },
        "1519603200000": {
            "eth_per_trade": 0.5430574166436676,
            "kyced_addresses": 0,
            "new_unique_addresses": 39,
            "total_burn_fee": 42.85336706164196,
            "total_eth_volume": 45.07376558142441,
            "total_trade": 83,
            "total_usd_amount": 31497.3427579499,
            "unique_addresses": 56,
            "usd_per_trade": 379.4860573246976
        },
        "1519689600000": {
            "eth_per_trade": 0.6014134385918366,
            "kyced_addresses": 0,
            "new_unique_addresses": 69,
            "total_burn_fee": 79.03472646631772,
            "total_eth_volume": 78.7851604555306,
            "total_trade": 131,
            "total_usd_amount": 55076.026979006005,
            "unique_addresses": 92,
            "usd_per_trade": 420.4276868626413
        },
        "1519776000000": {
            "eth_per_trade": 0.40083191776618454,
            "kyced_addresses": 0,
            "new_unique_addresses": 64,
            "total_burn_fee": 48.899026261678536,
            "total_eth_volume": 52.50898122737018,
            "total_trade": 131,
            "total_usd_amount": 36662.138255818456,
            "unique_addresses": 94,
            "usd_per_trade": 279.8636508077745
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-trade-summary`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from time stamp (millisecond)
toTime | true | integer | to time stamp (millisecond)
timeZone | true | integer | (in range [-12,14], default to 0): the integer specific which UTC timezone to query

## Get wallet stats summary

follow timeframe (day)

```shell
    curl -x GET http://localhost:8000/get-wallet-stats?\
    fromTime=1521914061000&\
    toTime=1523000461000&\
    walletAddr=0xb9e29984fe50602e7a619662ebed4f90d93824c7
```

> sample response

```json
{
    "data": {
        "1521936000000": {
            "eth_per_trade": 0.15169175185997197,
            "kyced_addresses": 0,
            "new_unique_addresses": 27,
            "total_burn_fee": 3.5843774403434443,
            "total_eth_volume": 9.101505111598318,
            "total_trade": 60,
            "total_usd_amount": 4738.284168671162,
            "unique_addresses": 40,
            "usd_per_trade": 78.97140281118602
        },
        "1522022400000": {
            "eth_per_trade": 0.1305336778977258,
            "kyced_addresses": 0,
            "new_unique_addresses": 13,
            "total_burn_fee": 1.2758795269915402,
            "total_eth_volume": 2.3496062021590642,
            "total_trade": 18,
            "total_usd_amount": 1230.3892752776494,
            "unique_addresses": 18,
            "usd_per_trade": 68.35495973764719
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-wallet-stats`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from time stamp (millisecond)
toTime | true |  integer | to time stamp (millisecond)
timeZone | false (default 0)| integer | in range [-12,14] the integer specific which UTC timezone to query
walletAddr | true | string | to specific which wallet Address to query data from. It must be larger than 2^128 to be valid.

## Get wallet list 

```shell
    curl -x GET http://localhost:8000/get-wallet-address
```

> sample response

```json
{
    "data": [
        "0xb9e29984fe50602e7a619662ebed4f90d93824c7",
        "0xf1aa99c69715f423086008eb9d06dc1e35cc504d"
    ],
    "success": true
}
```

Return a list of wallet address that has ever traded with core

### HTTP Request

**GET**
`<host>:8000/get-wallet-address`

## Get exchanges status

```shell
    curl -x GET http://localhost:8000/get-exchange-status
```

> sample response:

```json
{
    "data": {
        "binance": {
            "timestamp": 1521532176702,
            "status": true
        },
        "bittrex": {
            "timestamp": 1521532176704,
            "status": true
        },
        "huobi": {
            "timestamp": 1521532176703,
            "status": true
        }
    },
    "success": true
}
```

###HTTP Request

**GET**
`<host>:8000/get-exchange-status`

## Get country stats

> sample response: 

```json
{
    "data": {
        "1522368000000": {
            "eth_per_trade": 1.1759348083481784,
            "kyced_addresses": 0,
            "new_unique_addresses": 23,
            "total_burn_fee": 40.10625390027786,
            "total_eth_volume": 51.741131567319854,
            "total_trade": 44,
            "total_usd_amount": 19804.392524011764,
            "unique_addresses": 26,
            "usd_per_trade": 450.09983009117644
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-country-stats`

### URL params:
Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from timestamp (millisecond)
toTime | true | integer | to timestamp (millisecond)
country | true | string |  internatinal country 
timezone | true | integer | timezone to get country stats from -11 to 14

## Get heatmap

Return list of countries sort by total ETH value

> sample response:

```json
{
    "data": [
        {
            "country": "US",
            "total_eth_value": 51.741131567319854,
            "total_fiat_value": 19804.392524011764
        },
        {
            "country": "unknown",
            "total_eth_value": 31.28130484378119,
            "total_fiat_value": 12268.937507634406
        },
        {
            "country": "TW",
            "total_eth_value": 15,
            "total_fiat_value": 5916.6900000000005
        },
        {
            "country": "KR",
            "total_eth_value": 13.280037553077175,
            "total_fiat_value": 5016.70456645198
        },
        {
            "country": "JP",
            "total_eth_value": 10.277090646,
            "total_fiat_value": 3857.271305900826
        },
        {
            "country": "TH",
            "total_eth_value": 8.241091466923997,
            "total_fiat_value": 3195.368602817533
        },
        {
            "country": "CA",
            "total_eth_value": 3.8122812821017558,
            "total_fiat_value": 1445.8819158742285
        },
        {
            "country": "AU",
            "total_eth_value": 2.6,
            "total_fiat_value": 969.02
        },
        {
            "country": "DE",
            "total_eth_value": 1.823287,
            "total_fiat_value": 697.502009413
        },
        {
            "country": "ID",
            "total_eth_value": 1.7178731840736186,
            "total_fiat_value": 674.8439050493492
        },
        {
            "country": "RO",
            "total_eth_value": 1.4009999999999998,
            "total_fiat_value": 529.075415
        },
        {
            "country": "VN",
            "total_eth_value": 1.3951777988339262,
            "total_fiat_value": 548.8376078547749
        },
        {
            "country": "CN",
            "total_eth_value": 1.0121575386522288,
            "total_fiat_value": 401.6824093511598
        },
        {
            "country": "PL",
            "total_eth_value": 0.379699,
            "total_fiat_value": 144.141714079
        },
        {
            "country": "FR",
            "total_eth_value": 0.319624,
            "total_fiat_value": 122.92586391999998
        },
        {
            "country": "SG",
            "total_eth_value": 0.15642985716526572,
            "total_fiat_value": 64.06928945889221
        },
        {
            "country": "ES",
            "total_eth_value": 0.09344946,
            "total_fiat_value": 35.176806429959996
        },
        {
            "country": "XX",
            "total_eth_value": 0.09,
            "total_fiat_value": 36.86148
        },
        {
            "country": "IN",
            "total_eth_value": 0.0714026952146661,
            "total_fiat_value": 27.977050948875906
        },
        {
            "country": "AR",
            "total_eth_value": 0.02751473,
            "total_fiat_value": 10.92519129691
        },
        {
            "country": "RU",
            "total_eth_value": 0.024162,
            "total_fiat_value": 9.61210186
        },
        {
            "country": "SE",
            "total_eth_value": 0.023,
            "total_fiat_value": 9.132541
        },
        {
            "country": "LV",
            "total_eth_value": 0.01,
            "total_fiat_value": 3.9209899999999998
        },
        {
            "country": "AL",
            "total_eth_value": 0.003,
            "total_fiat_value": 1.126449
        }
    ],
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-heat-map`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from timestamp (millisecond)
toTime | true | integer | to timestamp (millisecond)
timezone | true | integer | timezone to get country stats from -11 to 14

## Get heat map for token

```shell
curl -x GET \
    "http://localhost:8000/get-token-heatmap?\
    fromTime=1518307200000&\
    token=EOS&\
    freq=D&\
    toTime=1518911999999"
```

> sample response:

```json
{
    "data": [
        {
            "country": "US",
            "volume": 2883.620428022146,
            "eth_volume": 29.97000000311978,
            "usd_volume": 28584.013502715607
        },
        {
            "country": "unknown",
            "volume": 663.7763113279779,
            "eth_volume": 6.848675774186141,
            "usd_volume": 5710.033060275751
        },
        {
            "country": "JP",
            "volume": 189.38349888667832,
            "eth_volume": 1.99,
            "usd_volume": 1881.86987
        },
        {
            "country": "KR",
            "volume": 93.83012247596538,
            "eth_volume": 1,
            "usd_volume": 857.766
        },
        {
            "country": "SI",
            "volume": 73.000042,
            "eth_volume": 0.7584920000216375,
            "usd_volume": 696.7810908998771
        },
        {
            "country": "IL",
            "volume": 9.757144977962138,
            "eth_volume": 0.1,
            "usd_volume": 85.47670000000001
        },
        {
            "country": "TH",
            "volume": 9.459436814264475,
            "eth_volume": 0.1,
            "usd_volume": 84.1759
        },
        {
            "country": "DE",
            "volume": 9.311558446913438,
            "eth_volume": 0.09904,
            "usd_volume": 85.93066944
        },
        {
            "country": "VN",
            "volume": 1.8918873628528947,
            "eth_volume": 0.019789900740301923,
            "usd_volume": 16.536080320374314
        }
    ],
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/get-token-heatmap`


### URL Params:

Parameter | Required | Type |  Description
--------- | -------- | ---- | -----------
fromTime | true | integer | timestamp to get data from (millisecond)
toTime |true | integer | timestamp to get data to (millisecond)
token | true | string | name of token to get heatmap
freq | true | string | frequencty to get volume ("M", "H", "D" - Minute, Hour, Day)

## Get gold data

> sample response:

```json
{
    "data": {
        "Timestamp": 1526923808631,
        "DGX": {
            "Valid": true,
            "Timestamp": 0,
            "success": "",
            "data": [
                {
                    "symbol": "DGXETH",
                    "price": 0.06676463,
                    "time": 1526923801
                },
                {
                    "symbol": "ETHUSD",
                    "price": 694.4,
                    "time": 1526923801
                },
                {
                    "symbol": "ETHSGD",
                    "price": 931.89,
                    "time": 1526923801
                },
                {
                    "symbol": "DGXUSD",
                    "price": 46.36,
                    "time": 1526923801
                },
                {
                    "symbol": "EURUSD",
                    "price": 1.17732,
                    "time": 1526923801
                },
                {
                    "symbol": "USDSGD",
                    "price": 1.34201,
                    "time": 1526923801
                },
                {
                    "symbol": "XAUUSD",
                    "price": 1291.468,
                    "time": 1526923801
                },
                {
                    "symbol": "USDJPY",
                    "price": 111.061,
                    "time": 1526923801
                }
            ],
            "Error": ""
        },
        "OneForgeETH": {
            "Value": 1.85646,
            "Text": "1 XAU is worth 1.85646 ETH",
            "Timestamp": 1526923803,
            "Error": false,
            "Message": ""
        },
        "OneForgeUSD": {
            "Value": 1291.57,
            "Text": "1 XAU is worth 1291.57 USD",
            "Timestamp": 1526923803,
            "Error": false,
            "Message": ""
        },
        "GDAX": {
            "Valid": true,
            "Error": "",
            "trade_id": 34527604,
            "price": "695.56000000",
            "size": "0.00894700",
            "bid": "695.55",
            "ask": "695.56",
            "volume": "50497.82498957",
            "time": "2018-05-21T17:30:04.729000Z"
        },
        "Kraken": {
            "Valid": true,
            "network_error": "",
            "error": [],
            "result": {
                "XETHZUSD": {
                    "a": [
                        "696.66000",
                        "1",
                        "1.000"
                    ],
                    "b": [
                        "696.33000",
                        "4",
                        "4.000"
                    ],
                    "c": [
                        "696.33000",
                        "0.10776064"
                    ],
                    "v": [
                        "13536.83019524",
                        "16999.30348103"
                    ],
                    "p": [
                        "707.93621",
                        "710.18316"
                    ],
                    "t": [
                        5361,
                        8276
                    ],
                    "l": [
                        "693.97000",
                        "693.97000"
                    ],
                    "h": [
                        "721.38000",
                        "724.80000"
                    ],
                    "o": "715.65000"
                }
            }
        },
        "Gemini": {
            "Valid": true,
            "Error": "",
            "bid": "694.50",
            "ask": "695.55",
            "volume": {
                "ETH": "11418.5646926",
                "USD": "8064891.13775284649999999999999999999704534",
                "timestamp": 1526923800000
            },
            "last": "695.36"
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/gold-feed`