package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/adserver/config"
)

func InitServer() {
	gin.SetMode(config.Config().Server.RunMode)
	r := gin.Default()

	RegisterRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/ad", GetAd)

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "AD SERVER OK",
			"code":   200,
		})
	})

}
