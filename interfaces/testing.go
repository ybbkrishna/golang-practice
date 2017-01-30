package main

import (
  "github.com/Knetic/govaluate"
	"fmt"
	"reflect"
)

func main()  {
  var keys map[string]interface{} = map[string]interface{} {
    "hello": [...]int {1,2,3},
  }
  exp, err := govaluate.NewEvaluableExpression("hello")
  if err != nil {
    fmt.Println(err)
  }
  val, err := exp.Evaluate(keys)
  fmt.Println(reflect.TypeOf(val))
  fmt.Println(val)
}
