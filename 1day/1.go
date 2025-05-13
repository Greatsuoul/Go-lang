package main

import "fmt"

func main() {
	// Implicit declaration
	name := "Prasad"
	s := 10
	d := 6.88

	// Explicit declaration
	var sname string = "Kanade"
	var a int = 20
	var b float64 = 3.88

	// Taking input from user
	var age int
	fmt.Println("Enter the age:-")
	fmt.Scanf("%d", &age)

	// Printing of data

	fmt.Printf("My name is %s\n Surname is %s\n %d years old\n", name, sname, age)

	//Printing values with some mathematical operation
	fmt.Println(s + a)
	fmt.Println(s - a)
	fmt.Println(s * a)
	fmt.Println(s / a)
	fmt.Println(s % a)
	// float operations

	fmt.Println(b + d)
	fmt.Println(b / d)

}
