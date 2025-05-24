package shareds

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckController interface {
	Status(ctx *gin.Context)
}
type healthcheckController struct {
}

func NewHealthcheckController() *healthcheckController {
	return &healthcheckController{}
}

// Status sends status 200 if the service is up
func (c *healthcheckController) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "version": "1.0"})
}
