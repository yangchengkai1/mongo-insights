package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var D []mongo.WriteModel

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

	newInsertOne := mongo.NewInsertOneModel()
	newInsertOne.SetDocument(bson.M{"_id": 30200015, "uid": 3020000, "status": 1})
	D = append(D, newInsertOne)

	newUpdateOne := mongo.NewUpdateOneModel()
	newUpdateOne.SetFilter(bson.M{"_id": bson.M{"$eq": 3020005}})
	newUpdateOne.SetUpdate(bson.M{"$set": bson.M{"uid": 3030000}})
	D = append(D, newUpdateOne)

	newUpdateMany := mongo.NewUpdateManyModel()
	newUpdateMany.SetFilter(bson.M{"user": bson.M{"$eq": "yang"}})
	newUpdateMany.SetUpdate(bson.M{"$set": bson.M{"status": 0}})
	D = append(D, newUpdateMany)

	newReplaceOne := mongo.NewReplaceOneModel()
	newReplaceOne.SetFilter(bson.M{"_id": bson.M{"$eq": 3020000}})
	D = append(D, newReplaceOne)

	newDeleteOne := mongo.NewDeleteOneModel()
	newDeleteOne.SetFilter(bson.M{"status": bson.M{"$eq": 1}})
	newReplaceOne.SetReplacement(bson.M{"status": 1, "root": 1})
	D = append(D, newDeleteOne)

	newDelete := mongo.NewDeleteManyModel()
	newDelete.SetFilter(bson.M{"_id": bson.M{"$eq": 3010000}})
	D = append(D, newDelete)

	result, err := collection.BulkWrite(context.Background(), D)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
