package domain

import (
	"encoding/json"
	"log"
)

type Data struct {
	JsonData interface{}
}

func (data *Data) MapDataToMongo(jsonStr string) {
	err := json.Unmarshal([]byte(jsonStr), &data.JsonData)
	if err != nil {
		log.Fatal(err)
	}
}
