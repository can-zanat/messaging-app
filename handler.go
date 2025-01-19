package main

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	service actions
}

type actions interface {
	StartSending() error
	StopSending() error
	GetSentMessages() (*[]Message, error)
}

var disableAutoStart = false

func NewHandler(service actions) *Handler {
	h := &Handler{service: service}

	if !disableAutoStart {
		go func() {
			err := service.StartSending()
			if err != nil {
				fmt.Println("Failed to start sending process:", err)
			}
		}()
		fmt.Println("Sending process started.")
	}

	return h
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/start-sending", h.StartSending)
	app.Post("/stop-sending", h.StopSending)
	app.Get("/sent-messages", h.GetSentMessages)
}

func (h *Handler) StartSending(ctx *fiber.Ctx) error {
	err := h.service.StartSending()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{"message": "send process started"})
}

func (h *Handler) StopSending(ctx *fiber.Ctx) error {
	err := h.service.StopSending()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{"message": "start process ended"})
}

func (h *Handler) GetSentMessages(ctx *fiber.Ctx) error {
	messages, err := h.service.GetSentMessages()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(messages)
}
