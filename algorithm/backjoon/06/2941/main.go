package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	maps := map[string]int{
		"c=":  1,
		"c-":  1,
		"dz=": 1,
		"d-":  1,
		"lj":  1,
		"nj":  1,
		"s=":  1,
		"z=":  1,
	}
	count := 0
	for i := 0; i < len(input); i++ {
		temp := ""
		if str := input[i]; string(str) == "d" {
			if i < len(input)-2 {
				temp = input[i : i+3]

				if maps[temp] == 1 {

					i += 2
					count++
					continue
				}
			}
			if i < len(input)-1 {
				temp = input[i : i+2]

				if maps[temp] == 1 {

					i += 1
					count++
					continue
				}
			}
			count++
		} else {
			if i < len(input)-1 {
				temp = input[i : i+2]
				if maps[temp] == 1 {
					i += 1
					count++
					continue
				}
			}
			count++
		}

	}
	fmt.Print(count)
}
