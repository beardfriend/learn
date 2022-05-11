package main

import "fmt"

// 처음 나오고 이후에 또 나오면 에러

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		var word string
		fmt.Scan(&word)
		letters := make(map[uint8]bool)
		var prev uint8
		isGroup := true
		for j := 0; j < len(word); j++ {
			_, exist := letters[word[j]]
			if !exist {
				letters[word[j]] = true
			} else if prev != word[j] {
				isGroup = false
				break
			}
			prev = word[j]
		}
		if isGroup {
			count++
		}
	}
	fmt.Print(count)
}
