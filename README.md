# Compare Drugstore Prices

Compare prices between drugstores (Watson and Poya).

## Setup

First, download the project:

    go get github.com/hsiaoairplane/compare-drugstore-price

## Crawling

The crawling steps:
1) Send HTTP GET request with parameter query name to multiple drugstores' URL with goroutine
2) Parse all the products' name and price from HTTP GET response HTML content
3) Return the product name, product price, and shop name back to client

## Crawling with in-memory cache (w/ timeout mechanism)

The crawling steps:
1) Get the query name, product name, product price, shop name, and update time from in-memory cache
2) If query name exist in-memory cache, go to 7)
3) If query name not exist in-memory cache, go to 4)
4) Send HTTP GET request with parameter query name to multiple drugstores' URL with goroutine
5) Parse all the products' name and price from HTTP GET response HTML content
6) Save the query name, product name, product price, shop name to in-memory cache
7) Return the product name, product price, and shop name back to client

## RESTful APIs

* CRUD

|    Method   |     URL     | Description |
|-------------|-------------|-------------|
| GET | <http://localhost/v1/search?name={name}> | Query product name's price for all drugstore shop without sorting. |
| GET | <http://localhost/v1/search?name={name}&sort=increase> | Query product name's price for all drugstore shop with sorting (ascending order). |
| GET | <http://localhost/v1/search?name={name}&sort=decrease> | Query product name's price for all drugstore shop with sorting (descending order). |

* HTTP Response JSON arry with JSON format

|    Field     | Type(Length) |  Description |
|--------------|--------------|--------------|
|     shop     |  String(16)  |   Shop name  |
|     name     |  String(128) | Product name |
|     price    |  Integer     | Product price|

## TODO

* [ ] Support crawling cosmed HTML content
* [ ] Support [prometheus](https://prometheus.io) metrics API
* [ ] Analyze in-memory cache hit rate and also analyze the timeout threshold for in-memory cache

## Godep

* Add dependency

```sh
godep save ./...
```

* Restore currently vendored deps to the $GOPATH

```sh
godep restore
```
