package handlers

import (
	"net/http"

	"github.com/bsv-blockchain/go-sdk/script"
	"github.com/gin-gonic/gin"
	mnee "github.com/mnee-xyz/go-mnee-1sat-sdk"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/models" // Import models
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

type TransferRequest struct {
	Request []struct {
		Address string  `json:"address" binding:"required" example:"1G6CB3Ch4zFkPmuhZzEyChQmrQPfi86qk3"`
		Amount  float64 `json:"amount" binding:"required" example:"0.1"`
	} `json:"request" binding:"required"`
	Wifs []string `json:"wifs" binding:"required" example:"L1dRKo...,K2..."`
}

type RawTxRequest struct {
	RawTxHex string `json:"rawTxHex" binding:"required" example:"01000000..."`
}

// TransferSync godoc
// @Summary      Synchronous Transfer
// @Description  Executes a transfer using multiple WIFs and waits for cosigner response. Returns final TxID.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body TransferRequest true "Transfer Parameters"
// @Success      200     {object} models.TransferSyncSuccessResponse
// @Failure      422     {object} models.GenericFailureResponse
// @Failure      400     {object} models.GenericFailureResponse
// @Failure      500     {object} models.GenericFailureResponse
// @Router       /transaction/transfer [post]
func TransferSync(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.GenericFailureResponse{Success: false, Message: "Unprocessable Entity: " + err.Error()})
		return
	}

	var dtos []mnee.TransferMneeDTO
	for _, r := range req.Request {
		address, err := script.NewAddressFromString(r.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "Invalid wallet address in request: " + r.Address})
			return
		}

		if r.Amount <= 0 {
			c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "Amount must be greater than 0"})
			return
		}

		atomicAmount := uint64(r.Amount * 100000)
		dtos = append(dtos, mnee.TransferMneeDTO{
			Address: address.AddressString,
			Amount:  atomicAmount,
		})
	}

	resp, err := services.Instance.SynchronousTransfer(c.Request.Context(), req.Wifs, dtos, false, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// TransferAsync godoc
// @Summary      Asynchronous Transfer
// @Description  Executes a transfer using multiple WIFs and returns a ticket ID immediately for polling.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body TransferRequest true "Transfer Parameters"
// @Success      200     {object} models.TransferAsyncSuccessResponse
// @Failure      422     {object} models.GenericFailureResponse
// @Failure      400     {object} models.GenericFailureResponse
// @Failure      500     {object} models.GenericFailureResponse
// @Router       /transaction/transfer-async [post]
func TransferAsync(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.GenericFailureResponse{Success: false, Message: "Unprocessable Entity: " + err.Error()})
		return
	}

	var dtos []mnee.TransferMneeDTO
	for _, r := range req.Request {
		address, err := script.NewAddressFromString(r.Address)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "Invalid wallet address in request: " + r.Address})
			return
		}

		if r.Amount <= 0 {
			c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "Amount must be greater than 0"})
			return
		}

		atomicAmount := uint64(r.Amount * 100000)
		dtos = append(dtos, mnee.TransferMneeDTO{
			Address: address.AddressString,
			Amount:  atomicAmount,
		})
	}

	ticketID, err := services.Instance.AsynchronousTransfer(c.Request.Context(), req.Wifs, dtos, false, nil, nil, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ticketId": *ticketID,
		},
	})
}

// SubmitRawTxSync godoc
// @Summary      Submit Raw Transaction (Synchronous)
// @Description  Submits a pre-signed raw transaction hex and waits for the cosigner. Returns the final TxID.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body RawTxRequest true "Raw Hex"
// @Success      200     {object} models.TransferSyncSuccessResponse
// @Failure      422     {object} models.GenericFailureResponse
// @Failure      400     {object} models.GenericFailureResponse
// @Failure      500     {object} models.GenericFailureResponse
// @Router       /transaction/submit-rawtx [post]
func SubmitRawTxSync(c *gin.Context) {
	var req RawTxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.GenericFailureResponse{Success: false, Message: "Unprocessable Entity: " + err.Error()})
		return
	}

	if req.RawTxHex == "" {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "rawTxHex cannot be empty"})
		return
	}

	resp, err := services.Instance.SubmitRawTxSync(c.Request.Context(), req.RawTxHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

// SubmitRawTxAsync godoc
// @Summary      Submit Raw Transaction (Asynchronous)
// @Description  Submits a pre-signed raw transaction hex and returns a ticket ID immediately.
// @Tags         Transfer
// @Accept       json
// @Produce      json
// @Param        request body RawTxRequest true "Raw Hex"
// @Success      200     {object} models.TransferAsyncSuccessResponse
// @Failure      422     {object} models.GenericFailureResponse
// @Failure      400     {object} models.GenericFailureResponse
// @Failure      500     {object} models.GenericFailureResponse
// @Router       /transaction/submit-rawtx-async [post]
func SubmitRawTxAsync(c *gin.Context) {
	var req RawTxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.GenericFailureResponse{Success: false, Message: "Unprocessable Entity: " + err.Error()})
		return
	}

	if req.RawTxHex == "" {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "rawTxHex cannot be empty"})
		return
	}

	ticketID, err := services.Instance.SubmitRawTxAsync(c.Request.Context(), req.RawTxHex, nil, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ticketId": *ticketID,
		},
	})
}
