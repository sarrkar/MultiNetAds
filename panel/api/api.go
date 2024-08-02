package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/api/controller"
	"github.com/sarrkar/chan-ta-net/panel/config"
)

func InitServer() {
	gin.SetMode(config.Config().Server.RunMode)
	r := gin.Default()

	r.Static("/static", config.Config().Server.StaticDir)
	r.LoadHTMLGlob(config.Config().Server.TemplateDir)

	RegisterRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })

	adv := r.Group("/advertiser")
	{
		ctrl := controller.NewAdvertiserController()

		adv.GET("/", ctrl.Index)
		adv.POST("/login", ctrl.Login)
		adv.GET("/:advertiser_id", ctrl.Home)
		adv.POST("/:advertiser_id/add_credit", ctrl.AddCredit)
		adv.POST("/:advertiser_id/ad_create", ctrl.CreateAd)
	}

	pub := r.Group("/publisher")
	{
		ctrl := controller.NewPublisherController()

		pub.GET("/", ctrl.Index)
		pub.POST("/login", ctrl.Login)
		pub.GET("/:publisher_id", ctrl.Home)
		pub.POST("/:publisher_id/checkout", ctrl.Checkout)

	}

	api := r.Group("/api")
	{
		ctrl := controller.NewApiController()

		api.GET("/inc_impression/:ad_id/:adv_id/:pub_id", ctrl.IncImpression)
		api.GET("/inc_click/:ad_id/:adv_id/:pub_id", ctrl.IncClick)
		api.GET("/create_mock", ctrl.CreateMockData)
		api.POST("/toggle_status/:ad_id", ctrl.ToggleAdStatus)
		api.GET("/all_ads", ctrl.GetAds)
		api.GET("/all_publishers", ctrl.GetPubs)
		api.GET("/all_advertisers", ctrl.GetAdvs)
	}

}
