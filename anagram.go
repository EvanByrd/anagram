package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Get the dictionary file.
	dictionary, dictError := getDictionary("./dictionary.txt")
	if dictError != nil {
		panic(dictError)
	}

	var word string
	var wordError error
	// Prompt the user to input a word.
	word, wordError = getUserInput()
	if wordError != nil {
		panic(dictError)
	}

	for !strings.EqualFold(word, "exit\n") {
		// Collect anagrams out of dictionary
		anagrams, anagError := getAnagrams(dictionary, word)
		if anagError != nil {
			panic(anagError)
		}

		fmt.Println(anagrams)

		// Prompt the user to input another word.
		word, wordError = getUserInput()
		if wordError != nil {
			panic(dictError)
		}

	}
}

// Go fetch the dictionary from the specified path
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

// Display a prompt for user input and then return the input.
func getUserInput() (string, error) {
	textReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	word, err := textReader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return word, err
}

// Gather all of the anagrams of the given word.
func getAnagrams(dictionary []string, word string) ([]string, error) {
	var output []string
	wordLetters := SortString(word)
	for i := 0; i < len(dictionary); i++ {
		if strings.EqualFold(wordLetters, SortString(dictionary[i]+"\n")) && !strings.EqualFold(word, dictionary[i]+"\n") {
			output = append(output, dictionary[i])
		}
	}
	return output, nil
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
