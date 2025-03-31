package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error(
			"Failed to load .env file",
			slog.String("msg", err.Error()),
		)
		os.Exit(1)
	}

	rabbitConnection := os.Getenv("RABBITMQ_URL")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	conn, err := amqp.Dial(rabbitConnection)
	if err != nil {
		slog.Error(
			"RabbitMQ connection error",
			slog.String("msg", err.Error()),
		)
		os.Exit(1)
	}
	defer conn.Close()

	slog.Info("Connection was established successfully!")

	<-stop

	slog.Info("Shutting down server...")
}
