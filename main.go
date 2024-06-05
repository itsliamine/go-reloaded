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
	fmt.Println(len(filecontent))
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	fmt.Println(len(filecontent))
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)

	fmt.Println(filecontent)
}
