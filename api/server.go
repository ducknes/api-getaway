package api

import (
	"api-getaway/api/handlers"
	"api-getaway/cluster/authservice"
	"api-getaway/cluster/storageservice"
	"api-getaway/cluster/userservice"
	"api-getaway/settings"
	"api-getaway/tools/middlewares"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/fiber-swagger"

	_ "api-getaway/docs"
)

type Server struct {
	server *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{AppName: settings.AppName()})
	app.Use(cors.New(), healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/health",
		ReadinessEndpoint: "/ready",
	}), recover.New())

	return &Server{server: app}
}

func (s *Server) UserServiceHandlers(client *userservice.Client) {
	userGroup := s.server.Group("/user", middlewares.AuthMiddleware)
	userGroup.Get("/", handlers.GetUserHandler(client))
	userGroup.Get("/many", handlers.GetUsersHandler(client))
	userGroup.Post("/", handlers.SaveUserHandler(client))
	userGroup.Put("/", handlers.UpdateUserHandler(client))
	userGroup.Delete("/", handlers.DeleteUserHandler(client))
}

func (s *Server) StorageServiceHandlers(client *storageservice.Client) {
	storageGroup := s.server.Group("/products", middlewares.AuthMiddleware)
	storageGroup.Get("/one", handlers.GetProductHandler(client))
	storageGroup.Get("/", handlers.GetProductsHandler(client))
	storageGroup.Post("/", handlers.SaveProductsHandler(client))
	storageGroup.Put("/", handlers.UpdateProductsHandler(client))
	storageGroup.Delete("/", handlers.DeleteProductsHandler(client))
}

func (s *Server) AuthServiceHandlers(client *authservice.Client) {
	s.server.Post("/login", handlers.LoginHandler(client))
	s.server.Post("/logout", handlers.LogoutHandler(client))
	s.server.Post("/refresh", handlers.RefreshHandler(client))
	s.server.Post("/sign-up", handlers.SignUpHandler(client))
}

func (s *Server) Swagger() {
	s.server.Get("/swagger/*", fiberSwagger.WrapHandler)
}

func (s *Server) Start(port int) {
	if err := s.server.Listen(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
