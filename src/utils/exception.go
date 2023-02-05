package utils

import "fmt"

// 받은 error를 출력
func ExceptionHandler(err error) error {
	if err != nil {
		fmt.Println("에러 발생:")
		fmt.Println(err)
	}

	return err
}
