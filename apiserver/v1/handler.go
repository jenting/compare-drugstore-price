package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/jenting/compare-drugstore-price/data"
	"github.com/jenting/compare-drugstore-price/job"
)

const (
	sortAscend  = "ascend"
	sortDescend = "descend"
)

// Search implements the RESTful GET API.
func Search(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		glog.Error("Query name is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameter"})
		return
	}
	sort := c.Query("sort")
	if sort != "" && (sort != sortAscend && sort != sortDescend) {
		glog.Error("Invalid sort parameter")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid sort parameter"})
		return
	}

	job := job.CreateNewJob(name)

	// Wait job response
	retJSON := <-job.CrawlerRet

	switch sort {
	case sortAscend:
		data.SortAscendByPrice(retJSON)
	case sortDescend:
		data.SortDescendByPrice(retJSON)
	}

	c.JSON(http.StatusOK, retJSON)
}
