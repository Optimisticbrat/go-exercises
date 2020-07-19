package main

import "fmt"

func naiveSort(numbers []int) {
	j := 0
	countSwap := 0
	counter := 0
	for j < len(numbers)-1 {
		i := 0
		for i < len(numbers)-1 {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				countSwap++
			}
			counter++
			i++
		}
		j++
	}
	//fmt.Println(numbers)
	percentSwaps := 100 * countSwap / counter
	fmt.Println("Total number of checks", counter, "number of swap", countSwap, "Percentage number of swap", percentSwaps)
}

func mapSort(numbers []int) []int {
	counter := 0

	counts := make(map[int]int)
	for _, number := range numbers {
		value, ok := counts[number]
		if ok {
			counts[number] = value + 1
		} else {
			counts[number] = 1
		}
		counter++
	}
	// Create an empty list
	var list []int
	i := 0
	for i <= 9 {
		value, ok := counts[i]
		if ok {
			j := 0
			for j < value {
				list = append(list, i)
				j++
				counter++
			}
		}
		i++
	}
	fmt.Println("number of steps =", counter)
	return list
}

func main() {
	numbers := []int{7, 4, 3, 7, 6, 5, 1, 9, 8, 2, 2, 6, 4, 0}
	original := make([]int, len(numbers))
	copy(original, numbers)
	//numbers := []int{9, 9, 8, 7, 7, 6, 5, 4, 3, 2, 1, 0}
	naiveSort(numbers)
	fmt.Println("Original list = ", original, "\nSortage list =", numbers)
	sortedlist := mapSort(original)

	fmt.Println("Original list = ", original, "\nMap Sortage list =", sortedlist)
}
