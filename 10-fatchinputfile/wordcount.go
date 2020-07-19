package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func printMapElements(mapvalues map[string]int) {
	for key, value := range mapvalues {
		fmt.Println(key, "->", value)
	}
}

func mostFreq(words map[string]int) (string, int) {
	var freqWord string
	maxOccurance := 0
	for key, value := range words {
		if value > maxOccurance {
			freqWord = key
			maxOccurance = value
		}
	}
	return freqWord, maxOccurance
}

func mapcount(words []string) map[string]int {
	counter := 0
	counts := make(map[string]int)
	for _, word := range words {
		value, ok := counts[word]
		if ok {
			counts[word] = value + 1
		} else {
			counts[word] = 1
		}
		counter++
	}
	return counts
}

func main() {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile("fatchinginputfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen

	// text := string(content)
	var words []string
	text := string(content)
	words = strings.Split(text, " ")
	mapwords := mapcount(words)
	printMapElements(mapwords)
	mword, nocc := mostFreq(mapwords)
	fmt.Println("Most frequent word is:", mword, "and it occurs for:", nocc)
}
