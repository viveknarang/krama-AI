package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CLIENT mongo client
var CLIENT *mongo.Client

func connectDB(url string, port string) *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://" + url + ":" + port)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB at " + url + ":" + port)

	CLIENT = client
	return client

}

func find(db string, collec string, filter interface{}) interface{} {

	var result interface{}

	collection := CLIENT.Database(db).Collection(collec)

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result

}

func insert(db string, collec string, document interface{}) bool {

	collection := CLIENT.Database(db).Collection(collec)

	insertResult, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return true

}

func update(db string, collec string, filter interface{}, update interface{}) bool {

	collection := CLIENT.Database(db).Collection(collec)

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return true

}

func delete(db string, collec string, deleteCriteria interface{}) {

	collection := CLIENT.Database(db).Collection(collec)

	deleteResult, err := collection.DeleteMany(context.TODO(), deleteCriteria)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

}

func disconnectDB() {

	err := CLIENT.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")

}
