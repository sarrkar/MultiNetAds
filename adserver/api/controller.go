package api

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/adserver/client"
	"github.com/sarrkar/chan-ta-net/adserver/config"
	"github.com/sarrkar/chan-ta-net/common/helper"
)

type AdResponse struct {
	ID            uint   `json:"id"`
	AdvertiserID  uint   `json:"advertiser_id"`
	Title         string `json:"title"`
	ImageUrl      string `json:"image_url"`
	RedirectUrl   string `json:"redirect_url"`
	ImpressionUrl string `json:"impression_url"`
	ClickUrl      string `json:"click_url"`
}

func GetAd(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	publisherID, err := strconv.ParseUint(ctx.Query("publisher_id"), 10, 64)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid publisher_id"})
		return
	}
	title := ctx.Query("title")

	ads, publisher := client.GetBestAds(uint(publisherID), title)
	if len(ads) == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"publisher_id": publisherID, "publisher_category": publisher.Category, "ad_id": 0, "ad_category": "No match"})
		return
	}

	sort.Slice(ads, func(i, j int) bool {
		if ads[i].Impression != 0 && ads[j].Impression != 0 {
			return ads[i].BID*(ads[i].Click/ads[i].Impression) > ads[j].BID*(ads[j].Click/ads[j].Impression)
		}
		return ads[i].BID > ads[j].BID
	})
	ad := ads[0]

	ctx.IndentedJSON(http.StatusOK, AdResponse{
		ID:           ad.ID,
		AdvertiserID: ad.AdvertiserID,
		Title:        ad.Title,
		ImageUrl:     ad.ImageUrl,
		RedirectUrl:  ad.RedirectUrl,
		ImpressionUrl: fmt.Sprintf("%s/impression/%s",
			config.Config().Server.EventSeverExternalHost,
			helper.RandStr(config.Config().Server.OTLlength)),
		ClickUrl: fmt.Sprintf("%s/click/%s",
			config.Config().Server.EventSeverExternalHost,
			helper.RandStr(config.Config().Server.OTLlength)),
	})
}
