package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	Solve11653(scanner, writer)
}

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}

func Solve11653(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n == 1 {
		return
	}

	res := factorization(n)
	for _, v := range res {
		_, _ = fmt.Fprintln(writer, v)
	}
}

func factorization(n int) []int {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return append(factorization(i), factorization(n/i)...)
		}
	}
	return []int{n}
}
