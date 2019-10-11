package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type record struct {
	ID     int64  `bson:"id"`
	User   string `bson:"user,omitempty"`
	Repo   string `bson:"repo"`
	Status int    `bson:"status"`
}

func main() {
	//var ch = make(chan int, 3)
	//	var single = make(chan int, 1)

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
	go index(collection)
	insert("yankai", "comet", int64(100000), collection)

	fmt.Println("over")
}

func insert(user, repo string, number int64, collection *mongo.Collection) {

	for i := int64(0 + number); i < int64(number+100000); i++ {
		r := record{
			User:   user,
			Repo:   repo,
			ID:     i,
			Status: 1,
		}

		if _, err := collection.InsertOne(context.Background(), &r); err != nil {
			log.Fatal(err)
		}
	}
}

func index(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys: bsonx.Doc{{"id", bsonx.Int64(1)}},
	}

	if _, err := collection.Indexes().CreateOne(context.Background(), indexModel); err != nil {
		log.Fatal(err)
	}
}
