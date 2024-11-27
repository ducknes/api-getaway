package api

import (
	"api-getaway/api/handlers"
	"api-getaway/cluster/storageservice"
	"api-getaway/cluster/userservice"
	"api-getaway/settings"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	s.server.Get("/user", handlers.GetUserHandler(client))
	s.server.Get("/users", handlers.GetUsersHandler(client))
	s.server.Post("/user", handlers.SaveUserHandler(client))
	s.server.Put("/user", handlers.UpdateUserHandler(client))
	s.server.Delete("/user", handlers.DeleteUserHandler(client))
}

func (s *Server) StorageServiceHandlers(client *storageservice.Client) {
	s.server.Get("/product", handlers.GetProductHandler(client))
	s.server.Get("/products", handlers.GetProductsHandler(client))
	s.server.Post("/products", handlers.SaveProductsHandler(client))
	s.server.Put("/products", handlers.UpdateProductsHandler(client))
	s.server.Delete("/products", handlers.DeleteProductsHandler(client))
}

func (s *Server) Start(port int) {
	if err := s.server.Listen(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
