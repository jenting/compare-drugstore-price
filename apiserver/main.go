package apiserver

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/hsiaoairplane/compare-drugstore-price/apiserver/v1"
)

var (
	port = flag.Uint("port", 8080, "Server port should assigned")
)

// StartServer runs apiserver.
func StartServer(router *gin.Engine) {
	router.GET("/v1/search", v1.Search)

	glog.Infof("Start APIServer at port: %d", *port)
	router.Run(fmt.Sprintf(":%d", *port))
}
