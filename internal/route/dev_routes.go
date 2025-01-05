package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/internal/handler"
)

// registerDevRoutes registers development routes.
func registerDevRoutes(r *gin.RouterGroup, rtr *Router) {

	r.GET("/ping", handler.PingHandler)
}
