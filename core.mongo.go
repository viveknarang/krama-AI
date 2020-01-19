package main

import (
	"context"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MONGODBCLIENT mongo client
var MONGODBCLIENT *mongo.Client

func connectDB() bool {

	rlog.Debug("connectDB() handle function invoked ...")

	clientOptions := options.Client().ApplyURI("mongodb://" + MongoURL + ":" + MongoPort)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		rlog.Error("connectDB() Error: " + err.Error())
	}

	MONGODBCLIENT = client

	if pingMongoDB(false) {

		return true

	}

	return false

}

func pingMongoDB(silent bool) bool {

	rlog.Debug("pingMongoDB() handle function invoked ...")

	if MONGODBCLIENT == nil {
		return false
	}

	err := MONGODBCLIENT.Ping(context.TODO(), nil)

	if err != nil {
		rlog.Error("pingMongoDB() Error: " + err.Error())
		return false
	}

	if !silent {
		rlog.Debug("pingMongoDB() MongoDB responding at " + MongoURL + ":" + MongoPort)
	}

	return true

}

func findMongoDocument(db string, collec string, filter bson.M, opts *options.FindOptions) []*bson.D {

	rlog.Debug("findMongoDocument() handle function invoked ...")

	var result []*bson.D

	collection := MONGODBCLIENT.Database(db).Collection(collec)

	cur, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		rlog.Error("findMongoDocument() Error: " + err.Error())
		return nil
	}

	for cur.Next(context.TODO()) {

		var elem bson.D
		err := cur.Decode(&elem)
		if err != nil {
			rlog.Error("findMongoDocument() Error: " + err.Error())
		}
		result = append(result, &elem)
	}

	return result

}

func insertMongoDocument(db string, collec string, document interface{}) bool {

	rlog.Debug("insertMongoDocument() handle function invoked ...")

	collection := MONGODBCLIENT.Database(db).Collection(collec)

	insertResult, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		rlog.Error("insertMongoDocument() Error: " + err.Error())
		return false
	}

	return insertResult.InsertedID != nil

}

func updateMongoDocument(db string, collec string, filter interface{}, update interface{}) [2]int64 {

	rlog.Debug("updateMongoDocument() handle function invoked ...")

	collection := MONGODBCLIENT.Database(db).Collection(collec)

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	var result [2]int64

	if err != nil {
		rlog.Error("updateMongoDocument() Error: " + err.Error())
		result[0] = -1
		result[1] = -1
		return result
	}

	result[0] = updateResult.MatchedCount
	result[1] = updateResult.ModifiedCount

	return result

}

func deleteMongoDocument(db string, collec string, deleteCriteria interface{}) int64 {

	rlog.Debug("deleteMongoDocument() handle function invoked ...")

	collection := MONGODBCLIENT.Database(db).Collection(collec)

	deleteResult, err := collection.DeleteMany(context.TODO(), deleteCriteria)

	if err != nil {
		rlog.Error("deleteMongoDocument() Error: " + err.Error())
	}

	return deleteResult.DeletedCount

}

func disconnectDB() {

	rlog.Debug("disconnectDB() handle function invoked ...")

	err := MONGODBCLIENT.Disconnect(context.TODO())

	if err != nil {
		rlog.Error("disconnectDB() Error: " + err.Error())
	}

	rlog.Debug("disconnectDB() Connection to MongoDB closed.")

}
