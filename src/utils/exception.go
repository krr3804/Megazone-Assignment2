package utils

import (
	"fmt"
	"log"
)

const (
	WrongInput     string = "잘못된 입력입니다."
	WrongFileInput string = "잘못된 형식의 파일입니다."
	IdNotFound     string = "입력하신 ID에 맞는 데이터가 존재하지 않습니다."
)

// 받은 error를 출력
func ExceptionHandler(err error) error {
	fmt.Print("에러 발생: ")
	if err.Error() == "WFI" {
		fmt.Println(WrongFileInput)
	} else if err.Error() == "WI" {
		fmt.Println(WrongInput)
	} else if err.Error() == "INF" {
		fmt.Println(IdNotFound)
	} else {
		fmt.Println("ERROR!")
		log.Fatal(err)
	}

	return err
}
