package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/template1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", nil)
	})

	router.GET("/template2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template2.html", nil)
	})

	router.GET("/template3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template3.html", nil)
	})

	router.GET("/template4", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template4.html", nil)
	})

	router.GET("/template5", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template5.html", nil)
	})

	router.Run(":9001")
}
