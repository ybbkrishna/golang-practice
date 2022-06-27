package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	x, _ := json.Marshal(map[string]float64{
		"a": 1,
		"b": 2,
		"c": math.NaN(),
	})
	fmt.Println(string(x))
}
