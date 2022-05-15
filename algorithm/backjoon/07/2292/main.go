package main

import "fmt"

func main() {
	var room int
	fmt.Scanf("%d", &room)
	fmt.Println(getCount(room))
}

func getCount(room int) int {
	cost := 1
	for idx := 1; idx <= (room - 1); {
		idx += (6 * cost)
		cost++
	}
	return cost
}
