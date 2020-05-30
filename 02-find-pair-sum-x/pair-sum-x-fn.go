// From a list of numbers find a pair that sums to X
package main

import "fmt"

func pairSumX(numbers []int, x int) [][2]int {
	var pairs [][2]int
	for i, num1 := range numbers {
		var j = i + 1
		for j < len(numbers) {
			var num2 = numbers[j]
			if i != j && num1+num2 == x {
				var pair [2]int
				pair[0] = num1
				pair[1] = num2
				pairs = append(pairs, pair)
			}
			j++
		}

	}
	return pairs
}

func max(number1, number2 int) int {
	var max = number1
	if number1 < number2 {
		max = number2
	}
	return max
}

func main() {
	var numbers = []int{5, 2, 4, 7, 6, 3, 8, 1, 5}
	pairs := pairSumX(numbers, 8)
	fmt.Println(pairs)
	fmt.Println(pairSumX(numbers, 10))
	fmt.Println(max(5, 8))
}
