package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/Panel/api/router"
	"github.com/sarrkar/Chan-ta-net/Panel/config"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	RegisterRoutes(r, cfg)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	adv := r.Group("/advertiser")
	{
		advAds := adv.Group("/ad")
		advReport := adv.Group("/report")
		advFinance := adv.Group("/finance")

		router.AdvertiserAd(advAds, cfg)
		router.AdvertiserReport(advReport, cfg)
		router.AdvertiserFinance(advFinance, cfg)
	}

	pub := r.Group("/publisher")
	{
		pubPlace := pub.Group("/Place")
		pubReport := pub.Group("/report")
		pubFinance := pub.Group("/finance")

		router.PublisherPlace(pubPlace, cfg)
		router.PublisherReport(pubReport, cfg)
		router.PublisherFinance(pubFinance, cfg)
	}

	api := r.Group("/api")
	{
		Ad := api.Group("/ad")

		router.Ad(Ad, cfg)
	}

}
