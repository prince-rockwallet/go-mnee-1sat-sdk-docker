package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// GetBalance godoc
// @Summary      Get balance for a single address
// @Description  Retrieves the MNEE balance for a specific wallet address
// @Tags         Balance
// @Produce      json
// @Param        address   path      string  true  "Wallet Address"
// @Success      200       {object}  map[string]interface{}
// @Router       /balance/{address} [get]
func GetBalance(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Address is required"})
		return
	}

	balances, err := services.Instance.GetBalances(c.Request.Context(), []string{address})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	if len(balances) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Address not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": balances[0]})
}

// GetBalances godoc
// @Summary      Get balances for multiple addresses
// @Description  Retrieves balances for a comma-separated list of addresses
// @Tags         Balance
// @Produce      json
// @Param        addresses  query     string  true  "Comma-separated list of addresses"
// @Success      200        {object}  map[string]interface{}
// @Router       /balance [get]
func GetBalances(c *gin.Context) {
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

	balances, err := services.Instance.GetBalances(c.Request.Context(), addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": balances})
}
