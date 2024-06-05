package core

import (
	lib "go-reloaded/utils"
)

func Converting(s string) string {
	str := ""

	words := lib.SplitWhiteSpaces(s)

	for i, w := range words {
		if lib.Contains(w, "(hex)") {
			words[i-1] = lib.HexConvert(words[i-1])
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(w, "(bin)") {
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
