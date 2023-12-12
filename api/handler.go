package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hugogarcia/go-kakfa/models"
	"github.com/segmentio/kafka-go"
)

var usersDB = []models.User{
	{ID: uuid.New().String(), Name: "User 1"},
	{ID: uuid.New().String(), Name: "User 2"},
	{ID: uuid.New().String(), Name: "User 3"},
}

func GetAllUsersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, usersDB)
	}
}

func SendUsersHandler(producer *kafka.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usersJson, err := json.Marshal(usersDB)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		bytes, err := producer.WriteMessages(
			kafka.Message{Value: usersJson},
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		fmt.Printf("Message sent. Bytes: %d", bytes)

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Message sent"})
	}
}
