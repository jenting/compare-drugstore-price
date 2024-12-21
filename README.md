# Compare Drugstore Prices

Compare prices between drugstores (`Watson` and `Poya`).

[![Build Status](https://travis-ci.com/jenting/compare-drugstore-price.svg?branch=master)](https://travis-ci.com/jenting/compare-drugstore-price)

## go version

Please use go version >= 1.11

## Setup

First, download the project:
```sh
go get github.com/jenting/compare-drugstore-price
```


Then run the project:
```sh
./run.sh
```

## Crawling with in-memory cache (with timeout mechanism)

The crawling steps:
1) Get the query name, product name, product price, shop name, and update time from in-memory cache
2) If query name exist in-memory cache, go to 7)
3) If query name not exist in-memory cache, go to 4)
4) Send HTTP GET request with parameter query name to multiple drugstores' URL with goroutine
5) Parse all the products' name and price from HTTP GET response HTML content
6) Save the query name, product name, product price, shop name to in-memory cache
7) Return the product name, product price, and shop name back to client

## Architecture                                                                                                                                                                                                                                                
                                                                                                                     +------------+             
             +-------------------------------------------------------+ inserts job to job queue                      |  Client 1  |             
             |    |   |   |                                     |    |                                               +------------+             
             |    |   |   |       Job Queues                    |    |<----------------------+                                                  
             |    |   |   |                                     |    |                       |                       +------------+             
             +-------------------------------------------------------+                       |                       |  Client 2  |             
                                         |                                        +--------------------+             +------------+             
                                         |                                        |                    |  HTTP GET                                     
                                         |  Worker get job from job queue         |       APIServer    | <--------->       .                    
                                         |                                        |                    |                   .                    
                                         |                                        +--------------------+                   .                    
             +---------------------------v---------------------------+                       ^                                                  
             |                                                       |                       |                       +------------+             
             | +----+   +----+                               +----+  |                       |                       |  Client n  |             
             | | W1 |   | W2 |  ...                      ... | Wn |  |-----------------------+                       +------------+             
             | +----+   +----+                               +----+  |    go channel notify                                                     
             |                     Workers Pool                      |                                                                          
             |                                                       |                                                                          
             +-------------------------------------------------------+                                                                          
               ^                         |                 ^                                                                                        
               |                         |                 |                                                                                        
               |                         v                 |   Return Found Result                                                                  
               |              +---------------------+      |                                                                                        
               |              |                     |      |                                                                                        
               |              |   in-memory cache   |------+                                                                                        
               +------------->|                     |                                                                                           
               |              +---------------------+                                                                                           
               |  Write to cache         |                                                                                                      
               |                         |  Not Found                                                                                           
               |                         |                                                                                                      
               |             +-----------v----------+                                                                                           
               |             |                      |                                                                                           
               +-------------|        Crawler       |                                                                                           
       Return Crawler Result |                      |                                                                                           
                             +----------------------+                                                                                           
                                        ^ ^                                                                                                     
                                       /   \  Sync. Wait                                                                                        
                                      /     \                                                                                                   
                                     v       v                                                                                                  
                        +---------------+  +------------+                                                                                       
                        |               |  |            |                                                                                       
                        |    Watsons    |  |     Poya   |                                                                                       
                        |               |  |            |                                                                                       
                        +---------------+  +------------+                                                                                       
                                                                                                                                                
                                                                                                                                                

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
