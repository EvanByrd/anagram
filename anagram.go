package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	var dictionaryPath = flag.String("dict", "./dictionary.txt", "")
	dictStartTime := time.Now()

	// Get the dictionary file.
	dictionary, dictError := getDictionary(*dictionaryPath)
	if dictError != nil {
		panic(dictError)
	}
	dictEndTime := time.Now()
	fmt.Println(dictEndTime.Sub(dictStartTime))

	var word string
	var wordError error
	// Prompt the user to input a word.
	word, wordError = getUserInput()
	if wordError != nil {
		panic(dictError)
	}

	for !strings.EqualFold(word, "exit") {
		anagramStartTime := time.Now()

		// Collect anagrams out of dictionary
		anagrams, anagError := getAnagrams(dictionary, word)
		if anagError != nil {
			panic(anagError)
		}

		anagramEndTime := time.Now()

		if len(anagrams) == 0 {
			fmt.Println("No anagrams found for " + word)
		} else {
			fmt.Println(anagrams)
		}
		fmt.Println(anagramEndTime.Sub(anagramStartTime))

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
	fmt.Print("\nEnter a word: ")

	textReader := bufio.NewReader(os.Stdin)
	word, err := textReader.ReadString('\n')
	if err != nil {
		return "", err
	}

	output := strings.Replace(word, "\n", "", -1)

	return output, err
}

// Gather all of the anagrams of the given word.
func getAnagrams(dictionary []string, word string) ([]string, error) {
	var output []string
	wordLetters := SortString(word)

	for i := 0; i < len(dictionary); i++ {
		if strings.EqualFold(wordLetters, SortString(dictionary[i])) && !strings.EqualFold(word, dictionary[i]) {
			output = append(output, dictionary[i])
		}
	}

	return output, nil
}

// Sort a words letters alphabetically
func SortString(word string) string {
	splitWord := strings.Split(word, "")
	sort.Strings(splitWord)
	output := strings.Join(splitWord, "")

	return output
}
