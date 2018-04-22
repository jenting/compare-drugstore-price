package job

import (
	"flag"
	"sync"

	"github.com/golang/glog"
	"github.com/hsiaoairplane/compare-drugstore-price/crawler/poya"
	"github.com/hsiaoairplane/compare-drugstore-price/crawler/watsons"
	"github.com/hsiaoairplane/compare-drugstore-price/data"
)

var (
	maxWorkerNum = flag.Int("max-worker-num", 16, "Maximum worker number")
)

func jobWorker(id int) {
	for {
		select {
		case job := <-jobQueue:
			glog.Infof("Worker ID: %v QueryName: %v", id, job.QueryName)

			// Sync two goroutine with sync.WaitGroup
			wg := sync.WaitGroup{}
			var ret1 []data.ProductInfo
			var ret2 []data.ProductInfo
			var ret []data.ProductInfo

			// Add in main func()
			wg.Add(1)
			go func() {
				ret1 = watsons.Crawler(job.QueryName, job.Timeout)
				// Done in goroutine
				wg.Done()
			}()

			// Add in main func()
			wg.Add(1)
			go func() {
				ret2 = poya.Crawler(job.QueryName, job.Timeout)
				// Done in goroutine
				wg.Done()
			}()

			// Waits two goroutine done
			wg.Wait()

			if ret1 != nil {
				ret = append(ret, ret1...) // append two slices
			}
			if ret2 != nil {
				ret = append(ret, ret2...) // append two slices
			}

			// Notify caller by notify channel
			job.CrawlerRet <- ret

			// Close notify channel
			close(job.CrawlerRet)
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
