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
	dictionaryPath := getDictionaryPath()
	dictStartTime := time.Now()

	// Get the dictionary file.
	dictionary, sortedWordsDictionary, dictError := getDictionary(dictionaryPath)
	if dictError != nil {
		panic(dictError)
	}

	dictEndTime := time.Now()
	fmt.Println(dictEndTime.Sub(dictStartTime))

	inputLoop(dictionary, sortedWordsDictionary)
}

// Infinite loop of user input.
func inputLoop(dictionary []string, sortedWordsDictionary []string) {
	var word string
	var wordError error
	// Prompt the user to input a word.
	word, wordError = getUserInput()
	if wordError != nil {
		panic(wordError)
	}

	for !strings.EqualFold(word, "exit") {
		anagramStartTime := time.Now()

		// Collect anagrams out of dictionary
		anagrams, anagError := getAnagrams(dictionary, sortedWordsDictionary, word)
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
			panic(wordError)
		}
	}
}

func getDictionaryPath() string {
	var dictionaryPath string
	flag.StringVar(&dictionaryPath, "dict", "./dictionary.txt", "The path to the dictionary file.")
	flag.Parse()

	fmt.Println("Loading dictionary from " + dictionaryPath)
	return dictionaryPath
}

// Go fetch the dictionary from the specified path
func getDictionary(path string) ([]string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	var lines []string
	var sortedWordLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		sortedWordLines = append(sortedWordLines, SortString(scanner.Text()))
	}

	return lines, sortedWordLines, scanner.Err()
}

// Display a prompt for user input and then return the input.
func getUserInput() (string, error) {
	fmt.Print("\nEnter a word (Enter exit to quit): ")

	textReader := bufio.NewReader(os.Stdin)
	word, err := textReader.ReadString('\n')
	if err != nil {
		return "", err
	}

	output := strings.Replace(word, "\n", "", -1)

	return output, err
}

// Gather all of the anagrams of the given word.
func getAnagrams(dictionary []string, sortedWordDictionary []string, word string) ([]string, error) {
	var output []string
	wordLetters := SortString(word)

	for i := 0; i < len(sortedWordDictionary); i++ {
		if strings.EqualFold(wordLetters, sortedWordDictionary[i]) && !strings.EqualFold(word, dictionary[i]) {
			output = append(output, dictionary[i])
		}
	}

	return output, nil
}

// ------------------------------------------------------------------------------------------------------------------

type sortRunes []rune

// Sort a words letters alphabetically
func (runeList sortRunes) Less(i, j int) bool {
	return runeList[i] < runeList[j]
}

func (runeList sortRunes) Swap(i, j int) {
	runeList[i], runeList[j] = runeList[j], runeList[i]
}

func (runeList sortRunes) Len() int {
	return len(runeList)
}

func SortString(word string) string {
	runes := []rune(word)
	sort.Sort(sortRunes(runes))
	return string(runes)
}
