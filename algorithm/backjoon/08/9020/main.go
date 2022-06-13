package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	decimal := make([]bool, 10001)

	for i := 2; i <= 10000; i++ {
		if decimal[i] {
			continue
		}
		for j := i * 2; j <= 10000; j += i {
			decimal[j] = true
		}
	}
	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)
		for k := n / 2; k > 0; k-- {
			if !decimal[k] && !decimal[n-k] {
				fmt.Println(k, n-k)
				break
			}
		}

	}
}
