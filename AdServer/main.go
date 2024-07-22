package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ads []Ad
}

type Ad struct {
	ID       uint
	Title    string
	ImageUrl string
	BID      uint
}

type AdResponse struct {
	Title         string
	ImageUrl      string
	ImpressionUrl string
	ClickUrl      string
}

func randStr(length int) string {
	b := make([]byte, length)
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	// resp, err := http.Get("panel.local/api/ads")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	jsonTest := `
{
    "ads": [
        {
            "ID": 12,
            "Title": "title1",
            "ImageUrl": "image.storage/media/image12.jpg",
            "BID": 100
        },
        {
            "ID": 14,
            "Title": "title2",
            "ImageUrl": "image.storage/media/image14.jpg",
            "BID": 150
        }
    ]
}`
	b := []byte(jsonTest)

	var res Response
	json.Unmarshal(b, &res)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"code":   200,
		})
	})

	router.GET("/api/ad", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"code":   200,
			"data": AdResponse{
				Title:         res.Ads[0].Title,
				ImageUrl:      res.Ads[0].ImageUrl,
				ImpressionUrl: fmt.Sprintf("eventserver.local/impression/%d/%s", res.Ads[0].ID, randStr(10)),
				ClickUrl:      fmt.Sprintf("eventserver.local/click/%d/%s", res.Ads[0].ID, randStr(10)),
			},
		})
	})
	router.Run(":8080")
}
