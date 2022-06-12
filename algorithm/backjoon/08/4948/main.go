package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	stdio = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	N     = 123456
)

func main() {
	defer stdio.Flush()

	sieve := make([]bool, 2*N+1) // false: prime, true: composite
	sieve[1] = true

	for i := 2; i*i <= 2*N; i++ {
		if sieve[i] {
			continue
		}
		for j := 2 * i; j <= 2*N; j += i {
			sieve[j] = true
		}
	}

	for {
		var n int
		fmt.Fscan(stdio, &n)
		if n == 0 {
			break
		}

		cnt := 0
		for i := n + 1; i <= 2*n; i++ {
			if !sieve[i] {
				cnt++
			}
		}
		fmt.Fprintln(stdio, cnt)
	}
}
