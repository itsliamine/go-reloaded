package core

import (
	lib "go-reloaded/utils"
)

func Quotes(s string) string {
	str := ""
	words := lib.SplitWhiteSpaces(s)

	quoteBeginning := -1

	for i, word := range words {
		if word == "'" {
			if quoteBeginning == -1 {
				quoteBeginning = i
				words[i+1] = "'" + words[i+1]
				words = lib.RemoveSlice(words, i)
			} else {
				words[i-1] += word
				words = lib.RemoveSlice(words, i)
				quoteBeginning = -1
			}
		} else if len(word) > 1 && lib.Contains(word[1:len(word)-1], "'") {
			index := lib.Index(word, "'")
			if quoteBeginning != -1 {
				words[i-1] += word
				words = lib.RemoveSlice(words, i)
			} else {
				words[i] = word[:index]
				words = lib.Insert(words, i+1, word[index:])
			}
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
