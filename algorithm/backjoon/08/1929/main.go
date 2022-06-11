package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	a := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if a[i] {
			continue
		}
		if i >= m {
			fmt.Println(i)
		}
		for j := i * 2; j <= n; j += i {
			a[j] = true
		}
	}
}
