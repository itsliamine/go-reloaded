package core

import (
	lib "go-reloaded/utils"
)

func Punctuation(s string) string {
	str := ""
	words := lib.SplitWhiteSpaces(s)

	for i := 0; i < len(words); i++ {
		if !(words[i] != ";" && words[i] != ":" && words[i] != "." && words[i] != "," && words[i] != "?" && words[i] != "!") {
			words[i-1] += words[i]
			words = lib.RemoveSlice(words, i)
			continue
		} else if lib.GetPrefix(words[i]) {
			words[i-1] += string(words[i][0])
			words = lib.Insert(words, i+1, words[i][1:])
			words = lib.RemoveSlice(words, i)
		} else if lib.Contains(words[i][:len(words[i])-1], ".") {
			index := lib.Index(words[i][:len(words[i])-1], ".")
			words[i-1] += words[i][:index+1]
			words = lib.Insert(words, i+1, words[i][index+1:])
			words = lib.RemoveSlice(words, i)
		}

		for j, char := range words[i] {
			if char != ';' && char != ':' && char != '.' && char != ',' && char != '?' && char != '!' {
				break
			}
			if j == len(words[i])-1 {
				words[i-1] += words[i]
				words = lib.RemoveSlice(words, i)
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
