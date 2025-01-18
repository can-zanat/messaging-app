package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestService_GetSentMessages(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mongoStore := NewMockStore(mockController)
	service := NewService(mongoStore)

	t.Run("Test GetSentMessages success case", func(t *testing.T) {
		mockMessages := &[]Message{
			{ID: primitive.NewObjectID(), Content: "Message 1"},
			{ID: primitive.NewObjectID(), Content: "Message 2"},
		}

		mongoStore.EXPECT().GetSentMessages().Return(mockMessages, nil)

		resp, err := service.GetSentMessages()

		assert.Nil(t, err)
		assert.Equal(t, mockMessages, resp)
	})
	t.Run("Test GetSentMessages failure case", func(t *testing.T) {
		mongoStore.EXPECT().GetSentMessages().Return(nil, assert.AnError)

		resp, err := service.GetSentMessages()

		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestService_StopSending(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mongoStore := NewMockStore(mockController)
	service := NewService(mongoStore)

	t.Run("Test StopSending success case", func(t *testing.T) {
		service.running = true
		service.stopChan = make(chan struct{})

		err := service.StopSending()

		assert.Nil(t, err)
		assert.False(t, service.running)
		assert.Nil(t, service.stopChan)
	})
	t.Run("Test StopSending failure case", func(t *testing.T) {
		service.running = false

		err := service.StopSending()

		assert.NotNil(t, err)
		assert.Equal(t, "sending is already not running", err.Error())
	})
}

func TestService_StartSending(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mongoStore := NewMockStore(mockController)
	service := NewService(mongoStore)

	t.Run("Test StartSending success case", func(t *testing.T) {
		service.running = false

		err := service.StartSending()

		assert.Nil(t, err)
		assert.True(t, service.running)
		assert.NotNil(t, service.stopChan)
	})
	t.Run("Test StartSending failure case", func(t *testing.T) {
		service.running = true

		err := service.StartSending()

		assert.NotNil(t, err)
		assert.Equal(t, "sending is already running", err.Error())
	})
}
