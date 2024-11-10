package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	x := "Joe"
	fmt.Println(quote.Go())
	fmt.Printf("Hello %s\n", x) // Format String
	fmt.Println("Hello" + x)    // Concatenation
}
