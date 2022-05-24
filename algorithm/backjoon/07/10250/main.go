package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {

		var H, W, N int
		fmt.Scanln(&H, &W, &N)
		floor := N % H
		ho := N/H + 1
		if N%H == 0 {
			floor = H
			ho = N / H
		}
		fmt.Println(floor*100 + ho)
	}
}
