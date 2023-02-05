package main

import (
	"app/dao"
	"app/utils"
	"app/view"
	"encoding/json"
	"errors"
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
		err = errors.New("WI")
	}
	return err
}

// 1. 데이터 입력
func doInsert() error {
	// 파일이름 문자열로 받아서 parser에 전달
	fileName, err := view.GetFileInput()
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}
	view.PrintInsertSuccess()
	return nil
}

// 2. ID 목록 조회
func doGetIdList() error {
	result, err := dao.FindIdList()
	if err != nil {
		return err
	}
	view.PrintIdList(result)

	return nil
}

// 3. 데이터 출력
func doGetData() error {
	// 조회할 데이터의 id를 입력 받음
	id := view.GetIdInput()

	// DB에 조회
	data, err := dao.FindDataById(id)
	if err != nil {
		return err
	}

	// 가져온 data를 json 형식으로 변환
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// JSON -> YAML 변환
	yamlData, err := utils.JsonToYaml(jsonData)
	if err != nil {
		return err
	}

	// YAML 데이터 출력
	path, _ := os.Getwd()
	fileName := path + "\\data.yaml"
	err = view.PrintYamlData(fileName, yamlData)

	return err
}
