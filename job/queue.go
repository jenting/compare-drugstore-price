package job

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/jenting/compare-drugstore-price/data"
)

var (
	maxQueueSize = flag.Int("max-queue-size", 100, "Maximum job queue size")
	jobQueue     chan Info
)

// Info describe how jobs composed.
type Info struct {
	QueryName  string
	Timeout    time.Duration
	CrawlerRet chan data.ProductInfoList
	//NotifyChan   chan bool
}

func init() {
	glog.Infof("Start jobqueue size: %d", *maxQueueSize)
	jobQueue = make(chan Info, *maxQueueSize)
}

// CreateNewJob creates a new job and enqueues to jobqueue.
func CreateNewJob(QueryName string) Info {
	crawlerRet := make(chan data.ProductInfoList)

	job := Info{
		QueryName:  QueryName,
		Timeout:    30 * time.Second,
		CrawlerRet: crawlerRet,
	}

	// Job enqueue
	jobQueue <- job

	return job
}
