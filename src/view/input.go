package view

import (
	"fmt"
	"strconv"
)

const (
	WrongInput    string = "잘못된 입력입니다."
	InputFileName string = "파일명을 입력하세요: "
	Exit          string = "프로그램이 종료되었습니다."
)

func PrintMenu() int64 {
	var input string
	fmt.Println(
		"1. 데이터 입력\n",
		"2. ID 목록 조회\n",
		"3. 데이터 출력\n",
		"4. 종료\n",
		"메뉴를 선택하세요: ")
	fmt.Scanln(&input)
	return validateMenuInput(input)
}

func PrintWrongInput() {
	fmt.Println(WrongInput)
}

func PrintFileInput() {
	var input string
	fmt.Println(InputFileName)
	fmt.Scanln(&input)

}

func PrintToExit() {
	fmt.Println(Exit)
}

// func validateFileInput(input string) bool {
// 	if()
// }

func validateMenuInput(input string) int64 {
	in, err := strconv.ParseInt(input, 10, 8)
	if err != nil || (in < 1 || in > 4) {
		return 0
	}
	return in
}
