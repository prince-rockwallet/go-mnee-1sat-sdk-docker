package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bsv-blockchain/go-sdk/script"
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
			address, err := script.NewAddressFromString(trimmed)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid wallet address: " + trimmed})
				return
			}
			addresses = append(addresses, address.AddressString)
			addresses = append(addresses, trimmed)
		}
	}

	if len(addresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No valid addresses provided"})
		return
	}

	var fromScore int
	fromQuery := c.Query("fromScore")
	if fromQuery == "" {
		fromScore = 0
	} else {
		var err error
		fromScore, err = strconv.Atoi(fromQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "fromScore must be a valid integer"})
			return
		}
		if fromScore < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "fromScore must not be negative"})
			return
		}
	}

	var limit int
	limitQuery := c.Query("limit")
	if limitQuery == "" {
		limit = 10
	} else {
		var err error
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "limit must be a valid integer"})
			return
		}
		if limit < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "limit must not be negative"})
			return
		}
	}

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
