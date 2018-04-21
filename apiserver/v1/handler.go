package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/hsiaoairplane/compare-drugstore-price/job"
)

// Search implements the RESTful GET API.
func Search(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		glog.Error("query name is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameter"})
		return
	}

	// Create channel
	notifyChan := make(chan bool, 1)
	ji := job.JobInfo{QueryName: name, NotifyChan: notifyChan}

	// Enqueue query job
	job.Enqueue(ji)

	// Wait job response
	<-notifyChan

	// Close channel
	close(notifyChan)
}
