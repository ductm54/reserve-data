# Rebalance Quadratic

## Get rebalance quadratic

**signing required**

```shell
curl -X GET "http://localhost:8000/rebalance-quadratic"
```

> sample response

```json
{
    "data": {
        "EOS": {
            "rebalance_quadratic": {
                "a": 750,
                "b": 500,
                "c": 0
            }
        },
        "ETH": {
            "rebalance_quadratic": {
                "a": 750,
                "b": 500,
                "c": 0
            }
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/rebalance-quadratic`


## Set rebalance quadratic equation

**signing required**

```shell
curl -X "POST" "http://localhost:8000/set-rebalance-quadratic" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "value={
  \"EOS\":{
    \"rebalance_quadratic\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0
    }
  },
  \"ETH\": {
    \"rebalance_quadratic\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0
    }
  }
}"
```

> sample response

```json
  {
    "success": true,
  }
```

### HTTP Request

**POST**
`<host>:8000/set-rebalance-quadratic`

Post form: json encoding data of rebalance quadratic equation

## Get pending rebalance quadratic

**signing required**

```shell
curl -X GET "http://localhost:8000/pending-rebalance-quadratic"
```

> sample response

```json
{
    "data": {
        "EOS": {
            "rebalance_quadratic": {
                "a": 750,
                "b": 500,
                "c": 0
            }
        },
        "ETH": {
            "rebalance_quadratic": {
                "a": 750,
                "b": 500,
                "c": 0
            }
        }
    },
    "success": true
}
```

### HTTP Request

**GET**
`<host>:8000/pending-rebalance-quadratic`


## Confirm rebalance quadratic equation

**signing required**

```shell
curl -X "POST" "http://localhost:8000/confirm-rebalance-quadratic" \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode "value={
  \"EOS\":{
    \"rebalance_quadratic\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0
    }
  },
  \"ETH\": {
    \"rebalance_quadratic\": {
      \"a\": 750,
      \"b\": 500,
      \"c\": 0
    }
  }
}"
```

> sample response

```json
  {
    "success": true,
  }
```

### HTTP Request

**POST**
`<host>:8000/confirm-rebalance-quadratic`

Post form: json encoding data of pwis equation



## Reject rebalance quadrtic equation

**signing required**

```shell
curl -X POST "http://localhost:8000/reject-rebalance-quadratic" \
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
`<host>:8000/reject-rebalance-quadratic`