/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name”
and “address”, respectively. Your program should use Marshal()
to create a JSON object from the map, and
then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type u struct {
		name    string
		address string
	}

	var user u

	fmt.Println("Enter a name: ")
	fmt.Scan(&user.name)
	fmt.Println("Enter an address: ")
	fmt.Scan(&user.address)

	myMap := make(map[string]string)

	myMap["name"] = user.name
	myMap["address"] = user.address

	barr, err := json.Marshal(myMap)

	fmt.Println(string(barr), err)

}
