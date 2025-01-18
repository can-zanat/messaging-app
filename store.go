package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBStore interface {
	GetTwoMessages() error
	LogSentMessages() error
	GetSentMessages() error
}

type MongoDBStore struct {
	Client *mongo.Client
}

func NewStore() *MongoDBStore {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
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

func (m *MongoDBStore) GetTwoMessages() error {
	return nil
}

func (m *MongoDBStore) LogSentMessages() error {
	return nil
}

func (m *MongoDBStore) GetSentMessages() error {
	return nil
}
