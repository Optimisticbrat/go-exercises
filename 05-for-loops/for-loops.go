package main

import "fmt"

func main() {
	//numbers := []int{1, 4, 5, 2, 6, 7, 9}
	//var i = 0
	//fmt.Println(numbers)
	//for i < len(numbers) {
	//	fmt.Println("value is ", numbers[i], "at index", i)
	//	i++
	//}
	fruits := []string{"banana", "apple", "grapes", "orange", "strawberry", "mango"}

	for i, fruit := range fruits {
		if len(fruit) == 5 {
			fmt.Println("Index ", i, "has number of letters", len(fruit))
		}
	}
	//var i = len(fruits) - 3
	//for i < len(fruits) {
	//	fmt.Println(fruits[i])
	//	i++
	//}
}

/*	fruits := []string{"banana", "apple", "grapes", "orange", "mango"}

	fmt.Println("Range-based for")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	fmt.Println("Range-based for with index")
	for i, fruit := range fruits {
		fmt.Println("Index:", i, " fruit: ", fruit)
	}

	fmt.Println("index-based for")
	var i = 0
	for i < len(fruits) {
		fmt.Println("Index:", i, " fruit: ", fruits[i])
		i++
	}
*/
