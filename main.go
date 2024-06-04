package main

import (
	"fmt"
	src "go-reloaded/core"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		return
	}

	filecontent := src.FileOpen(args[1])
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	filecontent = src.Quotes(filecontent)
	filecontent = src.Punctuation(filecontent)
	filecontent = src.Vowels(filecontent)

	fmt.Println(filecontent)
}
