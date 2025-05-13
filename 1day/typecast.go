package main

import (
	"fmt"
	"strconv" // to use Itoa and Atoi
)
func main() {
	// convetrting int to float 
	var age int = 30
	var ff float64 = float64(age)
	fmt.Printf("%.4f\n",ff)

	// converting float to int 
	var as float64 = 12.34223
	var i int = int(as)
	fmt.Printf("%d\n",i)

	// conversion of Integer to string
	var s int = 12
	var t string = strconv.Itoa(s)
	fmt.Printf("%q\n",t)

	// converting string into an integer 
	var m string = "90"
	n,err := strconv.Atoi(m)
	fmt.Printf("%d\n",n)
	fmt.Print(err)


}
