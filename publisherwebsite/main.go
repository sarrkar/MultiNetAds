package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/publisherwebsite/config"
)

func main() {
	gin.SetMode(config.Config().Server.RunMode)
	r := gin.Default()

	r.LoadHTMLGlob(config.Config().Server.TemplateDir)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/template/:publisher_id", func(c *gin.Context) {
		publisherID := c.Param("publisher_id")
		title := "default"

		if publisherID == "1" {
			title = "تکنولایف تکنولوژی در یک قدمی تو"
		} else if publisherID == "2" {
			title = "آموزش ریاضی سوم راهنمایی برای کودکان استثنایی"
		} else if publisherID == "3" {
			title = "ورزش 3 برای سلامتی و هیجان"
		}

		c.HTML(http.StatusOK, "template.html", gin.H{"PublisherID": publisherID, "Title": title})
	})

	err := r.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
