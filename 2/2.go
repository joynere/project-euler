package main

import "fmt"

// sum even fibonacci numbers under 4mio
func sumEvenFibbonacci(x, y, sum int) int {
	if y >= 4e6 {
		return sum
	} else if y%2 == 0 {
		sum += y
	}

	return sumEvenFibbonacci(y, x+y, sum)
}

func main() {
	fmt.Println(sumEvenFibbonacci(1, 2, 0))
}
