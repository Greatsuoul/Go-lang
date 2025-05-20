package main

import "fmt"

func main() {
	am := []int{19, 23, 34, 56, 38}
	var name [5]string = [5]string{"Anand", "vishnu", "jagdish", "hare", "ugram"}
	sum := [...]int{123, 298, 34, 234, 94, 23}

	// slicing array
	ns := name[1:5]
	ns1 := sum[1:4]

	fmt.Println(ns)
	for index, values := range ns {
		fmt.Println(index, "=>", values)
	}

	fmt.Println(len(ns))
	fmt.Println(cap(ns))

	sll := make([]int, 5, 10)
	fmt.Println(sll)

	//Adding values to Array
	sll = append(sll, 10, 20, 30)
	fmt.Println(sll)
	fmt.Println(len(sll))
	fmt.Println(cap(sll))

	//combining two arrays
	ns2 := append(sll, ns1...)
	fmt.Println(ns2)
	fmt.Println(len(ns2))
	fmt.Println(cap(ns2))

	//deleting specific emlement
	i := 2
	ni1 := am[:i]
	ni2 := am[i+1:]
	nio := append(ni1, ni2...)
	fmt.Println(ni1)
	fmt.Println(ni2)
	fmt.Println(nio)

	//Copy function
	fun := make([]int, 3)
	numm := copy(fun, am)
	fmt.Println(numm)
	fmt.Println(fun)
	fmt.Println(nio)

}
