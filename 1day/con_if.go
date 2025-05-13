package main

import "fmt"

func main() {

	var a, b int = 240, 30
	if a == b {
		fmt.Println("Equal")
	} else if a < b {
		fmt.Println("Less")
	} else {
		fmt.Println("Greater")
	}
}