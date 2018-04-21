package job

import (
	"flag"

	"github.com/golang/glog"
)

var (
	maxQueueSize = flag.Int("max-queue-size", 100, "Maximum job queue size")
)

var (
	JobQueue   chan JobInfo
	CrawlerRet chan CrawlerJSON
)

// JobInfo describe how jobs composed.
type JobInfo struct {
	QueryName  string
	NotifyChan chan bool
}

type CrawlerJSON struct {
	showName     string
	productName  string
	productPrice string
}

func init() {
	glog.Infof("Start jobqueue size: %d", *maxQueueSize)
	JobQueue = make(chan JobInfo, *maxQueueSize)
	CrawlerRet = make(chan CrawlerJSON, *maxQueueSize)
}

// Enqueue puts job to jobqueue
func Enqueue(job JobInfo) {
	JobQueue <- job
}
