package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/ad-server/client"
	"github.com/sarrkar/chan-ta-net/ad-server/config"
	"github.com/sarrkar/chan-ta-net/ad-server/helper"
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
	ad := client.GetBestAds()
	ctx.IndentedJSON(http.StatusOK, AdResponse{
		ID:           ad.ID,
		AdvertiserID: ad.AdvertiserID,
		Title:        ad.Title,
		ImageUrl:     ad.ImageUrl,
		RedirectUrl:  ad.RedirectUrl,
		ImpressionUrl: fmt.Sprintf("%s/impression/%s", config.Config().Server.EventSeverExternalHost,
			helper.RandStr(config.Config().Server.OTLlength)),
		ClickUrl: fmt.Sprintf("%s/click/%s", config.Config().Server.EventSeverExternalHost,
			helper.RandStr(config.Config().Server.OTLlength)),
	})
}
