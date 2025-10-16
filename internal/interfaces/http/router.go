package http

import (
	"net/http"
	"vibe-storm/internal/application/dto"
	"vibe-storm/internal/interfaces/http/handlers"
	"vibe-storm/pkg/config"
	"vibe-storm/pkg/database"
	"vibe-storm/pkg/middleware"

	"github.com/labstack/echo/v4"
)

// SetupRoutes configures all the application routes
func SetupRoutes(e *echo.Echo) {
	// Initialize dependencies
	cfg := config.Load()
	db := database.GetDB()

	deps := handlers.HandlerDeps{
		DB:     db,
		Config: cfg,
	}

	// API v1 routes
	v1 := e.Group("/api/v1")

	// Health check
	v1.GET("/health", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.HealthCheckHandler{Deps: deps}
		return handler.Handle(c)
	}))

	// Authentication routes (public)
	v1.POST("/auth/signup", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.SignupHandler{Deps: deps}
		return handler.Handle(c)
	}))

	v1.POST("/auth/signin", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.SigninHandler{Deps: deps}
		return handler.Handle(c)
	}))

	v1.POST("/auth/refresh", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.RefreshTokenHandler{Deps: deps}
		return handler.Handle(c)
	}))

	// Protected authentication routes
	protected := v1.Group("/auth")
	protected.Use(middleware.JWTMiddleware(cfg.JWT.Secret))
	protected.GET("/me", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.MeHandler{Deps: deps}
		return handler.Handle(c)
	}))

	// User management routes (protected)
	v1.GET("/users", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.GetUsersHandler{Deps: deps}
		return handler.Handle(c)
	}))

	v1.POST("/users", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.CreateUserHandler{Deps: deps}
		return handler.Handle(c)
	}))

	v1.GET("/users/:id", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.GetUserHandler{Deps: deps}
		return handler.Handle(c)
	}))

	v1.PUT("/users/:id", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.UpdateUserHandler{Deps: deps}
		return handler.Handle(c)
	}))

	v1.DELETE("/users/:id", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.DeleteUserHandler{Deps: deps}
		return handler.Handle(c)
	}))

	// Web routes for templates
	e.GET("/", WrapHandler(func(c echo.Context) (interface{}, error) {
		handler := &handlers.HomePageHandler{Deps: deps}
		return handler.Handle(c)
	}))

	// Static files
	e.Static("/static", "web/static")

	// 404 handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusNotFound

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				c.NoContent(code)
			} else {
				errorResponse := dto.NewStructuredErrorResponse(
					dto.ErrCodeNotFound,
					c.Request().URL.Path,
				)
				c.JSON(code, errorResponse)
			}
		}
	}
}
