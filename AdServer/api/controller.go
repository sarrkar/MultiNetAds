package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/AdServer/client"
	"github.com/sarrkar/Chan-ta-net/AdServer/config"
	"github.com/sarrkar/Chan-ta-net/AdServer/helper"
)

type AdResponse struct {
	Title         string `json:"title"`
	ImageUrl      string `json:"image_url"`
	ImpressionUrl string `json:"impression_event"`
	ClickUrl      string `json:"click_event"`
}

func GetAd(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ad := client.GetBestAds()
	ctx.IndentedJSON(http.StatusOK, AdResponse{
		Title:    ad.Title,
		ImageUrl: ad.ImageUrl,
		ImpressionUrl: fmt.Sprintf("%s/click/%d/%s", config.Config().Server.EventSeverHost,
			ad.ID, helper.RandStr(config.Config().Server.OTLlength)),
		ClickUrl: fmt.Sprintf("%s/impression/%d/%s", config.Config().Server.EventSeverHost,
			ad.ID, helper.RandStr(config.Config().Server.OTLlength)),
	})
}
