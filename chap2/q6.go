/*Write a function that calculates the average of a float64 slice.*/
package main

import (
	"fmt"
)

func main() {
	arr := []float64{1, 2, 3, 4, 5, 6}
	val := average(arr)
	fmt.Println(val)
}

func average(arr []float64) float64 {
	var sum, avg float64 = 0, 0
	for _, val := range arr {
		sum += val
	}
	avg = sum / float64(len(arr))
	return avg
}
