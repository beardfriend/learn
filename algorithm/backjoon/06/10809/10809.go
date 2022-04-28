package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &s)

	var alphabet = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var inputArray []string
	for _,v:= range []rune(s) {
		inputArray = append(inputArray, fmt.Sprintf("%c",v))	
	}
	for i := 0; i < len(alphabet); i++ {
		bol := false
			for j:=0; j < len(inputArray); j++ {

				if inputArray[j] == alphabet[i] {
					bol = true
					fmt.Printf("%d ", j)
					break
				} 

			}
			if !bol {
				fmt.Print("-1 ")
			}

	}

}
