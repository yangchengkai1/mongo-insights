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

	collection := client.Database("simple").Collection("article")

	result, err := collection.BulkWrite(context.Background(), bson.M{
		"insertOne":  bson.M{"id": 3020000, "uid": 3020000, "status": 1},
		"insertMany": bson.M{},
		"updates": bson.M{
			"q": bson.M{"_id": bson.M{"$eq": 3020000}},
			"u": bson.M{"uid": bson.M{"$set": 3030000}},
		},
		"deleteOne": bson.M{
			"filter": bson.M{"_id": bson.M{"$eq": 3010000}},
		},
	})

}
