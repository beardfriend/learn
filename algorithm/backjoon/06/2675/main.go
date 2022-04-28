// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main () {
// 	var s int
// 	reader:= bufio.NewReader(os.Stdin)

// 	fmt.Fscanln(reader,&s)
// 	type data struct {
// 		TotoalCount int
// 		StringArray []string
// 	}
// 	var allData  []data
// 	for i :=0; i < s; i++ {
// 		datas := data{}
// 		var stringData string
// 		fmt.Scanln(&datas.TotoalCount,&stringData)
// 		for _,v := range []rune(stringData) {
// 			datas.StringArray =append(datas.StringArray,fmt.Sprintf("%c",v))
// 		}
// 		allData = append(allData, datas)
// 	}

// 	for i := 0 ;i < len(allData); i++ {
// 		for j:=0;  j < len(allData[i].StringArray); j++ {
// 			y:= strings.Repeat(allData[i].StringArray[j],allData[i].TotoalCount)
// 			fmt.Print(y)
// 		}
// 		fmt.Println()
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, r int
	var str string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &t)
	defer writer.Flush()

	for i:=0; i<t; i++ {
		fmt.Fscanf(reader, "%d %s\n", &r, &str)
		for j:=0; j<len(str); j++ {
			for k:=0; k<r; k++ {
				fmt.Fprint(writer, string(str[j]))
			}
		}
		fmt.Fprint(writer, "\n")
	}
}