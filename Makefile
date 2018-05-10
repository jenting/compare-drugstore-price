#!/usr/bin/make -f

all:
	go build
	
clean:
	@[ -f compare-drugstore-price ] && rm -r compare-drugstore-price || true

.PHONY: all clean
