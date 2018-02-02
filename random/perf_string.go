package main

import "bytes"
import "strings"

func formatLogic(in string) string {
	var buffer bytes.Buffer
	for _, rune := range in {
		if rune == '(' || rune == ')' {
			buffer.WriteRune(' ')
			buffer.WriteRune(rune)
			buffer.WriteRune(' ')
		} else {
			buffer.WriteRune(rune)
		}
	}
	return buffer.String()
}

func formatLogicReplace(in string) string {
	in = strings.Replace(in, "(", " ( ", -1)
	return strings.Replace(in, ")", " ) ", -1)
}
