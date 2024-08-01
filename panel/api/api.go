package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/api/router"
	"github.com/sarrkar/chan-ta-net/panel/config"
)

func InitServer() {
	gin.SetMode(config.Config().Server.RunMode)
	//	r := gin.New()
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/static", config.Config().Server.StaticDir)

	r.LoadHTMLGlob(config.Config().Server.TemplateDir)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	RegisterRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", config.Config().Server.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	adv := r.Group("/advertiser")
	{
		router.Advertiser(adv)
	}

	pub := r.Group("/publisher")
	{
		router.Publisher(pub)

	}

	api := r.Group("/api")
	{
		Ad := api.Group("/ad")
		Pub := api.Group("/publisher")
		router.Ad(Ad)
		router.Pub(Pub)
	}

}
