package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLargestPalindromicProduct())
}

func findLargestPalindromicProduct() int {
	product := 0
	largestProduct := 0

	// smallest 3-digit palindromic product is 101 * 101 = 10201
	for i := 999; i > 101; i-- {
		for j := 999; j > 101; j-- {
			product = i * j

			if isPalindrome(product) && product > largestProduct {
				largestProduct = product
			}
		}
	}

	return largestProduct
}

// check if a given number is a palindrome
func isPalindrome(num int) bool {
	return num == reverse(num)
}

// reverse a given number (1234 -> 4321)
func reverse(in int) int {
	reversed := 0

	for i := 0; in > 0; i++ {
		remainder := in % 10
		in = (in - remainder) / 10
		reversed = reversed*10 + remainder
	}

	return reversed
}
