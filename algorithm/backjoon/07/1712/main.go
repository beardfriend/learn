package main

import (
	"fmt"
)

// 고정비용 : [임대로 ,제산세, 보험료, 급여 ...] A만원
// 생산 비용 : [재료, 인건비] B만원 가변

// 노트북 가격 C만원 총수입 > 고정비용 + 가변비용
// 되는 상황을 수식으로 나타내라.

func main() {
	var fixed, produce, income int
	fmt.Scanln(&fixed, &produce, &income)

	if produce >= income {
		n := -1
		fmt.Print(n)
	} else {
		n := fixed/(income-produce) + 1
		fmt.Print(n)
	}
}

func first() {
	var fixed, produce, income int
	fmt.Scanln(&fixed, &produce, &income)

	n := 1
	if produce > income {
		n = -1
		fmt.Print(n)
		return
	}

	for fixed+(produce*n) >= income*n {
		n++
	}
	fmt.Print(n)
}
