package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/Panel/api/router"
	"github.com/sarrkar/Chan-ta-net/Panel/config"
)

func InitServer() {
	gin.SetMode(config.Config().Server.RunMode)
	r := gin.Default()

	r.LoadHTMLGlob("api/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	RegisterRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	adv := r.Group("/advertiser")
	{
		adv.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "advertiser.html", nil)
		})

		advAds := adv.Group("/ad")
		advReport := adv.Group("/report")
		advFinance := adv.Group("/finance")

		router.AdvertiserAd(advAds)
		router.AdvertiserReport(advReport)
		router.AdvertiserFinance(advFinance)
	}

	pub := r.Group("/publisher")
	{
		pub.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "publisher.html", nil)
		})

		pubPlace := pub.Group("/Place")
		pubReport := pub.Group("/report")
		pubFinance := pub.Group("/finance")

		router.PublisherPlace(pubPlace)
		router.PublisherReport(pubReport)
		router.PublisherFinance(pubFinance)
	}

	api := r.Group("/api")
	{
		Ad := api.Group("/ad")

		router.Ad(Ad)
	}

}
