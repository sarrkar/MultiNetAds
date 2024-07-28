package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/api/controller"
	"github.com/sarrkar/chan-ta-net/panel/config"
)

func PublisherPlace(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ad-fetcher.html", gin.H{
			"api":    config.Config().Server.AdSeverExternalAPI,
			"pub_id": c.Param("publisher_id"),
		})
	})
}

func PublisherReport(r *gin.RouterGroup) {
	// TODO
}

func PublisherFinance(r *gin.RouterGroup) {
	ctrl := controller.NewPublisherController()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pub_finance.html", nil)
	})

	r.GET("/balance", func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("publisher_id"))
		id := uint(uid)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		publisher, err := ctrl.GetPublisher(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.HTML(http.StatusOK, "balance.html", gin.H{"Name": publisher.Name, "Balance": publisher.Balance})
	})

	r.GET("/checkout", func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("publisher_id"))
		id := uint(uid)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		publisher, err := ctrl.GetPublisher(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		publisher.Balance = 0
		ctrl.DB.Save(publisher)

		c.HTML(http.StatusOK, "checkout.html", nil)
	})
}

func Publisher(r *gin.RouterGroup) {
	ctrl := controller.NewPublisherController()

	r.GET("/", func(c *gin.Context) {
		publishers, err := ctrl.GetPublishers()
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.HTML(http.StatusOK, "publisher.html", gin.H{"Publishers": publishers})
	})

	r.POST("/submit-name", func(c *gin.Context) {
		name := c.PostForm("name")
		publisher, err := ctrl.NewPublisher(name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusFound, "/publisher/"+strconv.Itoa(int(publisher.ID)))
	})

	r.GET("/:publisher_id", func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("publisher_id"))
		id := uint(uid)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		publisher, err := ctrl.GetPublisher(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.HTML(http.StatusOK, "pub_dashboard.html", gin.H{"Name": publisher.Name, "ID": publisher.ID})
	})

	pubPlace := r.Group("/:publisher_id/add-script")
	pubReport := r.Group("/:publisher_id/reports")
	pubFinance := r.Group("/:publisher_id/finance")

	PublisherPlace(pubPlace)
	PublisherReport(pubReport)
	PublisherFinance(pubFinance)
}
