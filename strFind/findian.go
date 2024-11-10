package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	var flag bool
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	input = strings.ToLower(input)
	for i := 0; i < len(input); i++ {
		if input[0] == 'i' || input[i] == 'a' || input[len(input)-1] == 'n' {
			flag = true
		} else {
			flag = false
		}
	}
	if !flag {
		fmt.Println("Not Found!")
	} else {
		fmt.Println("Found")
	}
}
