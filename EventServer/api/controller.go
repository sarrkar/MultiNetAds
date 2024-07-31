package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/event-server/client"
	"github.com/sarrkar/chan-ta-net/event-server/helper"
)

var set helper.Set = helper.NewSet()
var clt client.Client = client.NewKafkaClinet()

func ClickHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	otlKey := ctx.Param("OTL_key")
	adID, adOK := ctx.GetQuery("ad_id")
	advID, advOK := ctx.GetQuery("adv_id")
	pubID, pubOK := ctx.GetQuery("pub_id")
	redirectURL, redirectOK := ctx.GetQuery("redirect_url")

	if !set.Check(otlKey) {
		set.Add(otlKey)
		if adOK && advOK && pubOK {
			go clt.AddClick(adID, advID, pubID)
		}
	}

	if redirectOK {
		ctx.Redirect(http.StatusMovedPermanently, redirectURL)
	} else {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
			"code":   400,
		})
	}
}

func ImpressionHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	otlKey := ctx.Param("OTL_key")
	adID, _ := ctx.GetQuery("ad_id")
	advID, _ := ctx.GetQuery("adv_id")
	pubID, _ := ctx.GetQuery("pub_id")

	if !set.Check(otlKey) {
		set.Add(otlKey)
		go clt.AddImperession(adID, advID, pubID)
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"code":   200,
	})
}
