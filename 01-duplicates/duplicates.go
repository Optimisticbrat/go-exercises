// Eliminate duplicate numbers
package main

import "fmt"

func main() {

	var numbers = []int{5, 2, 2, 7, 6, 2, 3, 3, 4, 5}
	var uniq []int
	for _, num := range numbers {
		var dup = false
		for _, n := range uniq {
			if num == n {
				dup = true
			}
		}
		if dup == false {
			uniq = append(uniq, num)
		}
	}
	fmt.Println("numbers:", numbers)
	fmt.Println("non-duplicate numbers:", uniq)
}

//-----------------------------------------------------------
//Even and odd number write seperately//
/*func main() {
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
}*/
//------------------------------------------------------------
//Prime number program//
/*func isprime(n int) bool {
	prime := true
	i := 2
	for i < n {
		if n%i == 0 {
			prime = false
		}
		i++
	}
	return prime
}
	func main() {
	var numbers = []int{7, 35, 42, 12, 19, 76, 87}
	var primes []int
	var nonprimes []int
	for _, num := range numbers {
		//fmt.Printf("%d is Prime: %t\n", num, isprime(num))
		if isprime(num) == true {
			primes = append(primes, num)
		} else {
			nonprimes = append(nonprimes, num)
		}
	}
	fmt.Println("Primes: ", primes, len(primes))
	fmt.Println("Non-primes:", nonprimes, len(nonprimes))
}*/
