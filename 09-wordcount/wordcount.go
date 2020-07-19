package main

import (
	"fmt"
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
	para := "Harmon Killebrew (June 29, 1936 – May 17, 2011) was an American professional baseball first baseman, third baseman, and left fielder. During his 22-year career in Major League Baseball, primarily with the Minnesota Twins, Killebrew was a prolific power hitter who, at the time of his retirement, had the fourth most home runs in major league history. Second only to Babe Ruth in home runs in the American League, he was inducted into the National Baseball Hall of Fame in 1984."

	var words []string
	words = strings.Split(para, " ")
	mapwords := mapcount(words)
	printMapElements(mapwords)
	mword, nocc := mostFreq(mapwords)
	fmt.Println("Most frequent word is:", mword, "and it occurs for:", nocc)

	//printMapElements(mapcount(strings.Split(para, " ")))
}
