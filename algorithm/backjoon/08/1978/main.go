package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, count int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a)

		var divisorCount int
		for j := 0; j < a; j++ {
			if a%(j+1) == 0 {
				divisorCount++
			}
		}
		if divisorCount == 2 {
			count++
		}
	}
	fmt.Println(count)
}

func areatostnes() {
	number := 1000
	var prime [1001]int

	for i := 2; i < number; i++ {
		prime[i] = i
	}
	for i := 2; i <= number; i++ {
		if prime[i] == 0 {
			continue
		}

		for j := i + i; j <= number; j += i {
			prime[j] = 0
		}

	}
	fmt.Print(prime)
}
