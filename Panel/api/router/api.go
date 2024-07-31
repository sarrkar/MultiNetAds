package router

import (

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/api/controller"
)

func Ad(r *gin.RouterGroup) {
	ctrl := controller.NewAdController()
	r.Static("/uploads", "./uploads")

	r.GET("/all_ads", ctrl.GetAds)
	r.GET("/inc_impression/:ad_id", ctrl.IncImpression)
	r.GET("/inc_click/:ad_id/:adv_id/:pub_id", ctrl.IncClick)
	r.GET("/create_mock", ctrl.CreateMockData)
	// r.Static("/static", "./static")
}
