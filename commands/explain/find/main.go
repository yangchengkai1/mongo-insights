package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:single@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("noindex")

	result := database.RunCommand(context.Background(), bson.M{
		"explain": bson.M{
			"find":   "test",
			"filter": bson.M{"_id": bson.M{"$eq": 1}},
		},
	})

	if result.Err() != nil {
		raw, _ := result.DecodeBytes()
		log.Println(raw.String())
		return
	}

	raw, _ := result.DecodeBytes()
	log.Println(raw.String())
}
