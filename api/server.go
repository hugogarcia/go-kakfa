package api

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugogarcia/go-kakfa/kafka"
)

const (
	ProducerPort = ":8080"
)

func Execute() error {
	producer, err := kafka.SetupProducer(context.Background(), kafka.UsersTopic)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/users", GetAllUsersHandler())
	router.POST("/users", SendUsersHandler(producer))

	fmt.Printf("Kafka PRODUCER 📨 started at http://localhost%s\n",
		ProducerPort)

	if err := router.Run(ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}

	return nil
}
