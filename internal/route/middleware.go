package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/internal/middleware"
)

// registerMiddleware registers middleware.
func registerMiddleware(e *gin.Engine) {
	e.Use(middleware.RecoveryOnPanic())
}
