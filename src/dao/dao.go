package dao

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB URI
const Uri = "mongodb://localhost:27017"

// DB 커넥션 생성
func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	return client, err
}

// DB 커넥션 해제
func disconnect(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())

	return err
}

// DB Insert
func InsertData(data map[string]interface{}) error {
	client, err := connect()
	if err != nil {
		return err
	}
	// 삽입할 컬렉션 로드
	userCollection := client.Database("db").Collection("manifest")

	// 데이터 삽입
	_, err = userCollection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	err = disconnect(client)
	return err
}

// Object id 목록 가져오기
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

	// go 자료형에 mapping
	var result []interface{}
	err = cur.All(context.TODO(), &result)

	return result, err

}

// 단일 데이터 조회
func FindDataById(id string) (map[string]interface{}, error) {
	client, err := connect()
	if err != nil {
		return nil, err
	}
	userCollection := client.Database("db").Collection("manifest")
	// Object Id로 조회
	opt := options.FindOne().SetProjection(bson.D{{"_id", 0}})

	// 매개변수로 받은 id(string)을 objectID 형태로 변환
	objectId, _ := primitive.ObjectIDFromHex(id)
	var result map[string]interface{}

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": objectId}, opt).Decode(&result)
	if err != nil {
		return nil, errors.New("INF")
	}

	return result, err
}
