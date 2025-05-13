package main

import "fmt"

func main() {
	var am [5]int = [5]int{19, 23, 34, 56, 38}
	var name [5]string = [5]string{"Anand", "vishnu", "jagdish", "hare", "ugram"}
	sum := [...]int{123, 298, 34, 234, 94, 23}
	fmt.Println(am)
	fmt.Println(name)
	// for i:=0 ; i<5;i++{
	fmt.Println(am[3])
	// }
	fmt.Println(sum)
	fmt.Println(len(sum))
	for index, values := range sum {
		fmt.Println(index, "=>", values)
	}

}
