package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	handler    Handler
	serverPort string
}

func (s *server) handlers(app *fiber.App) {
	app.Post("/beer-style", s.handler.Post)
	app.Get("/temperature", s.handler.Get)
	app.Put("/beer-style", s.handler.Update)
	app.Delete("/beer-style/:delete", s.handler.Delete)
}

func (s *server) Start() {
	app := fiber.New()

	s.handlers(app)
	log.Fatal(app.Listen(":" + s.serverPort))
}

func NewServer(handler Handler, serverPort string) *server {
	return &server{handler: handler, serverPort: serverPort}
}
