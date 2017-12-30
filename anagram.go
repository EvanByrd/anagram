package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println("You need to specify the dictionary file.")
	} else {
		setAnagram()
	}
}

func setAnagram() {
	fmt.Println("this will work eventually")
	// test
}
