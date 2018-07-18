---
title: API Reference

language_tabs: # must be one of https://git.io/vQNgJ
  - shell

toc_footers:
  - <a href='#'>Sign Up for a Developer Key</a>
  - <a href='https://github.com/lord/slate'>Documentation Powered by Slate</a>

includes:
  - exchangeinfo
  - authdata
  - getrates
  - rebalance_activities
  - stats
  - metrics/metrics
  - metrics/pwi_equation
  - metrics/target_quantity
  - metrics/rebalance_quadratic
  - metrics/stable_token
  - errors

search: true
---

# Introduction


# Authentication

All APIs that are marked with (signing required) must follow authentication mechanism below:

1. Must be urlencoded (**x-www-form-urlencoded**)
1. Must have `signed` header with value equals to `hmac512(secret, message)`
1. Must contain `nonce` param, its value is the unix time in millisecond, it must not be before or after server time by 10s
1. `message` is constructed in following way: all query params (nonce is included) and body key-values are merged into one urlencoded string with keys are sorted.
1. `secret` is configured secret string.

Example:  

- param query: `amount=0xde0b6b3a7640000&nonce=1514554594528&token=KNC`  
- secret: `vtHpz1l0kxLyGc4R1qJBkFlQre5352xGJU9h8UQTwUTz5p6VrxcEslF4KnDI21s1`  
- signed string: `2969826a713d13b399dd0d016dad3e95949aa81ed8703ec0258abebb5f0288b96272eef68275f12a32f7e396de3b5fd63ed12b530385e08e1b676c695aacb93b`

# General 

## Get time server 


```shell
curl -X GET "http://localhost:8000/timeserver"
```

> The above command returns JSON structured like this:

```json
{
  "data": "1517479497447",
  "success": true
}
```

This endpoint return current time server

### HTTP Request

`GET http://example.com/timeserver`


## Get all addresses are being used by core 

```shell
curl -X GET "http://localhost:8000/core/addresses"
```


> The above command returns JSON structured like this:

```json
{
  "data": {
    "tokens": {
      "EOS": "0x15fb2a9d7dadbb88f260f78dcbb574b3b76a8e06",
      "ETH": "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
      "KNC": "0x8dc114d77e857558aefbe8e1a50b460ff9578f1a",
      "OMG": "0x7606bd550f467546212649a9c25623dfca88dcd7",
      "SALT": "0xcc112cd38362bf3c07d226768fd5869e65296083",
      "SNT": "0x676f650000f420485b99ef0377a2e1c96eb3e821"
    },
    "exchanges": {
      "binance": {
        "EOS": "0x1ae659f93ba2fc0a1f379545cf9335adb75fa547",
        "ETH": "0x1ae659f93ba2fc0a1f379545cf9335adb75fa547",
        "KNC": "0x1ae659f93ba2fc0a1f379545cf9335adb75fa547",
        "OMG": "0x1ae659f93ba2fc0a1f379545cf9335adb75fa547",
        "SALT": "0x1ae659f93ba2fc0a1f379545cf9335adb75fa547",
        "SNT": "0x1ae659f93ba2fc0a1f379545cf9335adb75fa547"
      },
      "bittrex": {
        "EOS": "0xef6ee90c5bb23da2eb71b3daa8e57b204e5ac647",
        "ETH": "0xe0355aa3cc0a4e0e4b2a70acd90c2fa961f61b23",
        "KNC": "0x132478f1ec4b8e1256b11fdf3e00d97e4df5988f",
        "OMG": "0x9db6e8d2d133448dbcf755f19d540253da4ba043",
        "SALT": "0x385d619b530f00ab7d082683f7cdc37995ac76f2",
        "SNT": "0x3ef96f9de64c44b1ad392b10e2277a73ec14ff5f"
      }
    },
    "wrapper": "0xa54f27b5a72fc1ddc5c4bc6ed50391f457e4a46a",
    "pricing": "0x77925520469d0fcbb0311814c053bf9bafcd867b",
    "reserve": "0x2d1ceabd5a1cd16581ad199031601615a434a2cd",
    "feeburner": "0xa33a2f0745ee8e31b753ec33d22d363a62a123a4",
    "network": "0x643211b405c9a14139142e1104250bbcd94bd0ef"
  },
  "success": true
}
```


### HTTP Request

`GET http://example.com/core/addresses`


## Get prices all base-quote pair

```shell
curl -X GET "http://localhost:8000/prices"

```

> The above command returns JSON structured like this:

```json
{
  "data": {
    "ADX-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 1534.265,
            "Rate": 0.00250437
          },
          {
            "Quantity": 1147.78359847,
            "Rate": 0.00250435
          },
          {
            "Quantity": 426.37538021,
            "Rate": 0.00250429
          }
        ],
        "Asks": [
          {
            "Quantity": 4850.84,
            "Rate": 0.00277997
          },
          {
            "Quantity": 144.04135361,
            "Rate": 0.00277998
          },
          {
            "Quantity": 14.50780994,
            "Rate": 0.00278059
          }
        ],
        "ReturnTime": "1514114579641"
      }
    },
    "BAT-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 15173.85912685,
            "Rate": 0.00047374
          },
          {
            "Quantity": 130552,
            "Rate": 0.00047363
          },
          {
            "Quantity": 2149.78448276,
            "Rate": 0.0004734
          }
        ],
        "Asks": [
          {
            "Quantity": 660.96951182,
            "Rate": 0.00048652
          },
          {
            "Quantity": 476.36673132,
            "Rate": 0.00048663
          },
          {
            "Quantity": 53661.5,
            "Rate": 0.00048668
          }
        ],
        "ReturnTime": "1514114579480"
      }
    },
    "CVC-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 128.67287333,
            "Rate": 0.00099655
          },
          {
            "Quantity": 500,
            "Rate": 0.00098795
          },
          {
            "Quantity": 45.30924007,
            "Rate": 0.00098539
          }
        ],
        "Asks": [
          {
            "Quantity": 153.22180315,
            "Rate": 0.001
          },
          {
            "Quantity": 7010.72355807,
            "Rate": 0.00101567
          },
          {
            "Quantity": 2679.69026772,
            "Rate": 0.00101568
          }
        ],
        "ReturnTime": "1514114579642"
      }
    },
    "DGD-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 0.27508146,
            "Rate": 0.21463293
          },
          {
            "Quantity": 8.61103292,
            "Rate": 0.21463292
          },
          {
            "Quantity": 1,
            "Rate": 0.21462222
          }
        ],
        "Asks": [
          {
            "Quantity": 1.43683366,
            "Rate": 0.22554555
          },
          {
            "Quantity": 0.10879304,
            "Rate": 0.22554557
          },
          {
            "Quantity": 0.06252449,
            "Rate": 0.22554606
          }
        ],
        "ReturnTime": "1514114579641"
      }
    },
    "FUN-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 3550.94852427,
            "Rate": 0.00008065
          },
          {
            "Quantity": 24900,
            "Rate": 0.00008064
          },
          {
            "Quantity": 489855.39183168,
            "Rate": 0.00008063
          }
        ],
        "Asks": [
          {
            "Quantity": 3635.15493421,
            "Rate": 0.00008282
          },
          {
            "Quantity": 3905.9918732,
            "Rate": 0.00008293
          },
          {
            "Quantity": 1952.93876331,
            "Rate": 0.00008366
          }
        ],
        "ReturnTime": "1514114579574"
      }
    },
    "GNT-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 1E-8,
            "Rate": 0.0008552
          },
          {
            "Quantity": 2000,
            "Rate": 0.00084661
          },
          {
            "Quantity": 35.4,
            "Rate": 0.0008466
          }
        ],
        "Asks": [
          {
            "Quantity": 7209.279,
            "Rate": 0.00086879
          },
          {
            "Quantity": 399.58082001,
            "Rate": 0.0008688
          },
          {
            "Quantity": 7185.948,
            "Rate": 0.00086893
          }
        ],
        "ReturnTime": "1514114579457"
      }
    },
    "MCO-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 142.93534777,
            "Rate": 0.02437378
          },
          {
            "Quantity": 1.21116959,
            "Rate": 0.02437377
          },
          {
            "Quantity": 1.63701658,
            "Rate": 0.02437376
          }
        ],
        "Asks": [
          {
            "Quantity": 15.39680469,
            "Rate": 0.02503471
          },
          {
            "Quantity": 18.71484714,
            "Rate": 0.02503534
          },
          {
            "Quantity": 93.57423573,
            "Rate": 0.02503537
          }
        ],
        "ReturnTime": "1514114579481"
      }
    },
    "OMG-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 5.49,
            "Rate": 0.019857
          },
          {
            "Quantity": 13.62550123,
            "Rate": 0.0197758
          },
          {
            "Quantity": 10,
            "Rate": 0.01976677
          },
          {
            "Quantity": 6.92629385,
            "Rate": 0.01970274
          }
        ],
        "Asks": [
          {
            "Quantity": 6.73770653,
            "Rate": 0.02025768
          },
          {
            "Quantity": 7.49193537,
            "Rate": 0.02025774
          },
          {
            "Quantity": 1.48831433,
            "Rate": 0.02025781
          }
        ],
        "ReturnTime": "1514114579575"
      }
    },
    "PAY-ETH": {
      "bittrex": {
        "Valid": true,
        "Error": "",
        "Timestamp": "1514114579228",
        "Bids": [
          {
            "Quantity": 17.76916985,
            "Rate": 0.00576079
          },
          {
            "Quantity": 25,
            "Rate": 0.0057565
          },
          {
            "Quantity": 5.24,
            "Rate": 0.005728
          }
        ],
        "Asks": [
          {
            "Quantity": 136.4072,
            "Rate": 0.00581225
          },
          {
            "Quantity": 776.223,
            "Rate": 0.00583147
          },
          {
            "Quantity": 15.90915084,
            "Rate": 0.00583148
          }
        ],
        "ReturnTime": "1514114579574"
      }
    }
  },
  "success": true,
  "timestamp": "1514114582015",
  "version": 64
}
```

### HTTP request

`GET http://example.com/prices`