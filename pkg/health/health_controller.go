package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Ready(c *gin.Context) {
	c.String(http.StatusOK, "ready")
}

func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "work")
}
