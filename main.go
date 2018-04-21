package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/hsiaoairplane/compare-price-drugstore/apiserver"
	"github.com/hsiaoairplane/compare-price-drugstore/job"
)

func init() {
	// Default logging to console.
	flag.Set("logtostderr", "true")
}

func main() {
	// Parse flags.
	flag.Parse()

	// Disable debug mode of gin framework.
	gin.SetMode(gin.ReleaseMode)

	// Disable console color.
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Start job worker to de-job-queue.
	job.StartWorker()
	// Start APIServer
	apiserver.StartServer(router)
}
