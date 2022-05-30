package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var count int

	for {
		if N%5 != 0 {
			if N < 3 {
				count = -1
				break
			}
			N = N - 3
			count++
		} else {
			break
		}
	}
	if count != -1 {
		count += N / 5
	}
	fmt.Println(count)
}

// 5로 나눠질 때까지 3 빼는 걸 반복
