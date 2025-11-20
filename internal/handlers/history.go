package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

func GetHistory(c *gin.Context) {
	address := c.Param("address")
	fromScore, _ := strconv.Atoi(c.DefaultQuery("fromScore", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Address is required"})
		return
	}

	history, err := services.Instance.GetSpecificTransactionHistory(c.Request.Context(), []string{address}, fromScore, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"address": address,
			"history": history,
		},
	})
}
