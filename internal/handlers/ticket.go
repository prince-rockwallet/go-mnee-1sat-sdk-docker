package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/models" // Import models
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// PollTicket godoc
// @Summary      Poll Ticket Status
// @Description  Polls the status of a transaction ticket until it is processed or times out.
// @Tags         Transaction
// @Produce      json
// @Param        ticketId  path      string  true  "Ticket ID"
// @Success      200       {object}  models.GetTicketSuccessResponse
// @Failure      400       {object}  models.GenericFailureResponse
// @Failure      500       {object}  models.GenericFailureResponse
// @Router       /transaction/status/{ticketId} [get]
func PollTicket(c *gin.Context) {
	ticketID := c.Param("ticketId")
	if ticketID == "" {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: "ticketId is required"})
		return
	}

	ticket, err := services.Instance.PollTicket(c.Request.Context(), ticketID, 2*time.Second)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericFailureResponse{Success: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ticket,
	})
}
