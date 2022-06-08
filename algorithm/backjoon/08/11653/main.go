package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N == 1 {
		return
	}
	if N == 2 {
		fmt.Print(N)
		return
	}
	decimal := make([]int, 2, N)
	for i := 2; i <= N; i++ {
		decimal = append(decimal, i)
	}
	for _, i := range decimal {
		if decimal[i] == 0 {
			continue
		}
		for j := i + i; j <= N; j += i {
			decimal[j] = 0
		}
	}
	var newDecimal []int
	for i := 0; i < len(decimal); i++ {
		if decimal[i] == 0 {
			continue
		}
		newDecimal = append(newDecimal, i)
	}

	j := 0
	for j < len(newDecimal) {
		if N%newDecimal[j] != 0 {
			j++
			continue
		}
		N = N / newDecimal[j]
		fmt.Println(newDecimal[j])
	}
}
