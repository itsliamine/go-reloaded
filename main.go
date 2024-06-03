package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		return
	}

	inputFilename := args[1]
	//outputFilename := args[2]

	content, err := os.ReadFile(inputFilename)

	if err == nil {
		return
	}

	fmt.Println(content)
}
