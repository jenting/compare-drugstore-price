#!/bin/bash

set -o errexit

## go build
go build

## run binary
./compare-drugstore-price
