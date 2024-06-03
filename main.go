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

	fmt.Println(filecontent)

}
