package main

import "fmt"

func main() {
	//am := []int{19, 23, 34, 56, 38}
	in := map[int]string{
		1:  "Prasad",
		23: "Kanade",
		54: "Somnath",
	}

	// while data is available
	value1, found1 := in[23]
	fmt.Println("While Data and Key both are available")
	fmt.Println("Data is:- ", value1, " :: key status:- ", found1)

	// when Data isnt available
	value2, found2 := in[5]
	if found2 {
		fmt.Println("Data is:- ", value2, "  ::key status:-", found1)
	} else {
		fmt.Println("No data available")
	}

	// map after deleting 1 key-value pair
	delete(in,54)
	fmt.Println(in)

	// deleting Whole 
	for key ,_:= range in {
		delete(in,key)
	}
	fmt.Println(in)
	

}
