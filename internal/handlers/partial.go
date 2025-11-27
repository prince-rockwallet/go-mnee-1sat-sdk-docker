package handlers

import (
	"net/http"

	"github.com/bsv-blockchain/go-sdk/script"
	"github.com/gin-gonic/gin"
	mnee "github.com/mnee-xyz/go-mnee-1sat-sdk"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// PartialSign godoc
// @Summary      Partial Sign Transaction
// @Description  Builds and signs a transaction *only* with the provided WIFs. Returns hex.
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param        request body TransferRequest true "Transfer Parameters"
// @Success      200     {object} map[string]interface{}
// @Router       /transaction/partial-sign [post]
func PartialSign(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Unprocessable Entity: " + err.Error()})
		return
	}

	var dtos []mnee.TransferMneeDTO
	for _, r := range req.Request {
		address, err := script.NewAddressFromString(r.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid wallet address in request: " + r.Address})
			return
		}

		if r.Amount <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Amount must be greater than 0"})
			return
		}

		atomicAmount := uint64(r.Amount * 100000)
		dtos = append(dtos, mnee.TransferMneeDTO{
			Address: address.AddressString,
			Amount:  atomicAmount,
		})
	}

	hex, err := services.Instance.PartialSign(c.Request.Context(), req.Wifs, dtos, false, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"rawTxHex": *hex,
		},
	})
}
