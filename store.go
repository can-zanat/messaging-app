package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MessageQuantity = 2
const ContentMaxLength = 250

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

func (m *MongoDBStore) GetTwoMessages() (*[]Message, error) {
	collection := m.Client.Database("messages").Collection("messages")

	filter := bson.M{"is_sent": false}

	// Sort by _id field in ascending order and get only 2 records
	opts := options.Find().
		SetSort(bson.D{{Key: "_id", Value: 1}}).
		SetLimit(int64(MessageQuantity))

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []Message
	if err := cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	var filteredMessages []Message

	for _, msg := range messages {
		if len(msg.Content) > ContentMaxLength {
			log.Printf("Message ID %s exceeds content length limit with %d characters\n", msg.ID.Hex(), len(msg.Content))
		} else {
			filteredMessages = append(filteredMessages, msg)
		}
	}

	return &filteredMessages, nil
}

func (m *MongoDBStore) UpdateSentStatus(msg *[]Message) error {
	collection := m.Client.Database("messages").Collection("messages")

	ids := make([]primitive.ObjectID, 0, MessageQuantity)
	for _, message := range *msg {
		ids = append(ids, message.ID)
	}

	filter := bson.M{
		"_id": bson.M{"$in": ids},
	}

	update := bson.M{
		"$set": bson.M{"is_sent": true, "sent_time": time.Now().Format(time.RFC3339)},
	}

	_, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update is_sent: %w", err)
	}

	return nil
}

func (m *MongoDBStore) GetSentMessages() (*[]Message, error) {
	filter := bson.M{"is_sent": true}

	collection := m.Client.Database("messages").Collection("messages")

	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var messages []Message
	if err := cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return &messages, nil
}
