package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)
	var stringarray []string

	
	for i:= 0; i < len(input); i++ {
		stringarray = append(stringarray, strconv.Itoa(i))
	}
	fmt.Println(stringarray)
}