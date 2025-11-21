package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// GetHistory godoc
// @Summary      Get transaction history for multiple addresses
// @Description  Retrieves transaction history for one or more addresses with pagination
// @Tags         History
// @Produce      json
// @Param        addresses query     string  true  "Comma-separated list of Wallet Addresses"
// @Param        fromScore query     int     false "Starting score (default 0)"
// @Param        limit     query     int     false "Limit (default 10)"
// @Success      200       {object}  map[string]interface{}
// @Router       /transaction [get]
func GetHistory(c *gin.Context) {
	addrStr := c.Query("addresses")
	if addrStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "addresses query parameter is required"})
		return
	}

	rawAddresses := strings.Split(addrStr, ",")
	var addresses []string
	for _, addr := range rawAddresses {
		trimmed := strings.TrimSpace(addr)
		if trimmed != "" {
			addresses = append(addresses, trimmed)
		}
	}

	if len(addresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No valid addresses provided"})
		return
	}

	fromScore, _ := strconv.Atoi(c.DefaultQuery("fromScore", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	history, err := services.Instance.GetSpecificTransactionHistory(c.Request.Context(), addresses, fromScore, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"history": history,
		},
	})
}
