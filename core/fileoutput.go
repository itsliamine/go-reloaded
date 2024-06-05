package core

import (
	"os"
)

func FileOutput(fileContent, outputFileName string) bool {
	f, err := os.Create(outputFileName)
	if err != nil {
		return false
	}

	_, err = f.WriteString(fileContent + "\n")
	if err != nil {
		return false
	}

	err = f.Close()

	return err == nil
}
