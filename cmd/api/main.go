package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	mongoDBURL := os.Getenv("MONGO_DB_URI")

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	router := initializeRouter(client)
	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")

	log.Printf("Server is listening on %s...", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
