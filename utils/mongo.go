package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var VideoCollection *mongo.Collection

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	VideoCollection = client.Database("youtubeFetcher").Collection("videos")

	// Ensure the client is disconnected when the application exits
	go func() {
		<-ctx.Done()
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()
}
