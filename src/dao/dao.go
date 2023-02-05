package dao

import (
	"app/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Uri = "mongodb://localhost:27017"

func connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	utils.ExceptionHandler(err)
	return client
}

func disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	utils.ExceptionHandler(err)
}

func InsertData(data map[string]interface{}) {
	client := connect()
	userCollection := client.Database("db").Collection("manifest")
	insertResult, err := userCollection.InsertOne(context.TODO(), data)
	utils.ExceptionHandler(err)

	fmt.Println(insertResult)

	disconnect(client)
}

func FindIdList() []interface{} {
	client := connect()
	opt := options.Find().SetProjection(bson.D{{"_id", 1}})
	userCollection := client.Database("db").Collection("manifest")

	cur, err := userCollection.Find(context.TODO(), bson.D{}, opt)
	utils.ExceptionHandler(err)

	var result []interface{}
	err = cur.All(context.TODO(), &result)
	utils.ExceptionHandler(err)

	return result

}

func FindDataById(id string) map[string]interface{} {
	client := connect()
	opt := options.FindOne().SetProjection(bson.D{{"_id", 0}})
	userCollection := client.Database("db").Collection("manifest")
	objectId, _ := primitive.ObjectIDFromHex(id)
	var result map[string]interface{}
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": objectId}, opt).Decode(&result)
	utils.ExceptionHandler(err)

	return result
}
