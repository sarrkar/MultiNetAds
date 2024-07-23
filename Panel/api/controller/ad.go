package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/Panel/database"
	"gorm.io/gorm"
)

type AdController struct {
	Db *gorm.DB
}

func NewAdController() *AdController {
	return &AdController{
		Db: database.GetDb(),
	}
}

func (ctrl *AdController) Update(ctx *gin.Context) {
	// TODO
}

func (ctrl *AdController) Create(ctx *gin.Context) {
	// TODO
}

func (ctrl *AdController) Delete(ctx *gin.Context) {
	// TODO
}

func (ctrl *AdController) GetById(ctx *gin.Context) {
	// TODO
}
