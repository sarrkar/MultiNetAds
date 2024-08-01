package controller

import (
	"github.com/sarrkar/chan-ta-net/panel/database"
	"github.com/sarrkar/chan-ta-net/panel/models"
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

func (ctrl *PublisherController) GetPublishers() (publishers []*models.Publisher, err error) {
	if result := ctrl.DB.Find(&publishers); result.Error != nil {
		return nil, result.Error
	}
	return
}

func (ctrl *PublisherController) GetPublisher(id uint) (publisher *models.Publisher, err error) {
	if result := ctrl.DB.Where(&models.Publisher{ID: id}).First(&publisher); result.Error != nil {
		return nil, result.Error
	}
	return
}

func (ctrl *PublisherController) GetPublisherWithAds(id uint) (publisher *models.Publisher, err error) {
	if result := ctrl.DB.Where(&models.Publisher{ID: id}).Preload("Ads").First(&publisher); result.Error != nil {
		return nil, result.Error
	}
	return
}

func (ctrl *PublisherController) NewPublisher(name, category string) (*models.Publisher, error) {
	var publisher models.Publisher
	if result := ctrl.DB.Where(&models.Publisher{Name: name, Category: category}).FirstOrCreate(&publisher); result.Error != nil {
		return nil, result.Error
	}
	return &publisher, nil
}
