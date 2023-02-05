package main

import (
	"app/dao"
	"app/utils"
	"app/view"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// 프로그램 시작 지점
	for {
		/*
			1. 데이터 입력
			2. ID 목록 조회
			3. 데이터 출력
			4. 종료
			메뉴를 선택하세요:
		*/
		menu := view.GetMenu()

		// 프로그램 종료
		if menu == 4 {
			view.PrintToExit()
			break
		}
		// 모든 지점의 error를 반환 받아 예외 처리
		err := navigate(menu)
		if err != nil {
			utils.ExceptionHandler(err)
			continue
		}
	}
}

// 번호에 지정된 함수로 전달
func navigate(menu int64) error {
	var err error
	switch menu {
	case 1:
		err = doInsert()
	case 2:
		err = doGetIdList()
	case 3:
		err = doGetData()

	// 지정된 값이 들어오지 않으면 재입력 받음
	default:
		view.PrintWrongInput()
	}
	return err
}

// 1. 데이터 입력
func doInsert() error {
	// 파일이름 문자열로 받아서 parser에 전달
	fileName := view.GetFileInput()

	// YAML -> JSON 변환
	jsonData, err := utils.YamlToJson(fileName)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	// mongoDB에 입력할 수 있는 값으로 변환
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	// DB 입력
	err = dao.InsertData(data)

	return err
}

// 2. ID 목록 조회
func doGetIdList() error {
	result, err := dao.FindIdList()
	if err != nil {
		return err
	}
	for _, res := range result {
		fmt.Println(res)
	}

	return nil
}

// 3. 데이터 출력
func doGetData() error {
	id := view.GetIdInput()
	data, err := dao.FindDataById(id)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	yamlData, err := utils.JsonToYaml(jsonData)
	if err != nil {
		return err
	}
	path, _ := os.Getwd()
	fileName := path + "\\data.yaml"
	view.PrintYamlData(fileName, yamlData)

	return err
}
