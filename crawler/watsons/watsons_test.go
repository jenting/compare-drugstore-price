package watsons

import (
	"testing"
	"time"
)

const defaultTimeout = 30 // seconds

func TestWatsons(t *testing.T) {
	Crawler("洗髮精", time.Duration(defaultTimeout)*time.Second)
}
