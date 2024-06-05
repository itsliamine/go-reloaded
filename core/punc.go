package core

import (
	lib "go-reloaded/utils"
)

func Punc(filecontent string) string {
	for i := 1; i < len(filecontent); i++ {
		char := string(filecontent[i])

		if char == "," || char == "." || char == ":" || char == ";" || char == "?" || char == "!" {
			removed := false
			if string(filecontent[i-1]) == " " {
				filecontent = lib.RemoveChar(filecontent, i-1)
				removed = true
			}
			if i < len(filecontent)-1 {
				if removed {
					if string(filecontent[i]) != " " {
						filecontent = lib.Insert(filecontent, i, ' ')
					}
				} else if string(filecontent[i+1]) != " " {
					filecontent = lib.Insert(filecontent, i, ' ')
				}
			}

		}
	}

	return filecontent
}
