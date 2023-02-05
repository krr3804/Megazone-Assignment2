package view

import (
	"app/dao"
	"fmt"
	"os"
)

func PrintToExit() {
	fmt.Println("프로그램이 종료되었습니다.")
}

func PrintInsertSuccess() {
	fmt.Println("데이터가 정상적으로 입력되었습니다.")
}

// ID 목록 출력
func PrintIdList(result []dao.Data) {
	fmt.Println("현재 조회 가능한 데이터 목록: ")
	for _, res := range result {
		fmt.Println(res.ID)
	}
}

// YAML 데이터 출력
func PrintYamlData(fileName string, yamlData []byte) error {
	fmt.Println("데이터가 출력되었습니다.")
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(yamlData)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlData))
	return nil
}
