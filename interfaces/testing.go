package main

import (
  "github.com/Knetic/govaluate"
	"fmt"
  "reflect"
  "errors"
)

func main()  {
  testExp()
}

func testMap() {
  var keys = map[string]interface{} {
    "hello": [...]int {1,2,3},
  }
  var params = map[string]interface{} {
    "keys": keys,
  }
  exp, err := govaluate.NewEvaluableExpression("keys")
  if err != nil {
    fmt.Println(err)
  }
  val, err := exp.Evaluate(params)
  fmt.Println(val, err)
}

func testArray() {
  var keys map[string]interface{} = map[string]interface{} {
    "hello": [...]int {1,2,3},
  }
  exp, err := govaluate.NewEvaluableExpression("{1:2,2:3}")
  //exp, err := govaluate.NewEvaluableExpression("(1,2,3,5)")
  if err != nil {
    fmt.Println(err)
  }
  val, err := exp.Evaluate(keys)
  fmt.Println(reflect.TypeOf(val))
  fmt.Println(err)
  fmt.Println(val)
}

func Contains(
	args ...interface{},
) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New("Not enough arguments")
	}
	arr := reflect.ValueOf(args[0])
	fmt.Println(args)
	//TODO: log the types
	if arr.Kind() != reflect.Array &&
		arr.Kind() != reflect.Slice {
		return nil, errors.New("Expect arg0 to be an array or slice")
	}
	for i := 0; i < arr.Len(); i++ {
		val := arr.Index(i).Interface()
		if reflect.DeepEqual(val, args[1]) {
			return true, nil
		}
	}
	return false, nil
}

func testFunc() {
  functions := map[string]govaluate.ExpressionFunction {
      "contains": Contains,
  }
  expString := "contains((1,2,3,10), 10)"
  expression, _ := govaluate.NewEvaluableExpressionWithFunctions(expString, functions)

  result, _ := expression.Evaluate(nil)
  fmt.Println(result)
}

func testExp() {
  expString := "((1,2), 10, (1,2,3))"
  expression, _ := govaluate.NewEvaluableExpression(expString)
  result, _ := expression.Evaluate(nil)
  fmt.Println(result)
}
