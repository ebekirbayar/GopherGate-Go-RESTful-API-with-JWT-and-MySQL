package router

import (
	"GopherGate/pkg/handler"
	"GopherGate/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Router handles the routing configuration.
type Router struct {
	Handler         handler.Handler
	WebApiFramework *fiber.App
}

// NewRouter creates a new instance of Router.
func NewRouter(handler handler.Handler) *Router {
	return &Router{
		Handler: handler,
	}
}

// InitRouter initializes the router configuration.
func (r *Router) InitRouter() *Router {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "*",
	}))

	auth := app.Group("/api/auth")
	r.initAuthRoutes(auth)

	users := app.Group("/api/users")
	users.Use(middleware.Auth)
	r.initUserRoutes(users)

	r.WebApiFramework = app
	return r
}

// initAuthRoutes initializes authentication routes.
func (r *Router) initAuthRoutes(auth fiber.Router) {
	auth.Post("/register", r.Handler.Api.CreateUser)
	auth.Post("/login", r.Handler.Api.LoginUser)
}

// initUserRoutes initializes user-related routes.
func (r *Router) initUserRoutes(users fiber.Router) {
	users.Get("/me", r.Handler.Api.GetUser)
	users.Put("/:id", r.Handler.Api.UpdateUser)
	users.Delete("/:id", r.Handler.Api.DeleteUser)
	users.Get("", r.Handler.Api.GetUsers)
}
