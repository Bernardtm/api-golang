package server

import (
	"bernardtm/backend/configs"
	"bernardtm/backend/internal/core/shareds"
	"bernardtm/backend/internal/infra/di"
	"bernardtm/backend/internal/infra/middlewares"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// SetupRouter initializes the Gin router and sets up all the routes
func SetupRouter(container *di.Container, config *configs.AppConfig) *gin.Engine {
	gin.SetMode(config.GinMode)
	router := gin.Default()

	applyMiddlewares(router, config)
	configurePublicRoutes(router, container)
	configureProtectedRoutes(router, container)

	return router
}

// applyMiddlewares applies common middlewares to the router
func applyMiddlewares(router *gin.Engine, config *configs.AppConfig) {
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(gin.Recovery())
	router.Use(middlewares.CORS(config.CorsOrigin))
	router.Use(middlewares.SecurityHeadersMiddleware())
}

// configurePublicRoutes sets up public routes
func configurePublicRoutes(router *gin.Engine, c *di.Container) {
	// healthcheck
	router.GET("", c.HealthcheckController.Status)

	api := router.Group("/api/v1")

	// auth
	api.POST("/auth/login", c.AuthController.Login)
	api.POST("/auth/login/request-password-reset", c.AuthController.RequestPasswordReset)

}

// configureProtectedRoutes sets up protected routes
func configureProtectedRoutes(router *gin.Engine, c *di.Container) {
	jwtMiddleware := middlewares.NewJWTMiddleware(c.TokenService)
	jwtQueryMiddleware := middlewares.NewJWTQueryMiddleware(c.TokenService)

	api := router.Group("/api/v1")

	// auth
	api.POST("/auth/login/verify", jwtMiddleware.AuthMiddleware("2step_verification"), c.AuthController.Login2Step)
	api.POST("/auth/login/reset-password", jwtMiddleware.AuthMiddleware("password_reset_verification"), c.AuthController.ResetPassword)

	api.GET("/ws", jwtQueryMiddleware.AuthQueryMiddleware("api"), func(ctx *gin.Context) {
		c.SocketHandler.WebSocketHandler(ctx)
	})

	api.Use(jwtMiddleware.AuthMiddleware("api"))

	// files
	api.POST("/files", c.FilesController.Create)

	// menus
	api.GET("/menus/user", c.MenusController.GetMenusByUserID)

	// CRUD routes
	entities := map[string]shareds.CrudController{
		"menus":  c.MenusController,
		"users":  c.UserController,
		"status": c.StatusController,
		// Add more entities as needed
	}

	for path, controller := range entities {
		setupCrudRoutes(api, path, controller)
	}
}

func setupCrudRoutes(group *gin.RouterGroup, path string, controller shareds.CrudController) {
	routes := group.Group(path)
	{
		routes.GET("", controller.GetAll)
		routes.GET("/:id", controller.GetByID)
		routes.GET("/paginate", controller.Paginate)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
		routes.DELETE("/:id", controller.Delete)
	}
}
