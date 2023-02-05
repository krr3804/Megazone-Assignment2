package utils

import (
	"errors"
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

// YAML -> JSON Parser
func YamlToJson(fileName string) ([]byte, error) {
	// YAML 파일 입력 받음
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("WFI")
	}

	// JSON 형태의 []byte 자료형으로 변환 후 리턴
	jsonData, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		return nil, errors.New("WFI")
	}
	return jsonData, nil
}

// JSON -> YAML Parser
func JsonToYaml(jsonData []byte) ([]byte, error) {
	// YAML 형태의 []byte 자료형으로 변환 후 리턴
	yamlData, err := yaml.JSONToYAML(jsonData)

	return yamlData, err
}
