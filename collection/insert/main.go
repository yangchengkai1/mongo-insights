package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type record struct {
	ID     int64  `bson:"_id"`
	User   string `bson:"user"`
	Repo   string `bson:"repo"`
	Status int    `bson:"status"`
	UID    int64  `bson:"uid"`
}

func main() {
	var begin = make(chan bool)
	var done = make(chan bool)

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

	collection := client.Database("stocks").Collection("nasdaq")

	go index(collection, begin, done)
	go insert(collection, begin)

	select {
	case <-done:
		break
	}
}

func index(collection *mongo.Collection, begin, done chan bool) {
	indexModel := mongo.IndexModel{
		Keys: bsonx.Doc{{"none", bsonx.Int64(1)}},
	}

	startIndexTime := time.Now()
	begin <- true
	if _, err := collection.Indexes().CreateOne(context.Background(), indexModel); err != nil {
		log.Fatal(err)
	}
	endIndexTime := time.Now()

	duration := float64(1.0*(endIndexTime.UnixNano()-startIndexTime.UnixNano())) / float64(time.Second)
	log.Println("create index duration :", duration)

	done <- true
}

func insert(collection *mongo.Collection, begin chan bool) {
	<-begin

	for i := int64(0); i < int64(4000000); i++ {
		r := record{
			User:   "yankai",
			Repo:   "client",
			ID:     i,
			Status: 1,
			UID:    i,
		}

		startInsertTime := time.Now()
		if _, err := collection.InsertOne(context.Background(), &r); err != nil {
			log.Fatal(err)
		}
		endInsertTime := time.Now()

		duration := float64(1.0*(endInsertTime.UnixNano()-startInsertTime.UnixNano())) / float64(time.Second)
		log.Println("insert duration :", duration)
	}
}

func init() {
	file := "./" + "stock" + ".txt"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile) // 将文件设置为log输出的文件
}
