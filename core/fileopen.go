package core

import "os"

func FileOpen(filename string) string {

	content, err := os.ReadFile(filename)

	if err != nil {
		return ""
	}

	return string(content)
}
