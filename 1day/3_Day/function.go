package main

import "fmt"

func main() {
	//single return type
	d := add(3, 2)
	fmt.Println("Add:-", d)

	// multiple return type
	m, d, r := fun(5, 3)
	fmt.Println("div:-", d, ":: mul:-", m, ":: rem:-", r)


	// variadic func
	fmt.Print(vari(10,20,30,40,50))
}

func fun(a int, b int) (int, int, int) {
	mul := a * b
	div := a / b
	rem := a % b
	return mul, div, rem
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func vari(nu ...int) (sum int){
	for _,val:=range nu{
		sum+=val
	}	
return sum
}