//-----------------------------------------------------------
//Even and odd number write seperately//
package main

import "fmt"

func main() {
	var numbers = []int{7, 35, 42, 12, 19, 76, 87}
	var odd []int
	var even []int
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	fmt.Println("even:", even, len(even))
	fmt.Println("odd:", odd, len(odd))
}
