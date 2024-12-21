package cache

import (
	"flag"
	"time"

	cache "github.com/patrickmn/go-cache"

	"github.com/jenting/compare-drugstore-price/data"
)

var (
	cacheTimeoutSec = flag.Uint("cache-timeout-second", 300, "Cache timeout in seconds")

	// Create a cache with a default expiration time of 300 seconds, and which
	// purges expired items every 300 seconds
	c = cache.New(time.Duration(*cacheTimeoutSec)*time.Second, time.Duration(*cacheTimeoutSec)*time.Second)
)

// GetCache get value from cache with specify key
func GetCache(key string) (data.ProductInfoList, bool) {
	value, found := c.Get(key)
	var p data.ProductInfoList

	if found {
		p = value.(data.ProductInfoList)
	} else {
		p = nil
	}

	return p, found
}

// SetCache set value to cache with specify key
func SetCache(key string, value data.ProductInfoList) {
	c.Set(key, value, cache.DefaultExpiration)
}
