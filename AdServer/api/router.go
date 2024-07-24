package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/Chan-ta-net/AdServer/config"
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
			"status": "OK",
			"code":   200,
		})
	})

}