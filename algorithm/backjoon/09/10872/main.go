package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	result := factorial(N)
	fmt.Print(result)
}

func factorial(N int) int {
	if N == 0 {
		return 1
	}
	return N * factorial(N-1)
}
