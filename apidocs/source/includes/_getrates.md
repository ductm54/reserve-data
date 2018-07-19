# Rate

## Get token rates from blockchain

```
curl -X GET "http://127.0.0.1:8000/getrates"
```

> response:

```json
  {
    "data": {
        "ADX": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 371.0142432353458,
            "CompactBuy": 0,
            "BaseSell": 0.002538305711940429,
            "CompactSell": 0,
            "Rate": 0,
            "Block": 2420849
        },
        "BAT": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 1656.6398539506304,
            "CompactBuy": 0,
            "BaseSell": 0.0005684685,
            "CompactSell": 0,
            "Rate": 0,
            "Block": 2420849
        },
        "CVC": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 1051.2127184124374,
            "CompactBuy": -1,
            "BaseSell": 0.00089586775,
            "CompactSell": 1,
            "Rate": 0,
            "Block": 2420849
        },
        "DGD": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 5.662106994812361,
            "CompactBuy": 0,
            "BaseSell": 0.16632458088099816,
            "CompactSell": 0,
            "Rate": 0,
            "Block": 2420849
        },
        "EOS": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 121.11698932232625,
            "CompactBuy": -15,
            "BaseSell": 0.007775519999999998,
            "CompactSell": 15,
            "Rate": 0,
            "Block": 2420849
        },
        "ETH": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 0,
            "CompactBuy": 30,
            "BaseSell": 0,
            "CompactSell": -29,
            "Rate": 0,
            "Block": 2420849
        },
        "FUN": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 6805.131583093689,
            "CompactBuy": 33,
            "BaseSell": 0.000138387856475128,
            "CompactSell": -32,
            "Rate": 0,
            "Block": 2420849
        },
        "GNT": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 1055.0281030473377,
            "CompactBuy": -74,
            "BaseSell": 0.0010113802,
            "CompactSell": -47,
            "Rate": 0,
            "Block": 2420849
        },
        "KNC": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 229.65128829779712,
            "CompactBuy": 89,
            "BaseSell": 0.004100772,
            "CompactSell": -82,
            "Rate": 0,
            "Block": 2420849
        },
        "LINK": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 844.2527577938458,
            "CompactBuy": 101,
            "BaseSell": 0.0011154806,
            "CompactSell": -91,
            "Rate": 0,
            "Block": 2420849
        },
        "MCO": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 63.99319226272073,
            "CompactBuy": 21,
            "BaseSell": 0.014716371218820246,
            "CompactSell": -20,
            "Rate": 0,
            "Block": 2420849
        },
        "OMG": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 44.45707162223901,
            "CompactBuy": 30,
            "BaseSell": 0.021183301968644246,
            "CompactSell": -29,
            "Rate": 0,
            "Block": 2420849
        },
        "PAY": {
            "Valid": true,
            "Error": "",
            "Timestamp": "1515412582435",
            "ReturnTime": "1515412582710",
            "BaseBuy": 295.08854913901575,
            "CompactBuy": -13,
            "BaseSell": 0.003191406699999999,
            "CompactSell": 13,
            "Rate": 0,
            "Block": 2420849
        }
    },
    "success": true,
    "timestamp": "1515412583215",
    "version": 1515412582435
}
```

### HTTP Request
`<host>:8000/getrates`


## Get all token rates from blockchain

```shell
curl -X GET "http://127.0.0.1:8000/get-all-rates?\
fromTime=1530576000000&\
toTime=1530662400000"
```

> response

```json
{
    "data": [
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280618739",
            "ReturnTime": "1517280619071",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280618739",
                    "ReturnTime": "1517280619071",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280618739",
                    "ReturnTime": "1517280619071",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280618739",
                    "ReturnTime": "1517280619071",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280618739",
                    "ReturnTime": "1517280619071",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280618739",
                    "ReturnTime": "1517280619071",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280618739",
                    "ReturnTime": "1517280619071",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280621738",
            "ReturnTime": "1517280622251",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280621738",
                    "ReturnTime": "1517280622251",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280621738",
                    "ReturnTime": "1517280622251",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280621738",
                    "ReturnTime": "1517280622251",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280621738",
                    "ReturnTime": "1517280622251",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280621738",
                    "ReturnTime": "1517280622251",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280621738",
                    "ReturnTime": "1517280622251",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280624739",
            "ReturnTime": "1517280625052",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280624739",
                    "ReturnTime": "1517280625052",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280624739",
                    "ReturnTime": "1517280625052",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280624739",
                    "ReturnTime": "1517280625052",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280624739",
                    "ReturnTime": "1517280625052",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280624739",
                    "ReturnTime": "1517280625052",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280624739",
                    "ReturnTime": "1517280625052",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280627735",
            "ReturnTime": "1517280628664",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280627735",
                    "ReturnTime": "1517280628664",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280627735",
                    "ReturnTime": "1517280628664",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280627735",
                    "ReturnTime": "1517280628664",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280627735",
                    "ReturnTime": "1517280628664",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280627735",
                    "ReturnTime": "1517280628664",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280627735",
                    "ReturnTime": "1517280628664",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280630737",
            "ReturnTime": "1517280631266",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280630737",
                    "ReturnTime": "1517280631266",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280630737",
                    "ReturnTime": "1517280631266",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280630737",
                    "ReturnTime": "1517280631266",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280630737",
                    "ReturnTime": "1517280631266",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280630737",
                    "ReturnTime": "1517280631266",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280630737",
                    "ReturnTime": "1517280631266",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280633737",
            "ReturnTime": "1517280634096",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280633737",
                    "ReturnTime": "1517280634096",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280633737",
                    "ReturnTime": "1517280634096",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280633737",
                    "ReturnTime": "1517280634096",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280633737",
                    "ReturnTime": "1517280634096",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280633737",
                    "ReturnTime": "1517280634096",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280633737",
                    "ReturnTime": "1517280634096",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280636736",
            "ReturnTime": "1517280637187",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280636736",
                    "ReturnTime": "1517280637187",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280636736",
                    "ReturnTime": "1517280637187",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280636736",
                    "ReturnTime": "1517280637187",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280636736",
                    "ReturnTime": "1517280637187",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280636736",
                    "ReturnTime": "1517280637187",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280636736",
                    "ReturnTime": "1517280637187",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        },
        {
            "Version": 0,
            "Valid": true,
            "Error": "",
            "Timestamp": "1517280639741",
            "ReturnTime": "1517280640213",
            "Data": {
                "EOS": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280639741",
                    "ReturnTime": "1517280640213",
                    "BaseBuy": 87.21360760013062,
                    "CompactBuy": 0,
                    "BaseSell": 0.0128686459657361,
                    "CompactSell": 0,
                    "Rate": 0,
                    "Block": 5635245
                },
                "ETH": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280639741",
                    "ReturnTime": "1517280640213",
                    "BaseBuy": 0,
                    "CompactBuy": 32,
                    "BaseSell": 0,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "KNC": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280639741",
                    "ReturnTime": "1517280640213",
                    "BaseBuy": 307.05930436561505,
                    "CompactBuy": -34,
                    "BaseSell": 0.003084981280661941,
                    "CompactSell": 81,
                    "Rate": 0,
                    "Block": 5635245
                },
                "OMG": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280639741",
                    "ReturnTime": "1517280640213",
                    "BaseBuy": 65.0580993582104,
                    "CompactBuy": 32,
                    "BaseSell": 0.014925950060437398,
                    "CompactSell": -14,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SALT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280639741",
                    "ReturnTime": "1517280640213",
                    "BaseBuy": 152.3016783627643,
                    "CompactBuy": 9,
                    "BaseSell": 0.006196212698403499,
                    "CompactSell": 23,
                    "Rate": 0,
                    "Block": 5635245
                },
                "SNT": {
                    "Valid": true,
                    "Error": "",
                    "Timestamp": "1517280639741",
                    "ReturnTime": "1517280640213",
                    "BaseBuy": 4053.2170631085987,
                    "CompactBuy": 43,
                    "BaseSell": 0.000233599514875301,
                    "CompactSell": -3,
                    "Rate": 0,
                    "Block": 5635245
                }
            }
        }
    ],
    "success": true
}
```

### HTTP Request

`<host>:8000/get-all-rates`

### Query parameter:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | false | integer | get all rates from this timepoint (millisecond)
toTime | false | integer | get all rates to this timepoint (millisecond)

<aside class="notice">Restriction: toTime - fromTime <= 1 day (in millisecond)</aside>