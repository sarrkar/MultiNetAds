package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yektanet.com/displayads/panel/model"
)

func CreateAd(c *gin.Context) {
	var ad model.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.Lock()
	defer model.Unlock()
	ad.ID = model.NextAdID
	model.NextAdID++
	model.Ads[ad.ID] = ad

	c.JSON(http.StatusCreated, ad)
}

func ListAds(c *gin.Context) {
	model.Lock()
	defer model.Unlock()
	adList := make([]model.Ad, 0, len(model.Ads))
	for _, ad := range model.Ads {
		adList = append(adList, ad)
	}
	c.JSON(http.StatusOK, adList)
}

func ActivateAd(c *gin.Context) {
	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ad ID"})
		return
	}

	model.Lock()
	defer model.Unlock()
	ad, exists := model.Ads[adID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ad not found"})
		return
	}

	ad.Active = !ad.Active
	model.Ads[adID] = ad

	c.JSON(http.StatusOK, ad)
}

func ChargeAdvertiser(c *gin.Context) {
	type ChargeRequest struct {
		AdvertiserID int `json:"advertiser_id"`
		Amount       int `json:"amount"`
	}

	var req ChargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.Lock()
	defer model.Unlock()

	advertiser, exists := model.Advertisers[req.AdvertiserID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Advertiser not found"})
		return
	}

	advertiser.Balance += req.Amount
	model.Advertisers[req.AdvertiserID] = advertiser

	c.JSON(http.StatusOK, advertiser)
}

func ViewAdStats(c *gin.Context) {
	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ad ID"})
		return
	}

	model.Lock()
	defer model.Unlock()
	ad, exists := model.Ads[adID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ad not found"})
		return
	}

	stats := gin.H{
		"impressions": ad.Impressions,
		"clicks":      ad.Clicks,
		"ctr":         float64(ad.Clicks) / float64(ad.Impressions),
		"spent":       ad.Spent,
	}

	c.JSON(http.StatusOK, stats)
}

func RedirectAd(c *gin.Context) {
	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ad ID"})
		return
	}

	model.Lock()
	defer model.Unlock()
	ad, exists := model.Ads[adID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ad not found"})
		return
	}

	ad.Clicks++
	model.Ads[adID] = ad
	c.Redirect(http.StatusFound, ad.Link)
}
