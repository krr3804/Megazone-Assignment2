package utils

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

func YamlToJson(fileName string) []byte {
	yamlFile, err := ioutil.ReadFile(fileName)
	ExceptionHandler(err)
	jsonData, err := yaml.YAMLToJSON(yamlFile)
	ExceptionHandler(err)
	return jsonData
}

func JsonToYaml(jsonData []byte) []byte {
	yamlData, err := yaml.JSONToYAML(jsonData)
	ExceptionHandler(err)

	return yamlData
}
