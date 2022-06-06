package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&M)
	fmt.Scan(&N)

	number := 10000
	var slice [10001]int
	for i := 2; i < number; i++ {
		slice[i] = i
	}

	for _, i := range slice {
		if slice[i] == 0 {
			continue
		}
		for j := i + i; j <= number; j += i {
			slice[j] = 0
		}
	}
	var result []int

	for i := M; i <= N; i++ {
		if slice[i] == 0 {
			continue
		}
		result = append(result, slice[i])
	}
	if len(result) == 0 {
		fmt.Print("-1")
		return
	}
	var sum int
	for _, val := range result {
		sum += val
	}
	fmt.Println(sum)
	fmt.Println(result[0])
}
