package main

import "fmt"

func main() {
	var a, b, buff int

	// Taking input from user
	fmt.Print("Enter two values:-")
	fmt.Scanf("%d %d", &a, &b)

	buff = a
	a += b
	fmt.Println("addition is ", a)

	a = buff
	a -= b
	fmt.Println("subtraction is ", a)

	a = buff
	a *= b
	fmt.Println("multiplication is ", a)

	a = buff
	a /= b
	fmt.Println("division is ", a)

	a=buff
	a%=b
	fmt.Println("remainder is ", a)


}
