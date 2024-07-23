package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yektanet.com/displayads/panel/model"
)

func WithdrawPublisher(c *gin.Context) {
	type WithdrawRequest struct {
		PublisherID int `json:"publisher_id"`
		Amount      int `json:"amount"`
	}

	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.Lock()
	defer model.Unlock()

	publisher, exists := model.Publishers[req.PublisherID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}

	if publisher.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	publisher.Balance -= req.Amount
	model.Publishers[req.PublisherID] = publisher

	c.JSON(http.StatusOK, publisher)
}

func ViewPublisherStats(c *gin.Context) {
	publisherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publisher ID"})
		return
	}

	model.Lock()
	defer model.Unlock()
	publisher, exists := model.Publishers[publisherID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}

	stats := gin.H{
		"income":      publisher.Income,
		"impressions": publisher.Impressions,
		"clicks":      publisher.Clicks,
	}

	c.JSON(http.StatusOK, stats)
}
