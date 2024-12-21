package apiserver

import (
	v1 "compare-drugstore-price/apiserver/v1"
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	port = flag.Uint("port", 8080, "Server port should assigned")
)

// StartServer runs apiserver.
func StartServer(router *gin.Engine, sigAPIServerChan <-chan int) {
	router.GET("/v1/search", v1.Search)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}

	glog.Infof("Start APIServer at port: %d", *port)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Errorf("Shutting down the APIServer: %v", err)
		}
	}()

	// To gracefully stop all services
	<-sigAPIServerChan
	glog.Infof("Shutdown APIServer ...")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatalf("APIServer shutdown: %v", err)
	}
	glog.Info("Shutdown APIServer Done")
}
