package main

import "fmt"

func main() {
	var count int

	fmt.Scanf("%d", &count)
	arr := make([]int, 0, 1000)
	for i := 0; i < count; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}

	var temp int
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	for i := 0; i < count; i++ {
		fmt.Println(arr[i])
	}
}
