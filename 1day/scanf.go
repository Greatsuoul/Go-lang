package main

import "fmt"

func main() {
	// Taking input from user
	var age int
	fmt.Println("Enter the age:-")
	count, err := fmt.Scanf("%d", &age)

	// Printing of data
	fmt.Println("count =", count, ", err =", err)
	fmt.Printf("%d years old\n", age)

}
