package main

import "fmt"

func main() {
	var sum int = 1 // 현재위치
    var step int = 1 //지나가는 방의 개수
	var n int // 입력되는 수 , 즉 한계

	fmt.Scanf("%d", &n)
    
	for sum < n { // 현재위치가 한계보다 작을때
		step++
        sum += 6 * (step-1)
	}
	fmt.Print(step)
}
