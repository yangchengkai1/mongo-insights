package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	var ch = make(chan int, 3)

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

	collection := client.Database("yuque").Collection("test")

	go insert("yankai", "comet", int64(0), collection, ch)
	go insert("yangchengkai", "sor", int64(1000000), collection, ch)
	go insert("kai", "draw.io", int64(2000000), collection, ch)
	<-ch
	<-ch
	<-ch
	fmt.Println("over")
}

func insert(user, repo string, number int64, collection *mongo.Collection, ch chan int) {
	for i := int64(0 + number); i < int64(number+1000000); i++ {
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

	ch <- 1
}
