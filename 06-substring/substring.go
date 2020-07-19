package main

import "fmt"

func maxsumstring(numbers []int) (int, []int) {
	var slength = 1
	var sslice []int
	var maxSum = 0
	for slength < len(numbers) {
		i := 0
		for i < len(numbers)-slength {
			var sum = 0
			var j = 0
			var cslice []int
			for j < slength {
				cslice = append(cslice, numbers[i+j])
				sum = sum + numbers[i+j]
				j++
			}
			if sum > maxSum {
				maxSum = sum
				sslice = cslice
			}
			fmt.Println(i, sum)
			i++
		}
		slength++
	}
	return maxSum, sslice
}

func main() {
	var numbers = []int{2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum, sslice := maxsumstring(numbers)
	fmt.Println("Maximum sum is ", sum, "for", sslice)
}
