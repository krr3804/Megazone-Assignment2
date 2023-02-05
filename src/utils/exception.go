package utils

import "fmt"

func ExceptionHandler(err error) {
	if err != nil {
		fmt.Println("에러 발생:")
		fmt.Println(err)
	}
}
