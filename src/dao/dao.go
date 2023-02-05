package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Uri = "mongodb://localhost:27017"

func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	return client, err
}

func disconnect(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())

	return err
}

func InsertData(data map[string]interface{}) error {
	client, err := connect()
	if err != nil {
		return err
	}
	userCollection := client.Database("db").Collection("manifest")
	_, err = userCollection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	err = disconnect(client)
	return err
}

func FindIdList() ([]interface{}, error) {
	client, err := connect()
	if err != nil {
		return nil, err
	}
	userCollection := client.Database("db").Collection("manifest")
	opt := options.Find().SetProjection(bson.D{{"_id", 1}})
	cur, err := userCollection.Find(context.TODO(), bson.D{}, opt)
	if err != nil {
		return nil, err
	}

	var result []interface{}
	err = cur.All(context.TODO(), &result)

	return result, err

}

func FindDataById(id string) (map[string]interface{}, error) {
	client, err := connect()
	if err != nil {
		return nil, err
	}
	userCollection := client.Database("db").Collection("manifest")
	opt := options.FindOne().SetProjection(bson.D{{"_id", 0}})

	objectId, _ := primitive.ObjectIDFromHex(id)
	var result map[string]interface{}

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": objectId}, opt).Decode(&result)

	return result, err
}
