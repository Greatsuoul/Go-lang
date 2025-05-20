package main

import "fmt"

func main() {
	i := 77
	fmt.Printf("%T %v\n", &i, &i)
	fmt.Printf("%T %v\n", *(&i), *(&i))


	var ii *int=&i
	fmt.Print(*ii)
	fmt.Println(ii)
	
}
