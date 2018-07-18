# Target Quantity

## Set target quantity v2

**signing required**

> sample response:

```json
on success:
{"success":true}


on failure:
{"success":false,
 "reason":<error>}
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

> sample response:

```json

on success:
{"success":true}

on failure:
{
    "success":false,
    "reason":<error>
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

> sample response:

```json

on success:
{"success":true}


on failure:
{"success":false,
 "reason":<error>}
```

### HTTP Request

**POST**
`<host>:8000/v2/canceltargetqty`



## Get pending target quantity v2

**signing required**

Return the current pending target quantity 

### HTTP Request

**GET**
`<host>:8000/v2/pendingtargetqty`

```shell
curl -x GET \
  http://localhost:8000/v2/pendingtargetqty?nonce=111111
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

## Get target quantity v2

**signing required**

Return the current confirmed target quantity 

```shell
curl -x GET \
  http://localhost:8000/v2/targetqty?nonce=111111
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

### URL params:

Parameter | Required | Description
--------- | -------- | -----------
nonce | true | (uint64) : the nonce to conform to signing requirement