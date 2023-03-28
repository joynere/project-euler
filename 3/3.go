package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getLargestPrimeFactor(600851475143))
}

func getLargestPrimeFactor(num int) int {
	largestFactor := 0
	for i := 2; i < num; i = getNextPrime(i) {

		if num%i == 0 {
			largestFactor = i
			quot := num / i

			// if quotient is prime, finish and return largest prime factor
			if isPrime(quot) {
				if quot > largestFactor {
					return quot
				} else {
					return largestFactor
				}
				// if quotient is not prime, continue factorization with quotient
			} else {
				deep := getLargestPrimeFactor(quot)
				if deep > largestFactor {
					return deep
				} else {
					return largestFactor
				}
			}
		}
	}
	return 0
}

// Get next prime number > given number
func getNextPrime(num int) int {
	for i := num + 1; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(num int) bool {
	// only check for factors up to square root of number
	upperBound := int(math.Sqrt(float64(num)))

	for i := 2; i <= upperBound; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
