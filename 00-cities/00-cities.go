// Arrange cities by length of the name
package main

import "fmt"

func main() {
	var city = []string{"Surat", "Louisiana", "Montana", "Austin", "Fort", "Gordan", "Henryfort",
		"Illinois", "JacksonHill", "Kalamazoo", "Maine", "Madurai", "Boise", "Boston",
		"Carolina", "NewHampshire", "NewYork", "Ohio", "Utah", "Pennsylvania", "Albama",
		"Phoenix", "Queens", "RhodeIsland", "Minnesota", "Williamsport", "St. Louis",
		"Tennessee", "Vermont", "Wisconsin", "Xenia", "Albama", "York", "Zion", "Davis", "Delhi",
	}

	x := 0
	for j := range city {
		for i := range city {
			if i < (len(city)-1) && (len(city[i]) > len(city[i+1])) {
				temp := city[i]
				city[i] = city[i+1]
				city[i+1] = temp
			}
		}
		x = j
	}
	fmt.Println(x)
	fmt.Println(city)

	for _, c := range city {
		fmt.Printf("%s has total %d number of characters\n", c, len(c))
	}
}
