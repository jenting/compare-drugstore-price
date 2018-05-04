package job

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"

	"github.com/golang/glog"
	"github.com/hsiaoairplane/compare-drugstore-price/cache"
	"github.com/hsiaoairplane/compare-drugstore-price/crawler/poya"
	"github.com/hsiaoairplane/compare-drugstore-price/crawler/watsons"
	"github.com/hsiaoairplane/compare-drugstore-price/data"
)

var (
	maxWorkerNum = flag.Int("max-worker-num", 16, "Maximum worker number")
	exitSync     = sync.WaitGroup{}
)

func jobWorker(ctx context.Context, id int) {
	for {
		select {
		case job := <-jobQueue:
			glog.Infof("Worker ID: %v QueryName: %v", id, job.QueryName)

			var ret data.ProductInfoList
			ret, found := cache.GetCache(job.QueryName)
			// check data exists in cache
			if found {
				// Notify caller by notify channel
				job.CrawlerRet <- ret
			} else {
				// Sync two goroutine with sync.WaitGroup
				wg := sync.WaitGroup{}
				var ret1 data.ProductInfoList
				var ret2 data.ProductInfoList

				// Add in main func()
				wg.Add(1)
				go func() {
					// Done in goroutine
					defer wg.Done()
					ret1 = watsons.Crawler(job.QueryName, job.Timeout)
				}()

				// Add in main func()
				wg.Add(1)
				go func() {
					// Done in goroutine
					defer wg.Done()
					ret2 = poya.Crawler(job.QueryName, job.Timeout)
				}()

				// Waits two goroutine done
				wg.Wait()

				if ret1 != nil {
					ret = append(ret, ret1...) // append two slices
				}
				if ret2 != nil {
					ret = append(ret, ret2...) // append two slices
				}

				// save to cache
				cache.SetCache(job.QueryName, ret)

				// Notify caller by notify channel
				job.CrawlerRet <- ret
			}

			// Close notify channel
			close(job.CrawlerRet)

		case <-ctx.Done():
			glog.Infof("Terminate Worker ID: %v", id)
			exitSync.Done()
			return
		}
	}
}

// StartWorker initialize number of workers.
func StartWorker(sigAPIServerChan chan<- int) {
	glog.Infof("Start %d job workers", *maxWorkerNum)

	ctx, cancel := context.WithCancel(context.Background())
	for id := 1; id <= *maxWorkerNum; id++ {
		exitSync.Add(1)
		go jobWorker(ctx, id)
	}

	// To gracefully stop all services
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh:
			glog.Info("Receive os.Signal")
			cancel()
			return
		}
	}()
	exitSync.Wait()

	glog.Info("All workers shutdown")
	// Signal APIServer to shutdown when all workers stops
	sigAPIServerChan <- 1
}
