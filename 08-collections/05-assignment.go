/*
	Write a function that returns all the prime numbers beween the given start and end

print the generated prime numbers in the main function one by one
*/
package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	for _, primeNo := range primes {
		fmt.Println(primeNo)
	}
}

func genPrimes(start, end int) []int {
	primes := []int{}
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
