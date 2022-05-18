// 분수찾기

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	i := 1
	for num > i {
		num -= i
		i++
	}

	if i%2 == 1 {
		fmt.Printf("%d/%d", i+1-num, num)
	} else {
		fmt.Printf("%d/%d", num, i+1-num)
	}
}

// 1 2 3 4 5 6
