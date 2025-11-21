package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

// GetConfig godoc
// @Summary      Get System Config
// @Description  Returns current MNEE system configuration.
// @Tags         Config
// @Produce      json
// @Success      200       {object}  map[string]interface{}
// @Router       /config [get]
func GetConfig(c *gin.Context) {
	config, err := services.Instance.GetConfig(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    config,
	})
}
