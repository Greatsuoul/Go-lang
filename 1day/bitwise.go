package main

import "fmt"

func main() {
	// Taking input from user
	var a,b int = 20,12
	
	// Printing of data
	buff := a&b
	fmt.Println("AND is ", buff)

	buff = a|b
	fmt.Println("OR is ", buff)

	buff = a^b
	fmt.Println("XOR is ", buff)

	buff = a<<3
	fmt.Println("3 left shift is ", buff)

	buff = a>>3
	fmt.Println("3 left shift is ", buff)
}