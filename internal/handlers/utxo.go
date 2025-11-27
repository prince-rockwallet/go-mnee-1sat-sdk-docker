package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bsv-blockchain/go-sdk/script"
	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// GetAllUtxos godoc
// @Summary      Get all UTXOs for multiple addresses
// @Description  Retrieves all unspent transaction outputs for one or more addresses
// @Tags         UTXO
// @Produce      json
// @Param        addresses query     string  true  "Comma-separated list of Wallet Addresses"
// @Success      200       {object}  map[string]interface{}
// @Router       /utxos/all [get]
func GetAllUtxos(c *gin.Context) {
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

	txos, err := services.Instance.GetUnspentTxos(c.Request.Context(), addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": txos})
}

// GetPaginatedUtxos godoc
// @Summary      Get paginated UTXOs for multiple addresses
// @Description  Retrieves paginated unspent transaction outputs for one or more addresses
// @Tags         UTXO
// @Produce      json
// @Param        addresses query     string  true  "Comma-separated list of Wallet Addresses"
// @Param        page      query     int     false "Page number (default 1)"
// @Param        size      query     int     false "Page size (default 10)"
// @Success      200       {object}  map[string]interface{}
// @Router       /utxos/paginated [get]
func GetPaginatedUtxos(c *gin.Context) {
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

	var page int
	pageQuery := c.Query("page")
	if pageQuery == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "page must be a valid integer"})
			return
		}
		if page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "page must be greater than 0"})
			return
		}
	}

	var size int
	sizeQuery := c.Query("size")
	if sizeQuery == "" {
		size = 10
	} else {
		var err error
		size, err = strconv.Atoi(sizeQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "size must be a valid integer"})
			return
		}
		if size < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "size must be greater than 0"})
			return
		}
	}

	txos, err := services.Instance.GetPaginatedUnspentTxos(c.Request.Context(), addresses, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": txos})
}
