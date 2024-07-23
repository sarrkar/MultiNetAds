package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	mu.Lock()
	defer mu.Unlock()

	publisher, exists := Publishers[req.PublisherID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found"})
		return
	}

	if publisher.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	publisher.Balance -= req.Amount
	Publishers[req.PublisherID] = publisher

	c.JSON(http.StatusOK, publisher)
}