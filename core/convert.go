package core

import (
	lib "go-reloaded/utils"
)

func Converting(s string) string {
	str := ""

	words := lib.SplitWhiteSpaces(s)

	for i := 1; i < len(words); i++ {
		if lib.Contains(words[i], "(hex)") {
			words[i-1] = lib.HexConvert(words[i-1])
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(words[i], "(bin)") {
			words[i-1] = lib.BinConvert(words[i-1])
			words = lib.RemoveSlice(words, i)
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
