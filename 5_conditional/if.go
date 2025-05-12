package main

import "fmt"

func main() {
	name := "Alice"  //Implicit Declaration
	var age int = 20 //Explicit Declaration
	x := 20
	
	if age<x{
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}else if age>x{
		fmt.Printf("not in data")
	}else{
		fmt.Printf("Get out")
	}
}