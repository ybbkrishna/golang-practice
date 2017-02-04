package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

const PRODUCTION_RULES = "./test.yaml"

func LoadRulesFile(
	fname string,
	result interface{},
) error {
	var parsedResult = make(map[interface{}]interface{})
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, parsedResult); err != nil {
		return err
	}
	if filesToExtend := parsedResult["extends"]; filesToExtend != nil {
		switch filesToExtend.(type) {
		case []interface{}:
			files, _ := filesToExtend.([]interface{})
			for _, file := range files{
				fileName, _ := file.(string)
				LoadRulesFile(fileName, result)
			}
		case string:
			file, _ := filesToExtend.(string)
			LoadRulesFile(file, result)
		}
		delete(parsedResult, "extends");
	}
	merge(result.(map[interface{}]interface{}), parsedResult)
	return nil
}

func main() {
  var testConfig = make(map[interface{}]interface{})
  LoadRulesFile(PRODUCTION_RULES, testConfig)
	fmt.Println(testConfig)
}

func merge(
	dest map[interface{}]interface{},
	src map[interface{}]interface{},
) error {
	for k, v := range src {
		if dest[k] == nil {
			dest[k] = v
		} else {
			switch v.(type) {
			case map[interface{}]interface{}:
				switch dest[k].(type) {
				case map[interface{}]interface{}:
					desVal := dest[k].(map[interface{}]interface{})
					merge(desVal, v.(map[interface{}]interface{}))
				default:
					dest[k] = v
				}
			default:
				dest[k] = v
			}
		}
	}
	return nil
}
