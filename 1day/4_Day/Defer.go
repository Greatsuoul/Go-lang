package main

import "fmt"

func main() {

	defer add(8, 9)

	defer add(8, 56)

	add(44, 56)

}

func add(a int, b int) int {
	sum := a + b
	fmt.Println("Result: ", sum)
	return 0
}
