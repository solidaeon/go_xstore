# go_xstore

An artbitrary store API.

# Build

```shell
> docker build -t xstore:v0.0.1 .
> docker run -p 8080:8080 xstore:v0.0.1
```

# Sample API
### Get All Items
```shell
> curl http://localhost:8080/items
[]
```

### Put Item
```shell
> curl --location --request POST 'http://localhost:8080/item' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "name": "soap",
      "price": 5
    }'
```