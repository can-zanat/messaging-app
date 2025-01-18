package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Content   string             `json:"content" bson:"content"`
	Recipient string             `json:"recipient" bson:"recipient"`
	SentTime  string             `json:"sent_time" bson:"sent_time"`
	IsSent    bool               `json:"is_sent" bson:"is_sent"`
}
