package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/api/controller"
)

func Ad(r *gin.RouterGroup) {
	ctrl := controller.NewAdController()

	r.GET("/all_ads", ctrl.GetAds)
	r.GET("/inc_impression/:ad_id", ctrl.IncImpression)
	r.GET("/inc_click/:ad_id/:adv_id/:pub_id", ctrl.IncClick)
	r.GET("/create_mock", ctrl.CreateMockData)
}

func Pub(r *gin.RouterGroup) {
	pubctrl := controller.NewPublisherController()

	r.GET("/all_publishers", func(ctx *gin.Context) {
		result, err := pubctrl.GetPublishers()
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	})
}
