package main

import (
	"github.com/gin-gonic/gin"
	"yektanet.com/displayads/panel/api"
)

func main() {
	r := gin.Default()

	r.POST("/advertiser/create_ad", api.CreateAd)
	r.GET("/advertiser/list_ads", api.ListAds)
	r.POST("/advertiser/activate_ad/:id", api.ActivateAd)
	r.POST("/advertiser/charge", api.ChargeAdvertiser)
	r.GET("/advertiser/view_stats/:id", api.ViewAdStats)
	r.GET("/advertiser/redirect/:id", api.RedirectAd)

	r.POST("/publisher/withdraw", api.WithdrawPublisher)
	r.GET("/publisher/view_stats/:id", api.ViewPublisherStats)

	r.Run(":8080")
}
