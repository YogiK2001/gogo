package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := make([]int, 0, 3)
	var quit string

	for {
		var num int
		fmt.Println("Enter a number: ")
		fmt.Scan(&num)

		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("Sorted slice: ")
		for i := range slice {
			fmt.Println(slice[i])
		}
		fmt.Println("Enter 'q' to quit or any other key to continue: ")
		fmt.Scan(&quit)
		if quit == "X" {
			break
		}
	}
}
