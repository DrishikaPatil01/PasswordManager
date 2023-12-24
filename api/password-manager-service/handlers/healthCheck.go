package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Healthy Upstream")
}
