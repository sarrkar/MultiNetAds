package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/database"
	"github.com/sarrkar/chan-ta-net/panel/models"
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

	if result := ctrl.DB.Where(&models.Ad{Active: true}).Find(&ads); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &ads)
}

func (ctrl *AdController) IncImpression(ctx *gin.Context) {
	adId := ctx.Param("ad_id")
	var ad models.Ad

	if result := ctrl.DB.First(&ad, adId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ad.Impression++
	ctrl.DB.Save(&ad)

	ctx.JSON(http.StatusOK, &ad)
}

func (ctrl *AdController) IncClick(ctx *gin.Context) {
	adId := ctx.Param("ad_id")
	advId := ctx.Param("adv_id")
	pubId := ctx.Param("pub_id")

	var ad models.Ad
	var adv models.Advertiser
	var pub models.Publisher

	if result := ctrl.DB.First(&ad, adId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ad.Click++
	ctrl.DB.Save(&ad)

	if result := ctrl.DB.First(&adv, advId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	adv.Balance -= ad.BID
	ctrl.DB.Save(&adv)

	if result := ctrl.DB.First(&pub, pubId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	pub.Balance += (pub.CommissionPercent * ad.BID) / 100
	ctrl.DB.Save(&pub)

	ctx.Redirect(http.StatusMovedPermanently, ad.RedirectUrl)
}

func (ctrl *AdController) CreateMockData(ctx *gin.Context) {

	ctrl.DB.Exec("DELETE FROM ads")
	ctrl.DB.Exec("DELETE FROM advertisers")
	ctrl.DB.Exec("DELETE FROM publishers")

	advs := []models.Advertiser{
		{
			Name:    "adv1",
			Balance: 50000,
			Ads: []models.Ad{
				{
					Title:       "title1 adv1",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.google.com/",
					BID:         1000,
					Active:      true,
					Impression:  0,
					Click:       0,
				},
				{
					Title:       "title2 adv1",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         2000,
					Active:      true,
					Impression:  0,
					Click:       0,
				},
				{
					Title:       "title3 adv1",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         2000,
					Active:      false,
					Impression:  0,
					Click:       0,
				},
			},
		},
		{
			Name:    "adv2",
			Balance: 30000,
			Ads: []models.Ad{
				{
					Title:       "title1 adv2",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.google.com/",
					BID:         500,
					Active:      true,
					Impression:  0,
					Click:       0,
				},
				{
					Title:       "title2 adv2",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         1500,
					Active:      true,
					Impression:  0,
					Click:       0,
				},
				{
					Title:       "title3 adv2",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         1000,
					Active:      false,
					Impression:  0,
					Click:       0,
				},
			},
		},
	}

	pubs := []models.Publisher{
		{
			Name:              "pub1",
			Balance:           0,
			CommissionPercent: 20,
		},
		{
			Name:              "pub2",
			Balance:           15000,
			CommissionPercent: 15,
		},
		{
			Name:              "pub3",
			Balance:           1000,
			CommissionPercent: 10,
		},
		{
			Name:              "pub4",
			Balance:           0,
			CommissionPercent: 10,
		},
		{
			Name:              "pub5",
			Balance:           0,
			CommissionPercent: 30,
		},
	}

	ctrl.DB.Create(&advs)
	ctrl.DB.Create(&pubs)

	ctx.JSON(http.StatusOK, gin.H{
		"advs": &advs,
		"pubs": &pubs,
	})
}
