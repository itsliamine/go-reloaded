package core

import (
	lib "go-reloaded/utils"
)

func Vowels(s string) string {
	str := ""
	words := lib.SplitWhiteSpaces(s)

	for i, word := range words {
		if word == "a" && lib.IsVowel(words[i+1][0]) {
			words[i] = "an"
		}
	}

	for i, w := range words {
		if i == len(words)-1 {
			str += w
			break
		}
		str += w + " "
	}

	return str
}
