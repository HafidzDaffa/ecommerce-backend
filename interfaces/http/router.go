package http

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app         *fiber.App
	userHandler *UserHandler
}

func NewRouter(app *fiber.App, userHandler *UserHandler) *Router {
	return &Router{
		app:         app,
		userHandler: userHandler,
	}
}

func (r *Router) SetupRoutes() {
	api := r.app.Group("/api/v1")

	api.Get("/health", r.HealthCheck)

	users := api.Group("/users")
	users.Post("/register", r.userHandler.Register)
	users.Post("/login", r.userHandler.Login)
	users.Get("/:id", r.userHandler.GetUser)
	users.Put("/:id", r.userHandler.UpdateUser)
	users.Delete("/:id", r.userHandler.DeleteUser)
	users.Get("/", r.userHandler.ListUsers)
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "API is healthy"
// @Router /health [get]
func (r *Router) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
