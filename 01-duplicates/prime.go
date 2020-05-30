//------------------------------------------------------------
//Prime number program//
package main

import "fmt"

func isprime(n int) bool {
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
}
