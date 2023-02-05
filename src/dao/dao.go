package dao

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Uri = "mongodb://localhost:27017"

func connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData(data map[string]interface{}) {
	client := connect()
	userCollection := client.Database("db").Collection("manifest")
	insertResult, error := userCollection.InsertOne(context.TODO(), data)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(insertResult)

	disconnect(client)
}
