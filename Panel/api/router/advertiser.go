package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sarrkar/chan-ta-net/panel/api/controller"
	"github.com/sarrkar/chan-ta-net/panel/models"
)

func advertiserAd(r *gin.RouterGroup) {
	ctrl := controller.NewAdvertiserController()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "my_ads.html", nil)
	})

	r.GET("/list", func(c *gin.Context) {
		intId, err := strconv.Atoi(c.Param("advertiser_id"))
		id := uint(intId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		advertiser, err := ctrl.GetAdvertiserWithAds(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.HTML(http.StatusOK, "list_advertisement.html", gin.H{"Ads": advertiser.Ads})
	})

	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_advertisement.html", nil)
	})

	r.POST("/create", func(c *gin.Context) {
		title := c.PostForm("title")
		referralLink := c.PostForm("referral_link")
		imageLink := c.PostForm("image_link")

		intId, err := strconv.Atoi(c.Param("advertiser_id"))
		id := uint(intId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		bid, err := strconv.Atoi(c.PostForm("click_amount"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		advertiser, err := ctrl.GetAdvertiser(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		ad := &models.Ad{
			Title:        title,
			BID:          bid,
			RedirectUrl:  referralLink,
			ImageUrl:     imageLink,
			Active:       true,
			AdvertiserID: advertiser.ID,
		}

		ctrl.DB.Create(ad)

		c.Redirect(http.StatusFound, "/advertiser/"+strconv.Itoa(int(advertiser.ID))+"/ads/list")
	})

}

func advertiserReport(r *gin.RouterGroup) {

}

func advertiserFinance(r *gin.RouterGroup) {
	ctrl := controller.NewAdvertiserController()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "my_finance.html", nil)
	})

	r.GET("/balance", func(c *gin.Context) {
		intId, err := strconv.Atoi(c.Param("advertiser_id"))
		id := uint(intId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		advertiser, err := ctrl.GetAdvertiser(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.HTML(http.StatusOK, "balance.html", gin.H{"Name": advertiser.Name, "Balance": advertiser.Balance})
	})

	r.GET("/payment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "payment.html", nil)
	})

	r.POST("/add-credit", func(c *gin.Context) {
		intId, err := strconv.Atoi(c.Param("advertiser_id"))
		id := uint(intId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		
		amount, err := strconv.Atoi(c.PostForm("amount"))

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		advertiser, err := ctrl.GetAdvertiser(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		advertiser.Balance += amount
		ctrl.DB.Save(advertiser)

		c.Redirect(http.StatusFound, "/advertiser/"+strconv.Itoa(int(advertiser.ID))+"/finance/balance")
	})

}

func Advertiser(r *gin.RouterGroup) {
	ctrl := controller.NewAdvertiserController()

	r.GET("/", func(c *gin.Context) {
		advertisers, err := ctrl.GetAdvertisers()
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.HTML(http.StatusOK, "advertiser.html", gin.H{"Advertisers": advertisers})
	})

	r.POST("/submit-name", func(c *gin.Context) {
		name := c.PostForm("name")
		advertiser, err := ctrl.NewAdvertiser(name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Redirect(http.StatusFound, "/advertiser/"+strconv.Itoa(int(advertiser.ID)))
	})

	r.GET("/:advertiser_id", func(c *gin.Context) {
		intId, err := strconv.Atoi(c.Param("advertiser_id"))
		id := uint(intId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		advertiser, err := ctrl.GetAdvertiser(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.HTML(http.StatusOK, "adver_dashboard.html", gin.H{"Name": advertiser.Name, "ID": advertiser.ID})
	})

	advAds := r.Group("/:advertiser_id/ads")
	advReport := r.Group("/:advertiser_id/reports")
	advFinance := r.Group("/:advertiser_id/finance")

	advertiserAd(advAds)
	advertiserReport(advReport)
	advertiserFinance(advFinance)
}






