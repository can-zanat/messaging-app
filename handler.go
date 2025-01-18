package main

import (
	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	service actions
}

type actions interface {
	StartSending() error
	StopSending() error
	GetSentMessages() error
}

func NewHandler(service actions) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/start-sending", h.StartSending)
	app.Get("/stop-sending", h.StopSending)
	app.Post("/sent-messages", h.GetSentMessages)
}

func (h *Handler) StartSending(_ *fiber.Ctx) error {
	return nil
}

func (h *Handler) StopSending(_ *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetSentMessages(_ *fiber.Ctx) error {
	return nil
}
