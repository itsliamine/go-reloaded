package core

import lib "go-reloaded/utils"

func Vowel(filecontent string) string {
	for i := 0; i < len(filecontent); i++ {
		char := string(filecontent[i])
		if char == "a" {
			if string(filecontent[i+1]) == " " && lib.IsVowel(filecontent[i+2]) {
				filecontent = lib.Insert(filecontent, i+1, 'n')
			}
		}
	}
	return filecontent
}
