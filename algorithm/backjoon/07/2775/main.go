package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	var result int
	for i := 0; i < count; i += 1 {
		var k, n int
		fmt.Scan(&k)
		fmt.Scan(&n)
		result = getCount(k, n)
		fmt.Println(result)
	}

	ct := getCount(2, 2)
	fmt.Println(ct)
}

func getCount(k int, n int) int {
	if n == 0 {
		return 0
	} else if k == 0 {
		return n
	} else {
		return getCount(k, n-1) + getCount(k-1, n)
	}
}

// 거주조건 :
// a층 b호에 살려면
// a-1층의 1호부터 b호까지 사람들의 수의 합만큼 사람을 데려와야 함
// 비어있는 집 x
// 모든 거주민이 계약조건 지켰다.
// k층 n호 몇명 살고 있는지?

// 2 3
// 3층 1호:1명 2호:5명 3호:15명 4호: 35명
// 2층 1호:1명 2호:4명 3호:10명  4호: 20명
// 1층 1호:1명  2호:3명 3호:6명 4호: 10명
// 0층 1호:1명 2호:2명 3호:3명   4호: 4명

// 1 + 2 + 3

// 1 3 6 = 1 4 10

// n - 1 + n -2 + ... n - n = last
// 1 4

// 2 4

// 1 + 2 + 3 + 4 = 10
//

// 2 3 4 , 3 6 10, 4 10 20 6 10 15
