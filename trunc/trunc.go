/*
Write a program that prompts the user to enter a floating-point number and prints the integer, which is a truncated version of the floating-point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&input)
	fmt.Printf("Truncated integer: %d\n", int(math.Trunc(input)))
}
