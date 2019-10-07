package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type record struct {
	User   string `bson:"user,omitempty"`
	Repo   string `bson:"repo"`
	ID     int    `bson:"id"`
	Status int    `bson:"status"`
}

func main() {
	go github()

	//go yuque()

	//time.Sleep(1 * time.Minute)
}

func github() {
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

	collection := client.Database("github").Collection("record")

	for i := 0; i < 100000; i++ {
		r := record{
			User:   "yankai",
			Repo:   "comet",
			ID:     i,
			Status: 1,
		}

		if _, err := collection.InsertOne(context.Background(), &r); err != nil {
			log.Fatal(err)
		}
	}
}

/*
func github() {
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

	collection := client.Database("github").Collection("record")
	for i := 0; i < 20000; i++ {
		result, err := collection.UpdateOne(
			context.Background(),
			bson.D{
				{"id", i},
			},
			bson.D{
				{"$set", bson.D{
					{"user", "kai"},
				}},
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("user", result)
	}
}

func yuque() {
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

	collection := client.Database("github").Collection("record")
	for i := 0; i < 20000; i++ {
		result, err := collection.UpdateMany(
			context.Background(),
			bson.D{
				{"id", i},
			},
			bson.D{
				{"$set", bson.D{
					{"repo", "repo"},
				}},
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("repo", result)
	}
}
*/
