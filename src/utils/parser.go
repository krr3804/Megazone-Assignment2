package utils

import (
	"io/ioutil"
	"log"

	"sigs.k8s.io/yaml"
)

func YamlToJson(fileName string) []byte {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
