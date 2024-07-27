package router

import (
	"net/http"
//	"strconv"

	"github.com/gin-gonic/gin"
//	"github.com/sarrkar/chan-ta-net/panel/api/controller"
//	"github.com/sarrkar/chan-ta-net/panel/models"
)

func PublisherPlace(r *gin.RouterGroup) {
	// TODO
}

func PublisherReport(r *gin.RouterGroup) {
	// TODO
}

func PublisherFinance(r *gin.RouterGroup) {
	// TODO
}


func Publisher(r *gin.RouterGroup) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "publisher_name.html", nil)
	})
	pubPlace := r.Group("/place")
	pubReport := r.Group("/report")
	pubFinance := r.Group("/finance")

	PublisherPlace(pubPlace)
	PublisherReport(pubReport)
	PublisherFinance(pubFinance)
}