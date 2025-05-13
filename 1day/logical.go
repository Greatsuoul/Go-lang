package main

import "fmt"

func main() {
	var a int
	var b int
	x := 25
	// Taking input from user
	fmt.Print("Enter two values:-")
	fmt.Scanf("%d %d", &a, &b)

	// Printing of data
	fmt.Println((x>b)&&(x<a))
	fmt.Println((x>a)||(x<b))
	fmt.Println(!(x<a))

}
