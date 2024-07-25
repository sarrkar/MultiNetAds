package controller

import (
	"sync"

	"github.com/sarrkar/chan-ta-net/panel/database"
	"github.com/sarrkar/chan-ta-net/panel/models"
	"gorm.io/gorm"
)

type AdvertiserController struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewAdvertiserController() *AdvertiserController {
	return &AdvertiserController{
		DB: database.GetDb(),
	}
}

func (ctrl *AdvertiserController) GetAdvertisers() (advertisers []*models.Advertiser, err error) {
	if result := ctrl.DB.Find(&advertisers); result.Error != nil {
		return nil, result.Error
	}
	return
}

func (ctrl *AdvertiserController) GetAdvertiser(id int) (advertiser *models.Advertiser, err error) {
	if result := ctrl.DB.Where(&models.Advertiser{ID: id}).First(&advertiser); result.Error != nil {
		return nil, result.Error
	}
	return
}

func (ctrl *AdvertiserController) GetAdvertiserWithAds(id int) (advertiser *models.Advertiser, err error) {
	if result := ctrl.DB.Where(&models.Advertiser{ID: id}).Preload("Ads").First(&advertiser); result.Error != nil {
		return nil, result.Error
	}
	return
}

func (ctrl *AdvertiserController) NewAdvertiser(name string) (*models.Advertiser, error) {
	advertiser := &models.Advertiser{}
	if result := ctrl.DB.Where(&models.Advertiser{Name: name}).FirstOrCreate(advertiser); result.Error != nil {
		return nil, result.Error
	}
	return advertiser, nil
}
