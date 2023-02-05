package utils

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

// YAML -> JSON Parser
func YamlToJson(fileName string) ([]byte, error) {
	// YAML 파일 입력 받음
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// JSON 형태의 []byte 자료형으로 변환 후 리턴
	jsonData, err := yaml.YAMLToJSON(yamlFile)

	return jsonData, err
}

// JSON -> YAML Parser
func JsonToYaml(jsonData []byte) ([]byte, error) {
	// YAML 형태의 []byte 자료형으로 변환 후 리턴
	yamlData, err := yaml.JSONToYAML(jsonData)

	return yamlData, err
}
