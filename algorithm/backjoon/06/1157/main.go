package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	var lowerstr string
	reader:= bufio.NewReader(os.Stdin)
	fmt.Fscan(reader,&lowerstr)
	str:= strings.ToUpper(lowerstr)
	  
	var dataSet = make(map[string]int)
	for i:=0; i<len(str); i++ {
		str := fmt.Sprintf("%c",str[i])
		if dataSet[str] == 0 {
			dataSet[str] = 1
		} else {
			dataSet[str] += 1
		}
		
	}

	winner :=0
	var laststring string
	for v,_:= range dataSet {
		if winner < dataSet[v] {
			winner = dataSet[v]
			laststring = v
		} 
	}

	count := 0
	for v,_:= range dataSet {
		if dataSet[v] == winner {
			count = count+1
		}

	}

	if count > 1 {
		fmt.Print("?")
		return
	}

	fmt.Print(laststring)
	
}