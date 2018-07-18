# Stable Token Params

## Set stable token params

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
`<host>:8000/set-stable-token-params`

### Form Params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
value (string) | true | string | the json enconded string, represent a map (string : interface)

## Confirm stable token params

> sample response:

```json
on success:
{"success":true}

on failure:
{"success":false,
 "reason":<error>}
```

**signing required**

### HTTP Request

**POST**
`<host>:8000/confirm-stable-token-params`

### Form Params:

Parameter | Required | Type | Description
--------- | -------- | ---- | -----------
value | true | string | the json enconded string, represent a map (string : interface), must be equal to current pending.

## Reject stable token params

**signing required**

### HTTP Request

**POST**
`<host>:8000/reject-stable-token-params`

## Get pending stable token params

```shell
curl -x GET \
  http://localhost:8000/pending-token-params?nonce=111111
```
 
> sample response:

```json
{
  "data": {
    "DGX": {
      "AskSpread": 50,
      "BidSpread": 50,
      "PriceUpdateThreshold": 0.1
    }
  },
  "success": true
}
```

**signing required**

Return the current pending stable token params

### HTTP Request

**GET**
`<host>:8000/pending-stable-token-params`

## Get stable token params

```shell
    curl -x GET \
    http://localhost:8000/stable-token-params?nonce=111111
```
 
> sample response:

```json
{
  "data": {
    "DGX": {
      "AskSpread": 50,
      "BidSpread": 50,
      "PriceUpdateThreshold": 0.1
    }
  },
  "success": true

```

**signing required**

Return the current confirmed stable token params

### HTTP Request

**GET**
`<host>:8000/stable-token-params`