package main

import "strings"

// start: Jun 18 2021 5:50 PM
// end: Jun 18 2021 6:25 PM

func fullJustify(words []string, maxWidth int) []string {
	lines := makeLines(words, maxWidth)
	text := []string{}
	for i := 0; i < len(lines)-1; i++ {
		text = append(text, insertSpaces(lines[i], maxWidth))
	}
	lastLine := strings.Join(lines[len(lines)-1], " ")
	lastLine = lastLine + makeSpace(maxWidth-len(lastLine))
	text = append(text, lastLine)
	return text
}

func insertSpaces(words []string, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + makeSpace(maxWidth-len(words[0]))
	}
	value := ""
	lenOfAllWords := 0
	for i := 0; i < len(words); i++ {
		lenOfAllWords += len(words[i])
	}
	noOfSpaces := maxWidth - lenOfAllWords
	spacesPerSlot := noOfSpaces / (len(words) - 1)
	spacesLeft := noOfSpaces % (len(words) - 1)
	for i := 0; i < len(words)-1; i++ {
		value += words[i] + makeSpace(spacesPerSlot)
		if spacesLeft > 0 {
			spacesLeft--
			value += " "
		}
	}
	return value + words[len(words)-1]
}

func makeSpace(width int) string {
	value := ""
	for i := 0; i < width; i++ {
		value += " "
	}
	return value
}

func makeLines(words []string, maxWidth int) [][]string {
	lines := [][]string{}
	line := []string{}
	lineLen := 0
	for i := 0; i < len(words); i++ {
		if len(words[i])+lineLen > maxWidth {
			lines = append(lines, line)
			line = []string{}
			lineLen = 0
		}
		line = append(line, words[i])
		lineLen = lineLen + len(words[i]) + 1
	}
	lines = append(lines, line)
	return lines
}
