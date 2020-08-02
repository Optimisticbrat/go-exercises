package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//func printMapElements(mapvalues map[string]int) {
//	for key, value := range mapvalues {
//		fmt.Println(key, "->", value)
//	}
//}

func printMapElements(statescapitol map[string]string) {
	for key, value := range statescapitol {
		fmt.Println(key, "->", value)
	}
}

/*func mostFreq(words map[string]int) (string, int) {
	var freqWord string
	maxOccurance := 0
	for key, value := range words {
		if value > maxOccurance {
			freqWord = key
			maxOccurance = value
		}
	}
	return freqWord, maxOccurance
}*/

func mapstatescapitol(states []string) map[string]string {

	statescapitol := make(map[string]string)
	for _, state := range states {
		sctext := strings.Split(state, ",")
		statescapitol[sctext[0]] = sctext[1]
	}
	return statescapitol
}

func main() {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile("statescapitol.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen

	// text := string(content)
	var lines []string
	lines = strings.Split(string(content), "\n")
	statescapitol := mapstatescapitol(lines)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter state:")
	state, _ := reader.ReadString('\n')
	state = strings.TrimSpace(state)
	if capital, ok := statescapitol[state]; ok {
		fmt.Println("Capital of :", state, " is: ", capital)
	} else {
		fmt.Println("Invalid US state")
	}

}
