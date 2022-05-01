package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	words := strings.Split(input, " ")
	var count int
	for i:= range words {
		if words[i] != "\n" && words[i] != "" {
			count++
		}
	}
	fmt.Println(count)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )


// func main () {

// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	text:= scanner.Text()

// 	slice := strings.Split(text," ")
// 	count := 0
// 	for _, v := range slice {

// 		  if v != "" {
// 			count++
// 		  }	
// 	}
// 	fmt.Println(count)


// }