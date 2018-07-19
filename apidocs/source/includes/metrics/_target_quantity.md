# Target Quantity

## Set target quantity v2

**signing required**

```shell
curl -X POST "http://localhost:8000/v2/settargetqty" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "value={
    \"OMG\" : {
      \"TotalTarget\": 1500,
      \"ReserveTarget\": 1005,
      \"RebalanceThreshold\": 0.33,
      \"TransferThreshold\": 0.2
    }
  }"
```

> sample response:

```json
{
  "success":true
}
```

### HTTP Request

**POST**
`<host>:8000/v2/settargetqty`

### Form Params:

Parameter | Required | Description
--------- | -------- | -----------
value (string) | true | the json enconded string, represent a map (string : interface)


## Confirm target quantity v2

**signing required**

```shell
curl -X POST "http://localhost:8000/v2/confirmtargetqty" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "value={
    \"OMG\" : {
      \"TotalTarget\": 1500,
      \"ReserveTarget\": 1005,
      \"RebalanceThreshold\": 0.33,
      \"TransferThreshold\": 0.2
    }
  }"
```

> sample response:

```json
  {
    "success":true
  }
```

### HTTP Request

**POST**
`<host>:8000/v2/confirmtargetqty`

### URL Params:

Parameter | Required | Description
--------- | -------- | -----------
value (string) | true | the json enconded string, represent a map (string : interface), must be equal to current pending.



## Cancel set target quantity v2

**signing required**

```shell
curl -X POST "http://localhost:8000/v2/canceltargetqty
```

> sample response:

```json
  {
    "success":true
  }
```

### HTTP Request

**POST**
`<host>:8000/v2/canceltargetqty`



## Get pending target quantity v2

**signing required**

```shell
curl -X GET "http://localhost:8000/v2/pendingtargetqty"
```
 
> sample response:

```json
{
    "data": {
        "OMG": {
            "set_target": {
                "total_target": 0,
                "reserve_target": 0,
                "rebalance_threshold": 0,
                "transfer_threshold": 0
            }
        }
    },
    "success": true
}
```

Return the current pending target quantity 

### HTTP Request

**GET**
`<host>:8000/v2/pendingtargetqty`


## Get target quantity v2

**signing required**

Return the current confirmed target quantity 

```shell
curl -X GET "http://localhost:8000/v2/targetqty"
```
 
> sample response:

```json
{
  "data": {
    "OMG" : {
      "TotalTarget": 1500,
      "ReserveTarget": 1005,
      "RebalanceThreshold": 0.33,
      "TransferThreshold": 0.2
    }
  },
  "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/v2/targetqty`