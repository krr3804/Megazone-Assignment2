package view

import (
	"fmt"
)

const (
	WrongInput     string = "잘못된 입력입니다."
	WrongFileInput string = "잘못된 형식의 파일입니다."
	Exit           string = "프로그램이 종료되었습니다."
)

func PrintWrongInput() {
	fmt.Println(WrongInput)
}

func PrintWrongFileInput() {
	fmt.Println(WrongFileInput)
}

func PrintToExit() {
	fmt.Println(Exit)
}

func PrintYamlData(fileName string, yamlData []byte) {
	// file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// _, err = file.Write(yamlData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(string(yamlData))
}
