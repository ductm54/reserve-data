# PWI Equation

## Get pwis equation v2

**signing required**

```shell
    curl -X "GET" "http://localhost:8000/v2/pwis-equation" \
     -H 'Content-Type: application/x-www-form-urlencoded' \
```

> sample response:

```json
{
  "data": {
    "EOS": {
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
    "ETH": {
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
    }
  },
  "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/v2/pwis-equation`



## Get pending pwis equation v2

**signing required**

```shell
curl -X "GET" "http://localhost:8000/v2/pending-pwis-equation" \
     -H 'Content-Type: application/x-www-form-urlencoded' \
```

> sample response:

```json
{
  "data": {
    "EOS": {
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
    "ETH": {
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
    }
  },
  "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/v2/pending-pwis-equation`



## Set pwis equation v2

**signing required**

```shell
curl -X "POST" "http://localhost:8000/v2/set-pwis-equation" \
     -H 'Content-Type: application/x-www-form-urlencoded' \
     --data-urlencode "data={
  \"EOS\": {
    \"bid\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    },
    \"ask\": {
      \"a\": 800,
      \"b\": 600,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    }
  },
  \"ETH\": {
    \"bid\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    },
    \"ask\": {
      \"a\": 800,
      \"b\": 600,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    }
  }
}"
```

> sample response: 

```json
  {
    "success": true,
  }
```

### HTTP Request

**POST**
`<host>:8000/v2/set-pwis-equation`

Post form: json encoding data of pwis equation



## Confirm pending pwis equation v2

**signing required**

```shell
curl -X "POST" "http://localhost:8000/v2/confirm-pwis-equation" \
     -H 'Content-Type: application/x-www-form-urlencoded' \
     --data-urlencode "data={
  \"EOS\": {
    \"bid\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    },
    \"ask\": {
      \"a\": 800,
      \"b\": 600,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    }
  },
  \"ETH\": {
    \"bid\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    },
    \"ask\": {
      \"a\": 800,
      \"b\": 600,
      \"c\": 0,
      \"min_min_spread\": 0,
      \"price_multiply_factor\": 0
    }
  }
}"
```

> sample response:

```json
  {
    "success": true,
  }
```

### HTTP Request

**POST**
`<host>:8000/v2/confirm-pwis-equation`

Post form: json encoding data of pwis equation



## Reject pending pwis equation v2

**signing required**

```shell
    curl -X "POST" "http://localhost:8000/v2/reject-pwis-equation" \
     -H 'Content-Type: application/x-www-form-urlencoded'
```

> sample response

```json
  {
    "success": true,
  }
```

### HTTP Request

**POST**
`<host>:8000/v2/reject-pwis-equation`