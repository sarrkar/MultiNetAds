package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: sync this with real `panel.local/api/ad/:id` api
type Response struct {
	Ad Ad
}

type Ad struct {
	Url string
}

func main() {


	fmt.Println("salam")
	router := gin.Default()


	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})


	keys := map[string]bool{} // use map as a set
	setValue := true

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"code":   200,
		})
	})

	router.GET("/impression/:id/:OTLkey", func(c *gin.Context) {
		id := c.Param("id")
		OTLkey := c.Param("OTLkey")
		if _, ok := keys[OTLkey]; !ok {
			keys[OTLkey] = setValue

			fmt.Println("call to panel.local/api/inc_impression/" + id)
			// resp, err := http.Get("panel.local/api/inc_impression/" + id)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// defer resp.Body.Close()
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"code":   200,
		})
	})

	router.GET("/click/:id/:OTLkey", func(c *gin.Context) {
		id := c.Param("id")
		OTLkey := c.Param("OTLkey")

		if _, ok := keys[OTLkey]; !ok {
			keys[OTLkey] = setValue

			fmt.Println("call to panel.local/api/inc_click/" + id)
			// resp, err := http.Get("panel.local/api/inc_click/" + id)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// defer resp.Body.Close()
		}

		// resp, err := http.Get("panel.local/api/ad/" + id)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer resp.Body.Close()
		// b, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)

		jsonTest := `
{
	"Ad": {
		"Url": "http://google.com"
	}
}`
		b := []byte(jsonTest)

		var res Response
		json.Unmarshal(b, &res)

		fmt.Println("redirect to " + res.Ad.Url)
		c.Redirect(http.StatusMovedPermanently, res.Ad.Url)
	})

	router.Run(":8080")
}
