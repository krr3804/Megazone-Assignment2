package utils

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

func YamlToJson(fileName string) ([]byte, error) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	jsonData, err := yaml.YAMLToJSON(yamlFile)

	return jsonData, err
}

func JsonToYaml(jsonData []byte) ([]byte, error) {
	yamlData, err := yaml.JSONToYAML(jsonData)

	return yamlData, err
}
