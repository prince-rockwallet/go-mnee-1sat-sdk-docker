package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

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

func GetBalances(c *gin.Context) {
	addrStr := c.Query("addresses")
	if addrStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Addresses query param required"})
		return
	}

	addresses := strings.Split(addrStr, ",")
	balances, err := services.Instance.GetBalances(c.Request.Context(), addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": balances})
}
