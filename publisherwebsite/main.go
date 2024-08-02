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
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Publishers": []any{
				map[string]any{"name": "technolife", "id": "1"},
				map[string]any{"name": "faradars", "id": "2"},
				map[string]any{"name": "varzesh3", "id": "3"},
				map[string]any{"name": "zoomit", "id": "4"},
				map[string]any{"name": "quera", "id": "5"},
			},
		})
	})

	r.GET(
		"/template/:publisher_id",
		func(c *gin.Context) {
			publisherID := c.Param("publisher_id")
			title := "default"

			if publisherID == "1" {
				title = "تکنولایف - تکنولوژی در یک قدمی تو"
			} else if publisherID == "2" {
				title = "آموزش ریاضی سوم راهنمایی برای کودکان استثنایی - فرادرس"
			} else if publisherID == "3" {
				title = "ورزش 3 - برای سلامتی و هیجان"
			} else if publisherID == "4" {
				title = "اخبار گوشی موبایل بررسی جدید ترین گوشی های بازار - زومیت"
			} else if publisherID == "5" {
				title = "از آموزش و تمرین برنامه نویسی - کوئرا"
			}

			c.HTML(http.StatusOK, "template.html", gin.H{"PublisherID": publisherID, "Title": title})
		},
	)

	err := r.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
