package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/common/database"
	"github.com/sarrkar/chan-ta-net/common/models"
	"github.com/sarrkar/chan-ta-net/panel/config"
	"gorm.io/gorm"
)

type PublisherController struct {
	DB *gorm.DB
}

func NewPublisherController() *PublisherController {
	return &PublisherController{
		DB: database.GetDb(),
	}
}

func (ctrl PublisherController) Index(c *gin.Context) {
	var publishers []*models.Publisher
	if result := ctrl.DB.Order("ID").Find(&publishers); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	c.HTML(http.StatusOK, "publisher.html", gin.H{"Publishers": publishers})
}

func (ctrl PublisherController) Login(c *gin.Context) {
	name := c.PostForm("name")
	// commissionPercent, err := strconv.Atoi(c.PostForm("commission_percent"))
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	// if commissionPercent <= 0 || commissionPercent >= 100 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid commissionPercent value"})
	// 	return
	// }

	category := c.PostForm("category")
	customCategory := c.PostForm("custom_category")
	if customCategory != "" {
		category = customCategory
	}

	var publisher models.Publisher
	if result := ctrl.DB.Where(&models.Publisher{Name: name}).First(&publisher); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			publisher = models.Publisher{Name: name, CommissionPercent: 20, Category: category}
			ctrl.DB.Create(&publisher)
		} else {
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/publisher/%d", publisher.ID))
}

func (ctrl PublisherController) Home(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("publisher_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id := uint(uid)

	var publisher *models.Publisher
	if result := ctrl.DB.Where(&models.Publisher{ID: id}).First(&publisher); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.HTML(http.StatusOK, "pub_dashboard.html", gin.H{
		"Name":              publisher.Name,
		"ID":                publisher.ID,
		"Balance":           publisher.Balance,
		"Click":             publisher.Click,
		"Impression":        publisher.Impression,
		"Category":          publisher.Category,
		"CommissionPercent": publisher.CommissionPercent,
		"ScriptApi":         config.Config().Server.AdSeverExternalAPI,
	})
}

func (ctrl PublisherController) Checkout(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("publisher_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id := uint(uid)

	var publisher *models.Publisher
	if result := ctrl.DB.Where(&models.Publisher{ID: id}).First(&publisher); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	publisher.Balance = 0
	ctrl.DB.Save(publisher)
	c.Redirect(http.StatusFound, fmt.Sprintf("/publisher/%d", publisher.ID))
}
