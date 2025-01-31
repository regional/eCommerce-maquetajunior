package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionName = "ChatHistory"
)

func CreateChatMessageHandler(rw http.ResponseWriter, r *http.Request) {
    userId, err := ResolveClaims(rw, r, "userid")
    if err != nil {
        sendError(rw, http.StatusUnauthorized)
        return
    }

    var base64Message struct {
        Message string `json:"message"`
    }
    if err := json.NewDecoder(r.Body).Decode(&base64Message); err != nil {
        sendError(rw, http.StatusBadRequest)
        return
    }

    var message models.ChatMessage
    message.Message = base64Message.Message

    // Convertir userId a int
    if userIdFloat, ok := userId.(float64); ok {
        message.UserID = int(userIdFloat)
    } else {
        sendError(rw, http.StatusInternalServerError)
        return
    }
    message.Timestamp = time.Now()

    err = db.InsertDocument(CollectionName, message)
    if err != nil {
        sendError(rw, http.StatusInternalServerError)
        return
    }

    sendData(rw, message, http.StatusCreated)
}

func GetChatMessagesHandler(rw http.ResponseWriter, r *http.Request) {
	userId, err := ResolveClaims(rw, r, "userid")
    if err != nil {
        sendError(rw, http.StatusUnauthorized)
        return
    }

	var messages []models.ChatMessage
	var user int
	if userIdFloat, ok := userId.(float64); ok {
        user = int(userIdFloat)
    } else {
        sendError(rw, http.StatusInternalServerError)
        return
    }
    filter := bson.M{"userId": user}

    err = db.GetDocuments(CollectionName, filter, &messages)
    if err != nil {
        sendError(rw, http.StatusInternalServerError)
        return
    }

    sendData(rw, messages, http.StatusOK)
}
