package core

import lib "go-reloaded/utils"

func Letters(s string) string {
	str := ""

	words := lib.SplitWhiteSpaces(s)

	for i, w := range words {
		if lib.Contains(w, "(cap,") {
			num := lib.Atoi(string(words[i+1][0]))
			for x := i; x >= i-num; x-- {
				words[x] = lib.Capitalize(words[x])
			}
			words = lib.RemoveSlice(words, i)
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(w, "(cap)") {
			words[i-1] = lib.Capitalize(words[i-1])
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(w, "(low,") {
			num := lib.Atoi(string(words[i+1][0]))
			for x := i; x >= i-num; x-- {
				words[x] = lib.ToLower(words[x])
			}
			words = lib.RemoveSlice(words, i)
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(w, "(low)") {
			words[i-1] = lib.ToLower(words[i-1])
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(w, "(up,") {
			num := lib.Atoi(string(words[i+1][0]))
			for x := i; x >= i-num; x-- {
				words[x] = lib.ToUpper(words[x])
			}
			words = lib.RemoveSlice(words, i)
			words = lib.RemoveSlice(words, i)
		}

		if lib.Contains(w, "(up)") {
			words[i-1] = lib.ToUpper(words[i-1])
			words = lib.RemoveSlice(words, i)
		}

		str += w
	}

	return str
}
