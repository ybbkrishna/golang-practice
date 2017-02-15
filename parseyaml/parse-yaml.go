package main

import (
	"io/ioutil"

	"fmt"

	"gopkg.in/yaml.v2"
)

const PRODUCTION_RULES = "./test3.yaml"

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
			for _, file := range files {
				fileName, _ := file.(string)
				LoadRulesFile(fileName, result)
			}
		case string:
			file, _ := filesToExtend.(string)
			LoadRulesFile(file, result)
		}
		delete(parsedResult, "extends")
	}
	merge(result.(map[interface{}]interface{}), parsedResult)
	return nil
}

func main() {
	// var testConfig = make(map[interface{}]interface{})
	// LoadRulesFile(PRODUCTION_RULES, testConfig)
	// val := reflect.ValueOf(testConfig["test"])
	// switch val.Kind() {
	// case reflect.Map:
	// 	fmt.Println("its a map")
	// case reflect.Array:
	// 	fallthrough
	// case reflect.Slice:
	// 	fmt.Println("its array of maps")
	// default:
	// 	fmt.Println(val.Kind())
	// }

	// fmt.Println(testConfig)
	parseTest3()
}

func parseTest3() {
	var testConfig = make(map[interface{}]interface{})
	LoadRulesFile(PRODUCTION_RULES, testConfig)
	val := testConfig["test"].([]interface{})
	val2 := val[3].(map[interface{}]interface{})
	fmt.Println(val2)
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
