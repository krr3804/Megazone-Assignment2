package view

import (
	"fmt"
	"strconv"
)

func GetMenu() int64 {
	var input string
	fmt.Println("1. 데이터 입력\n2. ID 목록 조회\n3. 데이터 출력\n4. 종료\n메뉴를 선택하세요: ")
	fmt.Scan(&input)
	return validateMenuInput(input)
}

func GetFileInput() string {
	var input string
	fmt.Println("파일명을 입력하세요: ")
	fmt.Scan(&input)

	return input
}

func GetIdInput() string {
	var input string
	fmt.Println("ID를 입력하세요: ")
	fmt.Scan(&input)

	return input
}

// func validateFileInput(input string) bool {

// }

func validateMenuInput(input string) int64 {
	in, err := strconv.ParseInt(input, 10, 8)
	if err != nil || (in < 1 || in > 4) {
		return 0
	}
	return in
}
