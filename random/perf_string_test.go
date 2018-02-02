package main

import (
	"testing"
)

func BenchmarkFormatLogic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatLogic("(((()()()()((((sdfjasldkfjas90(adsfasdfa) + (ASDF) (ASD")
	}
}

func BenchmarkFormatLogicReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatLogicReplace("(((()()()()((((sdfjasldkfjas90(adsfasdfa) + (ASDF) (ASD")
	}
}
