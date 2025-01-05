package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/internal/handler/v1/users"
)

// registerUserRoutes registers user routes.
func registerUserRoutesV1(r *gin.RouterGroup, rtr *Router) {
	usersHandler := users.NewHandler(rtr.services.User)

	userRoutes := r.Group("/users")

	userRoutes.GET("/:id", usersHandler.View)
}
