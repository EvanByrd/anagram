package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Get the dictionary file.
	dictionary, dictError := getDictionary("./dictionary.txt")
	if dictError != nil {
		panic(dictError)
	}

	// Prompt the user to input a word.
	word, wordError := getUserInput()
	if wordError != nil {
		panic(dictError)
	}

	// Collect anagrams out of dictionary
	anagrams, anagError := getAnagrams(dictionary, word)
	if anagError != nil {
		panic(anagError)
	}

	fmt.Println(anagrams)
}

func getDictionary(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getUserInput() (string, error) {
	textReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	word, err := textReader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return word, err
}

func getAnagrams(dictionary []string, word string) ([]string, error) {
	return nil, nil
}
