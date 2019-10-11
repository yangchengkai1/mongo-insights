package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type record struct {
	ID     int64  `bson:"id"`
	User   string `bson:"user,omitempty"`
	Repo   string `bson:"repo"`
	Status int    `bson:"status"`
}

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

	collection := client.Database("index").Collection("test")

	filter := bson.M{"_id": 1000}
	update := bson.M{"$set": bson.M{"uid": 1000000}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("update uid result :", updateResult)

	filter = bson.M{"uid": 10000}
	update = bson.M{"$set": bson.M{"_id": 1000000}}

	updateResult, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("update _id result :", updateResult)
}
