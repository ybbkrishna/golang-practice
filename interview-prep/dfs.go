package main

import "fmt"

var digitToCharMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return digitToCharMap[string(digits[0])]
	}
	letterCombinations := letterCombinations(digits[1:])
	combinations := []string{}
	for i := 0; i < len(digitToCharMap[string(digits[0])]); i++ {
		for j := 0; j < len(letterCombinations); j++ {
			val := digitToCharMap[string(digits[0])][i] + letterCombinations[j]
			combinations = append(combinations, val)
		}
	}
	return combinations
}

func main() {
	fmt.Println(letterCombinations("22"))
}
