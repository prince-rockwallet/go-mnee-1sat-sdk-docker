package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mnee "github.com/mnee-xyz/go-mnee-1sat-sdk"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

type TransferRequest struct {
	Request []struct {
		Address string  `json:"address" binding:"required"`
		Amount  float64 `json:"amount" binding:"required"`
	} `json:"request" binding:"required"`
	Wif string `json:"wif" binding:"required"`
}

type RawTxRequest struct {
	RawTxHex string `json:"rawTxHex" binding:"required"`
}

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
