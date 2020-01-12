package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CLIENT mongo client
var CLIENT *mongo.Client

func connectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://" + MongoURL + ":" + MongoPort)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected at " + MongoURL + ":" + MongoPort)

	CLIENT = client
	return client

}

func findMongoDocument(db string, collec string, filter bson.M) []*bson.D {

	var result []*bson.D

	collection := CLIENT.Database(db).Collection(collec)

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	for cur.Next(context.TODO()) {

		var elem bson.D
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}

	return result

}

func insertMongoDocument(db string, collec string, document interface{}) bool {

	collection := CLIENT.Database(db).Collection(collec)

	insertResult, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return insertResult.InsertedID != nil

}

func updateMongoDocument(db string, collec string, filter interface{}, update interface{}) [2]int64 {

	collection := CLIENT.Database(db).Collection(collec)

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	var result [2]int64

	if err != nil {
		log.Fatal(err)
		result[0] = -1
		result[1] = -1
		return result
	}

	result[0] = updateResult.MatchedCount
	result[1] = updateResult.ModifiedCount

	return result

}

func deleteMongoDocument(db string, collec string, deleteCriteria interface{}) int64 {

	collection := CLIENT.Database(db).Collection(collec)

	deleteResult, err := collection.DeleteMany(context.TODO(), deleteCriteria)

	if err != nil {
		log.Fatal(err)
	}

	return deleteResult.DeletedCount

}

func disconnectDB() {

	err := CLIENT.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")

}
