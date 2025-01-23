package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("youtubeFetcher")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"publish_time": -1}, // Index for sorting by publish time
		Options: options.Index().SetUnique(false).SetName("publish_time_desc"),
	}
	_, err = DB.Collection("videos").Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal("Could not create index:", err)
	}
}
