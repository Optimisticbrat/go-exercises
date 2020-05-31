package main

import "fmt"

func main() {
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	ctypes := []string{"Heart", "Spade", "Diamond", "Clubs"}
	fmt.Println(numbers)
	fmt.Println(ctypes)

	for _, ctype := range ctypes {
		for _, number := range numbers {
			fmt.Println(" ", ctype, "of ", number)
		}
	}

}

//for i, fruit := range fruits {
//if len(fruit) == 5 {
//	fmt.Println("Index ", i, "has number of letters", len(fruit))
//}
//}
