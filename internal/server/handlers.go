package server

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Post(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
