package job

import (
	"flag"
	"sync"

	"github.com/golang/glog"
	"github.com/hsiaoairplane/compare-drugstore-price/crawler/poya"
	"github.com/hsiaoairplane/compare-drugstore-price/crawler/watsons"
)

var (
	maxWorkerNum = flag.Int("max-worker-num", 16, "Maximum worker number")
)

/*
type WorkerInfo struct {
	workerID int
	channel  chan interface{}
}
*/

func jobWorker(id int) {
	for {
		select {
		case job := <-JobQueue:
			glog.Infof("Worker ID: %v QueryName: %v", id, job.QueryName)

			// Sync two goroutine with sync.WaitGroup
			wg := sync.WaitGroup{}

			// Add in main func()
			wg.Add(1)
			go func() {
				watsons.Crawler(job.QueryName)
				// Done in goroutine
				wg.Done()
			}()

			// Add in main func()
			wg.Add(1)
			go func() {
				poya.Crawler(job.QueryName)
				// Done in goroutine
				wg.Done()
			}()

			// Waits two goroutine done
			wg.Wait()

			job.NotifyChan <- true
		}
	}
}

// StartWorker initialize number of workers.
func StartWorker() {
	glog.Infof("Start %d job workers", *maxWorkerNum)

	for id := 1; id <= *maxWorkerNum; id++ {
		go jobWorker(id)
	}
}
