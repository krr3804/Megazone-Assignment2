package view

import (
	"fmt"
)

func PrintToExit() {
	fmt.Println("프로그램이 종료되었습니다.")
}

func PrintIdList(result []interface{}) {
	for _, res := range result {
		fmt.Println(res)
	}
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
