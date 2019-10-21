package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type record struct {
	ID     int64  `bson:"_id"`
	UID    int64  `bson:"uuid"`
	User   string `bson:"user"`
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

	insert("yang", "insert", int64(3000001), collection)
}

func insert(user, repo string, number int64, collection *mongo.Collection) {
	for i := int64(0 + number); i < int64(number+10000); i++ {
		r := record{
			User:   user,
			Repo:   repo,
			ID:     i,
			Status: 1,
			UID:    i + 1,
		}

		if _, err := collection.InsertOne(context.Background(), &r); err != nil {
			log.Fatal(err)
		}
	}
}
