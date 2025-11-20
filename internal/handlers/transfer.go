package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mnee "github.com/mnee-xyz/go-mnee-1sat-sdk"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

type TransferRequest struct {
	Request []struct {
		Address string  `json:"address" binding:"required" example:"1G6CB3Ch4zFkPmuhZzEyChQmrQPfi86qk3"`
		Amount  float64 `json:"amount" binding:"required" example:"0.1"`
	} `json:"request" binding:"required"`
	Wif string `json:"wif" binding:"required" example:"L1..."`
}

type RawTxRequest struct {
	RawTxHex string `json:"rawTxHex" binding:"required" example:"01000000..."`
}

// TransferSync godoc
// @Summary      Synchronous Transfer
// @Description  Executes a transfer and waits for cosigner response. Returns final TxID.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body TransferRequest true "Transfer Parameters"
// @Success      200     {object} map[string]interface{}
// @Router       /transaction/transfer [post]
func TransferSync(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	var dtos []mnee.TransferMneeDTO
	for _, r := range req.Request {
		atomicAmount := uint64(r.Amount * 100000)
		dtos = append(dtos, mnee.TransferMneeDTO{
			Address: r.Address,
			Amount:  atomicAmount,
		})
	}

	resp, err := services.Instance.SynchronousTransfer(c.Request.Context(), []string{req.Wif}, dtos, false, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// TransferAsync godoc
// @Summary      Asynchronous Transfer
// @Description  Executes a transfer and returns a ticket ID immediately for polling.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body TransferRequest true "Transfer Parameters"
// @Success      200     {object} map[string]interface{}
// @Router       /transaction/transfer-async [post]
func TransferAsync(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	var dtos []mnee.TransferMneeDTO
	for _, r := range req.Request {
		atomicAmount := uint64(r.Amount * 100000)
		dtos = append(dtos, mnee.TransferMneeDTO{
			Address: r.Address,
			Amount:  atomicAmount,
		})
	}

	ticketID, err := services.Instance.AsynchronousTransfer(c.Request.Context(), []string{req.Wif}, dtos, false, nil, nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ticketId": *ticketID,
		},
	})
}

// SubmitRawTxAsync godoc
// @Summary      Submit Raw Transaction
// @Description  Submits a pre-signed raw transaction string.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body RawTxRequest true "Raw Hex"
// @Success      200     {object} map[string]interface{}
// @Router       /transaction/submit-rawtx [post]
func SubmitRawTxAsync(c *gin.Context) {
	var req RawTxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	ticketID, err := services.Instance.SubmitRawTxAsync(c.Request.Context(), req.RawTxHex, nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ticketId": *ticketID,
		},
	})
}
