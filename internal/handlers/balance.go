package handlers

import (
	"net/http"
	"strings"

	"github.com/bsv-blockchain/go-sdk/script"
	"github.com/gin-gonic/gin"
	mnee "github.com/mnee-xyz/go-mnee-1sat-sdk"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// GetBalance godoc
// @Summary      Get balance for a single address
// @Description  Retrieves the MNEE balance for a specific wallet address
// @Tags         Balance
// @Produce      json
// @Param        address   path      string  true  "Wallet Address"
// @Success      200       {object}  models.GetBalanceSuccessResponse
// @Failure      400       {object}  models.GenericFailureResponse
// @Failure      500       {object}  models.GenericFailureResponse
// @Router       /balance/{address} [get]
func GetBalance(c *gin.Context) {
	_address := c.Param("address")
	if _address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Address is required"})
		return
	}

	address, err := script.NewAddressFromString(_address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid wallet address: " + _address})
		return
	}

	balances, err := services.Instance.GetBalances(c.Request.Context(), []string{address.AddressString})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if len(balances) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": mnee.BalanceDataDTO{Address: &address.AddressString, Amt: 0.0, Precised: 0.0}})
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
// @Success      200        {object}  models.GetBalancesSuccessResponse
// @Failure      400        {object}  models.GenericFailureResponse
// @Failure      500        {object}  models.GenericFailureResponse
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
			address, err := script.NewAddressFromString(trimmed)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid wallet address: " + trimmed})
				return
			}
			addresses = append(addresses, address.AddressString)
		}
	}

	if len(addresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No valid addresses provided"})
		return
	}

	balances, err := services.Instance.GetBalances(c.Request.Context(), addresses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if len(balances) == 0 {
		for _, addr := range addresses {
			balances = append(balances, mnee.BalanceDataDTO{Address: &addr, Amt: 0.0, Precised: 0.0})
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": balances})
}
