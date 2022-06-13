package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(pibonachi(N))
}

func pibonachi(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return pibonachi(N-1) + pibonachi(N-2)
}
