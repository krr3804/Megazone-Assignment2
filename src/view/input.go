package view

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func GetMenu() int64 {
	var input string
	fmt.Println("1. 데이터 입력\n2. ID 목록 조회\n3. 데이터 출력\n4. 종료\n메뉴를 선택하세요: ")
	fmt.Scan(&input)
	return validateMenuInput(input)
}

func GetFileInput() (string, error) {
	var input string
	fmt.Println("파일명을 입력하세요: ")
	fmt.Scan(&input)
	var err error
	if !validateFileInput(input) {
		err = errors.New("WFI")
	}
	return input, err
}

func GetIdInput() string {
	var input string
	fmt.Println("ID를 입력하세요: ")
	fmt.Scan(&input)

	return input
}

func validateFileInput(input string) bool {
	matched1, _ := regexp.MatchString(".yaml$", input)
	matched2, _ := regexp.MatchString(",yml$", input)

	return matched1 || matched2
}

func validateMenuInput(input string) int64 {
	in, err := strconv.ParseInt(input, 10, 8)
	if err != nil || (in < 1 || in > 4) {
		return 0
	}
	return in
}
