package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/Panel/database"
	"github.com/sarrkar/Chan-ta-net/Panel/models"
	"gorm.io/gorm"
)

type AdController struct {
	DB *gorm.DB
}

func NewAdController() *AdController {
	return &AdController{
		DB: database.GetDb(),
	}
}

func (ctrl *AdController) GetAds(ctx *gin.Context) {
	var ads []models.Ad

	// TODO: where(Ad.adviser.Balance) > 0
	if result := ctrl.DB.Where(&models.Ad{Active: true}).Find(&ads); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &ads)
}

func (ctrl *AdController) IncImpression(ctx *gin.Context) {
	id := ctx.Param("id")
	var ad models.Ad

	if result := ctrl.DB.First(&ad, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ad.Impression++
	ctrl.DB.Save(&ad)

	ctx.JSON(http.StatusOK, &ad)
}

func (ctrl *AdController) IncClick(ctx *gin.Context) {
	id := ctx.Param("id")
	var ad models.Ad

	if result := ctrl.DB.First(&ad, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ad.Click++
	ctrl.DB.Save(&ad)

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Redirect(http.StatusMovedPermanently, ad.RedirectUrl)
}

func (ctrl *AdController) CreateMockData(ctx *gin.Context) {

	ads := []models.Ad{
		{
			Title:       "title 1",
			ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
			RedirectUrl: "https://www.google.com/",
			BID:         1000,
			Active:      true,
			Impression:  0,
			Click:       0,
		},
		{
			Title:       "title 2",
			ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
			RedirectUrl: "https://www.varzesh3.com/",
			BID:         2000,
			Active:      true,
			Impression:  0,
			Click:       0,
		},
	}

	adv := models.Advertiser{
		Name:    "adv 1",
		Balance: 50000,
		Ads:     ads,
	}

	ctrl.DB.Create(&adv)

	ctx.JSON(http.StatusOK, &adv)
}
