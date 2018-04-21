# Compare Drugstore Prices

Compare prices between drugstores (Watson and Poya).

## Crawling

The crawling steps:
1) Send HTTP GET request with parameter query name to multiple drugstores' URL with goroutine
2) Parse all the products' name and price from HTTP GET response HTML content
3) Return the product name, product price, and shop name back to client

## Crawling with persistent database

The crawling steps:
1) Get the query name, product name, product price, shop name, and update time from persistent database
2) If query name exist and current time minus update time less than or equals N days, go to 8)
3) If query name exist and current time minus update time large than N days, go to 5)
4) If query name not exist, go to 5)
5) Send HTTP GET request with parameter query name to multiple drugstores' URL with goroutine
6) Parse all the products' name and price from HTTP GET response HTML content
7) Save the query name, product name, product price, shop name, and update time to persistent database
8) Return the product name, product price, and shop name back to client

## Crawling with in-memory cache (w/ timeout mechanism) and persistent databae

The crawling steps:
1) Get the query name, product name, product price, shop name, and update time from in-memory cache
2) If query name exist in-memory cache, go to 11)
3) If query name not exist in-memory cache, go to 4)
4) Get the query name, product name, product price, shop name, and update time from persistent database
5) If query name exist and current time minus update time less than or equals N days, go to 11)
6) If query name exist and current time minus update time large than N days, go to 8)
7) If query name not exist, go to 8)
8) Send HTTP GET request with parameter query name to multiple drugstores' URL with goroutine
9) Parse all the products' name and price from HTTP GET response HTML content
10) Save the query name, product name, product price, shop name, and update time to in-memory cache and persistent database
11) Return the product name, product price, and shop name back to client

## RESTful APIs

* CRUD

|    Method   |     URL     | Description |
|-------------|-------------|-------------|
| GET | <http://localhost/v1/search?name={name}&sort=false> | Query product name's price for all drugstore shop without sorting. |
| GET | <http://localhost/v1/search?name={name}&sort=true&order=increase> | Query product name's price for all drugstore shop with sorting (ascending order). |
| GET | <http://localhost/v1/search?name={name}&sort=true&order=decrease> | Query product name's price for all drugstore shop with sorting (descending order). |

* HTTP Response JSON arry with JSON format

|    Field     | Type(Length) |  Description |
|--------------|--------------|--------------|
|     shop     |  String(16)  |   Shop name  |
|     name     |  String(128) | Product name |
|     price    |  Integer     | Product price|

## TODO

* [ ] Support crawling cosmed HTML content
* [ ] Support [prometheus](https://prometheus.io) metrics API
* [ ] Analyze in-memory cache hit rate
* [ ] Analyze threshold N days

## Godep

* Add dependency

```sh
godep save ./...
```

* Restore currently vendored deps to the $GOPATH

```sh
godep restore
```