// package main

// import (
//     "net/http"
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     router := gin.Default()

//     router.LoadHTMLGlob("templates/*")

//     router.GET("/template1", renderStaticTemplate("template1.html"))
//     router.GET("/template2", renderStaticTemplate("template2.html"))
//     router.GET("/template3", renderStaticTemplate("template3.html"))
//     router.GET("/template4", renderStaticTemplate("template4.html"))
//     router.GET("/template5", renderStaticTemplate("template5.html"))

//     router.Static("/static", "./static")

//     router.GET("/api/ad", func(c *gin.Context) {
//         ad := map[string]string{
//             "title": "This is a test ad",
//             "image": "/static/images/sample1.jpg",
//         }
//         c.JSON(http.StatusOK, ad)
//     })

//     router.Run(":8080")
// }

// func renderStaticTemplate(templateName string) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.HTML(http.StatusOK, templateName, nil)
//     }
// }

// package main

// import (
//     "github.com/gin-gonic/gin"
//     "net/http"
// )

// func main() {
//     router := gin.Default()

//     router.Static("/static", "./static")

//     router.LoadHTMLGlob("templates/*")

//     router.GET("/api/ad", func(c *gin.Context) {
//         ad := map[string]string{
//             "title": "This is a test ad from AdServer",
//             "image_url": "https://example.com/path/to/ad-image.jpg",
//         }
//         c.JSON(http.StatusOK, ad)
//     })

//     router.GET("/template1", func(c *gin.Context) {
//         c.HTML(http.StatusOK, "template1.html", nil)
//     })

//     router.GET("/template2", func(c *gin.Context) {
//         c.HTML(http.StatusOK, "template2.html", nil)
//     })

//     router.GET("/template3", func(c *gin.Context) {
//         c.HTML(http.StatusOK, "template3.html", nil)
//     })

//     router.Run(":8082")
// }

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdEvent struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/api/ad", func(c *gin.Context) {
		ad := map[string]string{
			"title":            "This is a test ad from AdServer",
			"image_url":        "http://localhost:8000/static/media/image12.jpg",
			"impression_event": "http://localhost:8080/impression/12/Y0JVOH0BQb",
			"click_event":      "http://localhost:8080/click/12/Y0JVOH0BQb",
		}
		c.JSON(http.StatusOK, ad)
	})

	router.POST("/api/impression", func(c *gin.Context) {
		var event AdEvent
		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		event.Type = "impression"
		sendEventToEventServer(event)

		c.JSON(http.StatusOK, gin.H{"status": "impression recorded"})
	})

	router.POST("/api/click", func(c *gin.Context) {
		var event AdEvent
		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		event.Type = "click"
		sendEventToEventServer(event)

		c.JSON(http.StatusOK, gin.H{"status": "click recorded"})
	})

	router.GET("/template1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", nil)
	})

	router.GET("/template2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template2.html", nil)
	})

	router.GET("/template3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template3.html", nil)
	})

	router.Run(":9001")
}

func sendEventToEventServer(event AdEvent) {
	jsonData, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling event:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/api/events", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending event to event server:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
	} else {
		log.Println("Event successfully sent to event server")
	}
}
