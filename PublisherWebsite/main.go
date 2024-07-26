package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/publisher-website/config"
)

func main() {
	router := gin.Default()
	gin.SetMode(config.Config().Server.RunMode)

	router.LoadHTMLGlob(config.Config().Server.TemplateDir)

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

	err := router.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
