package handlers

import (
	"fmt"
	"karhub/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type beerHandler struct {
	beerUsecase BeerUsecase
}

func (h *beerHandler) Post(ctx *fiber.Ctx) error {
	var beer entity.BeerStyle

	err := ctx.BodyParser(&beer)
	if err != nil {
		log.Error(fmt.Sprintf("error to handle post unmarshal: %s", err))
		return fmt.Errorf("sorry, something went wrong")
	}

	beerStyle, err := h.beerUsecase.Create(ctx.UserContext(), beer)
	if err != nil {
		log.Error(fmt.Sprintf("error to post create: %s", err))
		return fmt.Errorf("sorry, something went wrong")
	}

	return ctx.Status(200).JSON(beerStyle)
}

func (h *beerHandler) Get(ctx *fiber.Ctx) error {
	var temperature entity.BeerTemperature

	temperature.Temperature = ctx.QueryInt("q")

	beerTemperature, err := h.beerUsecase.Read(ctx.UserContext(), temperature)
	if err != nil {
		log.Error(fmt.Sprintf("error to get read : %s", err))
		return fmt.Errorf("sorry, something went wrong")
	}

	return ctx.Status(200).JSON(beerTemperature)
}

func (h *beerHandler) Update(ctx *fiber.Ctx) error {
	var beer entity.BeerStyle

	err := ctx.BodyParser(&beer)
	if err != nil {
		log.Error(fmt.Sprintf("error to handle update unmarshal: %s", err))
		return fmt.Errorf("sorry, something went wrong")
	}

	beerStyle, err := h.beerUsecase.Update(ctx.UserContext(), beer)
	if err != nil {
		log.Error(fmt.Sprintf("error to put update: %s", err))
		return fmt.Errorf("sorry, something went wrong")
	}

	return ctx.Status(200).JSON(beerStyle)
}

func (h *beerHandler) Delete(ctx *fiber.Ctx) error {
	beerStyle := ctx.Params("delete")

	err := h.beerUsecase.Delete(ctx.UserContext(), beerStyle)
	if err != nil {
		log.Error(fmt.Sprintf("error to delete: %s", err))
		return fmt.Errorf("sorry, something went wrong")
	}

	ctx.Status(200)
	return nil
}

func NewBeersHandler(beerUsecase BeerUsecase) *beerHandler {
	return &beerHandler{beerUsecase: beerUsecase}
}
