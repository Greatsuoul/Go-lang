package main

import "fmt"

func main() {
	var fruit string

	fmt.Print("Enter name of fruit: ")
	fmt.Scanln(&fruit)

	// x := [5]string{"Apple", "Pineapple", "Banana", "Pear", "Vadapav"}

	switch fruit {
	case "Apple", "Pineapple","Fig","Watermelon":
		fmt.Println("Tropical fruit")
	case "Brinjal","Cabbage":
		fmt.Println("Not a fruit")
	default:
		fmt.Println("Sorry,not recognized")
	}
}
