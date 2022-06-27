package main

import "strings"

// start: Jun 18 2021 5:05 PM
// end: Jun 18 2021 5:45 PM

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	numMap := []string{"", "Thousand", "Million", "Billion", "Trillion"}
	value := ""
	for i, j := num, 0; i > 0; i, j = i/1000, j+1 {
		if i%1000 == 0 {
			continue
		}
		curr := numberToWordsUnderThousand(i%1000) + " " + numMap[j]
		value = strings.Trim(curr, " ") + " " + value
	}
	return strings.Trim(value, " ")
}

func numberToWordsUnderThousand(num int) string {
	if num < 100 {
		return numberToWordsTwoDigits(num)
	}
	return strings.Trim(numberToWordsTwoDigits(num/100)+" "+"Hundred"+" "+numberToWordsTwoDigits(num%100), " ")
}

func numberToWordsUnderTwenty(num int) string {
	if num < 1 {
		return ""
	}
	numMap := map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}
	return numMap[num]
}

func numberToWordsTwoDigits(num int) string {
	if num < 20 {
		return numberToWordsUnderTwenty(num)
	}
	numMap := map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	if num%10 == 0 {
		return numMap[num/10]
	}
	return numMap[num/10] + " " + numberToWordsUnderTwenty(num%10)
}
