package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	argstest := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("You need to specify the dictionary file.")
	} else {
		fmt.Println(argstest)
		getDictionary()
	}
}

func setDictionary() {
	fmt.Println("this will work eventually")
}
