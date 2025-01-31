package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatMessage struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    UserID    int             `bson:"userId"`
    Message   string             `bson:"message"`
    Timestamp time.Time          `bson:"timestamp"`
}