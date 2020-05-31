package main

import "fmt"

func main() {
	var numbers = []int{5, 2, 4, 7, 6, 3, 8, 1, 5}
	var pairs [][2]int
	for i, num1 := range numbers {
		j := i + 1
		for j < len(numbers) {
			num2 := numbers[j]
			if i != j && num1+num2 == 9 {
				var pair [2]int
				pair[0] = num1
				pair[1] = num2
				pairs = append(pairs, pair)
			}
			j++
		}
	}
	fmt.Println(pairs)
}
