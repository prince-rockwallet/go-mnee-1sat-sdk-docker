package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

func GetAllUtxos(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Address is required"})
		return
	}

	txos, err := services.Instance.GetUnspentTxos(c.Request.Context(), []string{address})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": txos})
}

func GetPaginatedUtxos(c *gin.Context) {
	address := c.Param("address")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Address is required"})
		return
	}

	txos, err := services.Instance.GetPaginatedUnspentTxos(c.Request.Context(), []string{address}, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": txos})
}
