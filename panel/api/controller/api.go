package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/common/models"
	"gorm.io/gorm"
)

type ApiController struct {
	DB *gorm.DB
}

func NewApiController() *ApiController {
	return &ApiController{
		DB: database.GetDb(),
	}
}

func (ctrl *ApiController) GetAds(c *gin.Context) {
	var ads []models.Ad

	if result := ctrl.DB.Where(&models.Ad{Active: true}).Where("budget - click * bid > ?", 0).Find(&ads); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &ads)
}

func (ctrl *ApiController) GetPubs(c *gin.Context) {
	var publishers []*models.Publisher
	if result := ctrl.DB.Find(&publishers); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	c.JSON(http.StatusOK, publishers)
}

func (ctrl *ApiController) GetAdvs(c *gin.Context) {
	var advertisers []*models.Advertiser
	if result := ctrl.DB.Find(&advertisers); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	c.JSON(http.StatusOK, advertisers)
}

func (ctrl *ApiController) IncImpression(c *gin.Context) {
	adId := c.Param("ad_id")
	pubId := c.Param("pub_id")
	var ad models.Ad
	var pub models.Publisher

	if result := ctrl.DB.First(&ad, adId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ad.Impression++
	ctrl.DB.Save(&ad)

	if result := ctrl.DB.First(&pub, pubId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	pub.Impression++
	ctrl.DB.Save(&pub)

	c.JSON(http.StatusOK, &ad)
}

func (ctrl *ApiController) IncClick(c *gin.Context) {
	adId := c.Param("ad_id")
	advId := c.Param("adv_id")
	pubId := c.Param("pub_id")
	fmt.Printf("click %s %s %s", adId, advId, pubId)

	var ad models.Ad
	var adv models.Advertiser
	var pub models.Publisher

	if result := ctrl.DB.First(&ad, adId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ad.Click++
	ctrl.DB.Save(&ad)

	if result := ctrl.DB.First(&adv, advId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	adv.Balance -= ad.BID
	ctrl.DB.Save(&adv)

	if result := ctrl.DB.First(&pub, pubId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	pub.Balance += (pub.CommissionPercent * ad.BID) / 100
	pub.Click++
	ctrl.DB.Save(&pub)

	c.Redirect(http.StatusMovedPermanently, ad.RedirectUrl)
}

func (ctrl *ApiController) ToggleAdStatus(c *gin.Context) {
	adId := c.Param("ad_id")
	var ad models.Ad

	if result := ctrl.DB.Find(&ad, adId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ad.Active = !ad.Active
	ctrl.DB.Save(&ad)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ctrl *ApiController) CreateMockData(c *gin.Context) {

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
					Category:    "education",
					Budget:      20000,
				},
				{
					Title:       "title2 adv1",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         2000,
					Active:      true,
					Category:    "technology",
					Budget:      20000,
				},
				{
					Title:       "title3 adv1",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         2000,
					Active:      false,
					Category:    "education",
					Budget:      20000,
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
					Category:    "technology",
					Budget:      20000,
				},
				{
					Title:       "title2 adv2",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         1500,
					Active:      true,
					Category:    "education",
					Budget:      20000,
				},
				{
					Title:       "title3 adv2",
					ImageUrl:    "https://letsenhance.io/static/73136da51c245e80edc6ccfe44888a99/1015f/MainBefore.jpg",
					RedirectUrl: "https://www.varzesh3.com/",
					BID:         1000,
					Active:      false,
					Category:    "technology",
					Budget:      20000,
				},
			},
		},
	}

	pubs := []models.Publisher{
		{
			Name:              "technolife",
			Balance:           0,
			CommissionPercent: 20,
			Category:          "technology",
		},
		{
			Name:              "faradars",
			Balance:           15000,
			CommissionPercent: 15,
			Category:          "education",
		},
		{
			Name:              "varzesh3",
			Balance:           1000,
			CommissionPercent: 10,
			Category:          "health",
		},
		{
			Name:              "zoomit",
			Balance:           0,
			CommissionPercent: 10,
			Category:          "technology",
		},
		{
			Name:              "quera",
			Balance:           0,
			CommissionPercent: 30,
			Category:          "education",
		},
	}

	ctrl.DB.Create(&advs)
	ctrl.DB.Create(&pubs)

	c.JSON(http.StatusOK, gin.H{
		"advs": &advs,
		"pubs": &pubs,
	})
}
