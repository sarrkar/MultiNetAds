package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AdvertiserAd(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "my_ads.html", nil)
	})

	r.GET("/add_advertisement", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_advertisement.html", nil)
	})

	r.POST("/submit-advertisement", func(c *gin.Context) {
	//	title := c.PostForm("title")
	//	clickAmount := c.PostForm("click_amount")
	//	description := c.PostForm("description")

		c.HTML(http.StatusOK, "submit-advertisement.html", gin.H{
			"message": "تبلیغ با موفقیت اضافه شد!",
		})
	})


}

func AdvertiserReport(r *gin.RouterGroup) {
	// TODO
}

func AdvertiserFinance(r *gin.RouterGroup) {
	// TODO
}
