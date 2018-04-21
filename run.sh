#!/bin/bash

set -o errexit

## Start mysql docker service.
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

## go build
go build

## run binary
./compare-price-drugstore
