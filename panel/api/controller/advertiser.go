package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/common/helper"
	"github.com/sarrkar/chan-ta-net/common/models"
	"github.com/sarrkar/chan-ta-net/panel/config"
	"gorm.io/gorm"
)

type AdvertiserController struct {
	DB *gorm.DB
}

func NewAdvertiserController() *AdvertiserController {
	return &AdvertiserController{
		DB: database.GetDb(),
	}
}

func (ctrl AdvertiserController) Index(c *gin.Context) {
	var advertisers []*models.Advertiser
	if result := ctrl.DB.Order("ID").Find(&advertisers); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	c.HTML(http.StatusOK, "advertiser.html", gin.H{"Advertisers": advertisers})
}

func (ctrl AdvertiserController) Login(c *gin.Context) {
	name := c.PostForm("name")
	advertiser := &models.Advertiser{}
	if result := ctrl.DB.Where(&models.Advertiser{Name: name}).FirstOrCreate(advertiser); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/advertiser/%d", advertiser.ID))
}

func (ctrl AdvertiserController) Home(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("advertiser_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id := uint(intId)

	var advertiser *models.Advertiser
	if result := ctrl.DB.Where(&models.Advertiser{ID: id}).Preload("Ads").First(&advertiser); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := range advertiser.Ads {
		advertiser.Ads[i].Budget -= advertiser.Ads[i].BID * int(advertiser.Ads[i].Click)
	}

	sort.Slice(advertiser.Ads, func(i, j int) bool {
		return advertiser.Ads[i].ID < advertiser.Ads[j].ID
	})

	c.HTML(http.StatusOK, "adv_dashboard.html", gin.H{
		"Name":    advertiser.Name,
		"ID":      advertiser.ID,
		"Balance": advertiser.Balance,
		"Ads":     advertiser.Ads,
	})
}

func (ctrl AdvertiserController) AddCredit(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("advertiser_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id := uint(intId)

	amount, err := strconv.Atoi(c.PostForm("amount"))
	if err != nil || amount <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid BID value"})
		return
	}
	var advertiser *models.Advertiser
	if result := ctrl.DB.Where(&models.Advertiser{ID: id}).First(&advertiser); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	advertiser.Balance += amount
	ctrl.DB.Save(advertiser)
	c.Redirect(http.StatusFound, fmt.Sprintf("/advertiser/%d", advertiser.ID))
}

func (ctrl AdvertiserController) CreateAd(c *gin.Context) {
	title := c.PostForm("title")
	referralLink := c.PostForm("referral_link")
	category := c.PostForm("category")
	customCategory := c.PostForm("custom_category")

	file, err := c.FormFile("image_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	filename := helper.RandStr(10) + filepath.Base(file.Filename)
	filepath := filepath.Join(config.Config().Server.StaticDir, "uploads", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}
	imageLink := config.Config().Server.PanelExternalHost + "/static/uploads/" + filename

	if customCategory != "" {
		category = customCategory
	}

	intId, err := strconv.Atoi(c.Param("advertiser_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id := uint(intId)

	bid, err := strconv.Atoi(c.PostForm("click_amount"))
	if err != nil || bid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid BID value"})
		return
	}

	budget, err := strconv.Atoi(c.PostForm("budget"))
	if err != nil || budget <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Budget value"})
		return
	}
	var advertiser *models.Advertiser
	if result := ctrl.DB.Where(&models.Advertiser{ID: id}).Preload("Ads").First(&advertiser); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if budget > 2*advertiser.Balance {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Buget more than twice the balance"})
		return
	}

	ad := &models.Ad{
		Title:        title,
		BID:          bid,
		RedirectUrl:  referralLink,
		ImageUrl:     imageLink,
		Active:       true,
		AdvertiserID: advertiser.ID,
		Category:     category,
		Budget:       budget,
	}

	ctrl.DB.Create(ad)

	c.Redirect(http.StatusFound, fmt.Sprintf("/advertiser/%d", advertiser.ID))
}

func (ctrl AdvertiserController) ToggleAdStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("advertiser_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	adId := c.Param("ad_id")
	var ad models.Ad

	if result := ctrl.DB.Find(&ad, adId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ad.Active = !ad.Active
	ctrl.DB.Save(&ad)

	c.Redirect(http.StatusFound, fmt.Sprintf("/advertiser/%d", id))
}
