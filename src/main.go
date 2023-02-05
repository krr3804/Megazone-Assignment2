package main

import (
	"app/dao"
	_ "app/domain"
	"app/utils"
	"app/view"
	"encoding/json"
	"log"
)

func main() {
	for {
		menu := view.GetMenu()
		if menu == 4 {
			view.PrintToExit()
			break
		}
		navigate(menu)
	}
}

func navigate(menu int64) {
	switch menu {
	case 1:
		doInsert()
	case 2:
		doGetIdList()
	case 3:
		doGetData()

	default:
		view.PrintWrongInput()
	}

}

func doInsert() {
	fileName := view.GetFileInput()
	var jsonData []byte = utils.YamlToJson(fileName)
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	dao.InsertData(data)
}

func doGetIdList() {

}

func doGetData() {

}
