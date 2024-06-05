package core

import lib "go-reloaded/utils"

func Quote(filecontent string) string {
	quoteBeginning := -1
	for i := 0; i < len(filecontent); i++ {
		char := string(filecontent[i])
		if char == "'" {
			if quoteBeginning == -1 {
				if i < len(filecontent)-1 {
					if filecontent[i+1] == ' ' {
						filecontent = lib.RemoveChar(filecontent, i+1)
					}
				}
				quoteBeginning = i
			} else if quoteBeginning != -1 {
				if filecontent[i-1] == ' ' {
					filecontent = lib.RemoveChar(filecontent, i-1)
					quoteBeginning = -1
				}
			}
		}
	}
	return filecontent
}
