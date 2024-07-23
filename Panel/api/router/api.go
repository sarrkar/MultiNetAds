package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/Panel/api/controller"
)

func Ad(r *gin.RouterGroup) {
	ctrl := controller.NewAdController()

	r.POST("/", ctrl.Create)
	r.GET("/:id", ctrl.GetById)
	r.PUT("/:id", ctrl.Update)
	r.DELETE("/:id", ctrl.Delete)
}
