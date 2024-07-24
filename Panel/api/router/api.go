package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/Panel/api/controller"
)

func Ad(r *gin.RouterGroup) {
	ctrl := controller.NewAdController()

	r.GET("/all_ads", ctrl.GetAds)
	r.GET("/inc_impression/:id", ctrl.IncImpression)
	r.GET("/inc_click/:id", ctrl.IncClick)
	r.GET("/create_mock", ctrl.CreateMockData)
}
