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
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)
	if !src.FileOutput(filecontent, args[2]) {
		fmt.Println("Couldn't save file")
	}

	fmt.Println("Output: ")
	fmt.Println(filecontent)
	fmt.Println()
	fmt.Println("Content saved in file", args[2])
}
