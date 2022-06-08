package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	// var k, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)
	find(n)
}

func find(n int) {
	i := 2
	for n != 1 {
		if n%i == 0 {
			n /= i
			fmt.Println(i)
		} else {
			i += 1
		}
	}
}
