package main

import (
	"fmt"
	"strings"
)

func threeEqualParts(arr []int) []int {
	i, j := 1, 2
	for ; i < len(arr)-1; i++ {
		for j = i + 1; j < len(arr); j++ {
			// fmt.Println(i, j, arr[:i], arr[i:j], arr[j:])
			if compare(arr[:i], arr[i:j]) && compare(arr[i:j], arr[j:]) {
				return []int{i - 1, j}
			}
		}
	}
	return []int{-1, -1}
}

func compare(a, b []int) bool {
	a1 := strings.Join(strings.Fields(strings.Trim(fmt.Sprint(a), "[]")), "")
	a1 = strings.TrimLeft(a1, "0")
	b1 := strings.Join(strings.Fields(strings.Trim(fmt.Sprint(b), "[]")), "")
	b1 = strings.TrimLeft(b1, "0")
	// fmt.Println(a1, b1)
	return a1 == b1
}

func main() {
	fmt.Println(threeEqualParts([]int{0, 1, 0, 1, 1}))
}
