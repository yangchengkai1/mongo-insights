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
	ID     int64  `bson:"_id"`
	User   string `bson:"user"`
	Repo   string `bson:"repo"`
	Status int    `bson:"status"`
	UID    int64  `bson:"uid"`
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

	clone, err := collection.Clone()
	if err != nil {
		log.Fatal(err)
	}

	count, err := clone.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("clone", count, clone.Name())

	r := record{
		User:   "clone",
		Repo:   "clone",
		ID:     600001,
		Status: 1,
		UID:    600001,
	}

	if _, err := clone.InsertOne(context.Background(), &r); err != nil {
		log.Fatal(err)
	}

	count, err = clone.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("clone", count, clone.Name(), clone)

	count, err = collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("index", count, collection.Name(), collection)
}
