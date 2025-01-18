package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_StartSending(t *testing.T) {
	app := createTestApp()
	mockService, mockServiceController := createMockService(t)
	handler := NewHandler(mockService)

	t.Run("Should return 500 if service returns an error", func(t *testing.T) {
		mockService.EXPECT().StartSending().Return(assert.AnError)

		handler.RegisterRoutes(app)

		req := httptest.NewRequest("POST", "/start-sending", http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		mockServiceController.Finish()

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
	t.Run("Should return 200 if everything is ok", func(t *testing.T) {
		mockService.EXPECT().StartSending().Return(nil)

		handler.RegisterRoutes(app)

		req := httptest.NewRequest("POST", "/start-sending", http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		mockServiceController.Finish()

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestHandler_StopSending(t *testing.T) {
	app := createTestApp()
	mockService, mockServiceController := createMockService(t)
	handler := NewHandler(mockService)

	t.Run("Should return 500 if service returns an error", func(t *testing.T) {
		mockService.EXPECT().StopSending().Return(assert.AnError)

		handler.RegisterRoutes(app)

		req := httptest.NewRequest("POST", "/stop-sending", http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		mockServiceController.Finish()

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
	t.Run("Should return 200 if everything is ok", func(t *testing.T) {
		mockService.EXPECT().StopSending().Return(nil)

		handler.RegisterRoutes(app)

		req := httptest.NewRequest("POST", "/stop-sending", http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		mockServiceController.Finish()

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestHandler_GetSentMessages(t *testing.T) {
	app := createTestApp()
	mockService, mockServiceController := createMockService(t)
	handler := NewHandler(mockService)

	t.Run("Should return 500 if service returns an error", func(t *testing.T) {
		mockService.EXPECT().GetSentMessages().Return(nil, assert.AnError)

		handler.RegisterRoutes(app)

		req := httptest.NewRequest("GET", "/sent-messages", http.NoBody)

		resp, _ := app.Test(req)

		mockServiceController.Finish()

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
	t.Run("Should return 200 if everything is ok", func(t *testing.T) {
		mockService.EXPECT().GetSentMessages().Return(&[]Message{}, nil)

		handler.RegisterRoutes(app)

		req := httptest.NewRequest("GET", "/sent-messages", http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		mockServiceController.Finish()

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func createMockService(t *testing.T) (*Mockactions, *gomock.Controller) {
	t.Helper()

	mockServiceController := gomock.NewController(t)
	mockService := NewMockactions(mockServiceController)

	return mockService, mockServiceController
}

func createTestApp() *fiber.App {
	return fiber.New(fiber.Config{})
}
