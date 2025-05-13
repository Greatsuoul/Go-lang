package main

import "fmt"

//Global Declaration
var a int = 20

func main() {
	// Local Variable
	s := 10
	d := 6.88
	var b float64 = 3.88

	fmt.Println(s + a)
	fmt.Println(s - a)
	fmt.Println(s * a)
	fmt.Println(s / a)
	fmt.Println(s % a)
	// float operations

	fmt.Println(b + d)
	fmt.Println(b - d)

}
