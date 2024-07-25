package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/event-server/client"
	"github.com/sarrkar/chan-ta-net/event-server/helper"
)

var set helper.Set = helper.NewSet()

func ClickHandler(ctx *gin.Context) {
	otlKey := ctx.Param("OTL_key")
	adID, _ := ctx.GetQuery("ad_id")
	advID, _ := ctx.GetQuery("adv_id")
	pubID, _ := ctx.GetQuery("pub_id")
	redirectURL, _ := ctx.GetQuery("redirect_url")

	if !set.Check(otlKey) {
		set.Add(otlKey)
		go client.AddClick(adID, advID, pubID)
	}

	ctx.Redirect(http.StatusMovedPermanently, redirectURL)
}

func ImpressionHandler(ctx *gin.Context) {
	otlKey := ctx.Param("OTL_key")
	adID, _ := ctx.GetQuery("ad_id")
	advID, _ := ctx.GetQuery("adv_id")
	pubID, _ := ctx.GetQuery("pub_id")

	if !set.Check(otlKey) {
		set.Add(otlKey)
		go client.AdImperession(adID, advID, pubID)
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"code":   200,
	})
}
