# Metrics

Save metrics information

## Update Price Analytic Data

**signing required**

Set a record marking the condition because of which the set price is called. 

> sample response:

```json
on success:
{"success":true}

on failure:
{"success":false,
 "reason":<error>}
```

### HTTP Request

**GET**
`<host>:8000/update-price-analytic-data`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
timestamp | true | integer |  the timestamp of the action (real time ) in millisecond
value | true |  bool | the json enconded object to save. 

<aside class="notice">Note: the data sent over must be encoded in Json in order to make it valid for output operation
  In Python, the data would be encoded as:
   data = {"timestamp": timestamp, "value": json.dumps(analytic_data)} </aside>

## Get Price Analytic Data

**signing required**

```shell
curl -x GET \
  http://localhost:8000/get-price-analytic-data?fromTime=1522753160000&toTime=1522755792000
```
 
> sample response:

```json
{
  "data": [
    {
      "Timestamp": 1522755271000,
      "Data": {
        "block_expiration": false,
        "trigger_price_update": true,
        "triggering_tokens_list": [
          {
            "ask_price": 0.002,
            "bid_price": 0.003,
            "mid afp_old_price": 0.34555,
            "mid_afp_price": 0.6555,
            "min_spread": 0.233,
            "token": "OMG"
          },
          {
            "ask_price": 0.004,
            "bid_price": 0.005,
            "mid afp_old_price": 0.21555,
            "mid_afp_price": 0.4355,
            "min_spread": 0.133,
            "token": "KNC"
          }
        ]
      }
    }
  ],
  "success": true
}
```

List of price analytic data, sorted by timestamp 

**GET**
`<host>:8000/get-get-price-analytic-data`

### URL params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
fromTime | true | integer | from timestamp (millisecond)
toTime | true | integer | to timestamp (millisecond)


## Update exchange notifications 

> sample response:

```json
  {
    "success": true
  }
```

### HTTP Request

**POST**
`<host>:8000/exchange-notification`

### Form params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
exchange | true | string | exchange name
action | true | string | action name
token | true | string | token pair
fromTime | true | integer | from timestamp
toTime | true | integer | to timestamp
isWarning | true | bool | is exchange warning or not
msg | true | string | message for the notification


## Get exchange notifications

> sample response:

```json
{
    "data": {
        "binance": {
            "trade": {
                "OMG": {
                    "fromTime": 123,
                    "toTime": 125,
                    "isWarning": true,
                    "msg": "3 times"
                }
            }
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/exchange-notifications`

## Get rebalance status

> sample response:

```json
  {
    "success": true,
    "data": true
  }
```

Get rebalance status, if reponse is *true* then rebalance is enable, the analytic can perform rebalance, else reponse is *false*, the analytic hold rebalance ability.

### HTTP Request

**GET**
`<host>:8000/rebalancestatus`


## Hold rebalance

```shell
curl -X POST \
  http://localhost:8000/holdrebalance \
  -H 'content-type: multipart/form-data' \
```


> sample response

```json
  {
    "success": true
  }
```

### HTTP Request

**POST**
`<host>:8000/holdrebalance`


## Enable rebalance

```shell
curl -X POST \
  http://localhost:8000/enablerebalance \
  -H 'content-type: multipart/form-data' \
```

> sample response:

```json
  {
    "success": true
  }
```

### HTTP Request

**POST**
`<host>:8000/enablerebalance`


## Get setrate status

> sample response:

```json
  {
    "success": true,
    "data": true
  }
```

Get setrate status, if reponse is *true* then setrate is enable, the analytic can perform setrate, else reponse is *false*, the analytic hold setrate ability.

### HTTP Request

**GET**
`<host>:8000/setratestatus`


## Hold setrate

### HTTP Request

**POST**
`<host>:8000/holdsetrate`

```shell
curl -X POST \
  http://localhost:8000/holdsetrate \
  -H 'content-type: multipart/form-data' \
```

> sample response:

```json
  {
    "success": true
  }
```

## Enable setrate

### HTTP Request

**POST**
`<host>:8000/enablesetrate`

```shell
curl -X POST \
  http://localhost:8000/enablesetrate \
  -H 'content-type: multipart/form-data' \
```

> sample response:

```json
  {
    "success": true
  }
```