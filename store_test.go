package main

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoImage = "mongo:7.0.4"
)

func NewStoreWithURI(uri string) *MongoDBStore {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("Connection failure:", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Unable to access MongoDB:", err)
	}

	return &MongoDBStore{
		Client: client,
	}
}

func prepareTestStore(t *testing.T) (store *MongoDBStore, clean func()) {
	t.Helper()

	ctx := context.Background()

	mongodbContainer, err := mongodb.RunContainer(ctx, testcontainers.WithImage(mongoImage))
	if err != nil {
		t.Fatalf("Failed to start MongoDB container: %v", err)
	}

	clean = func() {
		if terminateErr := mongodbContainer.Terminate(ctx); terminateErr != nil {
			t.Fatalf("Failed to terminate MongoDB container: %v", terminateErr)
		}
	}

	containerURI, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("Failed to get container connection string: %v", err)
	}

	s := NewStoreWithURI(containerURI)

	return s, clean
}

func TestMongoDBStore_GetTwoMessages(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	store, clean := prepareTestStore(t)
	defer clean()

	collection := store.Client.Database("messages").Collection("messages")

	_, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "content", Value: "Hello, World!"},
		{Key: "is_sent", Value: false},
	})
	assert.Nil(t, err)

	_, err = collection.InsertOne(context.Background(), bson.D{
		{Key: "content", Value: "Hola Mundo!"},
		{Key: "is_sent", Value: false},
	})
	assert.Nil(t, err)

	tests := []struct {
		Name           string
		ExpectedResult []string
		ExpectedError  error
	}{
		{
			ExpectedResult: []string{"Hello, World!", "Hola Mundo!"},
			ExpectedError:  nil,
			Name:           "should return two messages correctly",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := store.GetTwoMessages()

			assert.Equal(t, tc.ExpectedError, err)

			var actualContents []string
			for _, msg := range *result {
				actualContents = append(actualContents, msg.Content)
			}

			assert.Equal(t, tc.ExpectedResult, actualContents)
		})
	}
}

func TestMongoDBStore_UpdateSentStatus(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	store, clean := prepareTestStore(t)
	defer clean()

	collection := store.Client.Database("messages").Collection("messages")

	_, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "content", Value: "Hello, World!"},
		{Key: "is_sent", Value: false},
	})
	assert.Nil(t, err)

	_, err = collection.InsertOne(context.Background(), bson.D{
		{Key: "content", Value: "Hola Mundo!"},
		{Key: "is_sent", Value: false},
	})
	assert.Nil(t, err)

	tests := []struct {
		Name          string
		ExpectedError error
	}{
		{
			ExpectedError: nil,
			Name:          "should update two messages correctly",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			messages, err := store.GetTwoMessages()
			assert.Nil(t, err)

			err = store.UpdateSentStatus(messages)
			assert.Equal(t, tc.ExpectedError, err)

			messages, err = store.GetTwoMessages()
			assert.Nil(t, err)

			for _, msg := range *messages {
				assert.True(t, msg.IsSent)
			}
		})
	}
}

func TestMongoDBStore_GetSentMessages(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	store, clean := prepareTestStore(t)
	defer clean()

	collection := store.Client.Database("messages").Collection("messages")

	_, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "content", Value: "Hello, World!"},
		{Key: "is_sent", Value: true},
	})
	assert.Nil(t, err)

	_, err = collection.InsertOne(context.Background(), bson.D{
		{Key: "content", Value: "Hola Mundo!"},
		{Key: "is_sent", Value: true},
	})
	assert.Nil(t, err)

	tests := []struct {
		Name           string
		ExpectedResult []string
		ExpectedError  error
	}{
		{
			ExpectedResult: []string{"Hello, World!", "Hola Mundo!"},
			ExpectedError:  nil,
			Name:           "should return two messages correctly",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := store.GetSentMessages()

			assert.Equal(t, tc.ExpectedError, err)

			var actualContents []string
			for _, msg := range *result {
				actualContents = append(actualContents, msg.Content)
			}

			assert.Equal(t, tc.ExpectedResult, actualContents)
		})
	}
}
