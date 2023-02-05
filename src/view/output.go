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
