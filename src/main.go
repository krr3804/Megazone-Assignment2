package main

import (
	"app/dao"
	_ "app/domain"
	"app/utils"
	"app/view"
	"encoding/json"
	"fmt"
	"os"
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
	utils.ExceptionHandler(err)
	dao.InsertData(data)
}

func doGetIdList() {
	result := dao.FindIdList()
	for _, res := range result {
		fmt.Println(res)
	}
}

func doGetData() {
	id := view.GetIdInput()
	data := dao.FindDataById(id)
	jsonData, err := json.Marshal(data)
	utils.ExceptionHandler(err)
	yamlData := utils.JsonToYaml(jsonData)
	path, _ := os.Getwd()
	fileName := path + "\\data.yaml"
	view.PrintYamlData(fileName, yamlData)
}
