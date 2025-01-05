package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/config"
	"github.com/go-nop/ginstarter/internal/repositories"
	"github.com/go-nop/ginstarter/internal/services"
)

// Router is a struct that holds the router.
type Router struct {
	appConfig    *config.App
	repositories *repositories.Repositories
	services     *services.Services
}

// NewRouter creates a new Router.
func NewRouter(opts ...Option) *Router {
	r := &Router{}
	for _, opt := range opts {
		opt(r)
	}

	return r
}

// Init initializes the router.
func (rtr *Router) Init() *gin.Engine {
	if rtr.appConfig.DebugMode {
		gin.SetMode(gin.DebugMode)
	}

	e := gin.Default()

	// Register middleware.
	registerMiddleware(e)

	// create a new group for api versions.
	api := e.Group("/api")
	apiV1 := api.Group("/v1")

	// Register development routes for all api versions.
	registerDevRoutes(api, rtr)

	// Register user routes.
	registerUserRoutesV1(apiV1, rtr)

	return e
}
