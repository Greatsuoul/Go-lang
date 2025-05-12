package main

import "fmt"

func main() {
	w := func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("bool")
		case float32,float64:
			fmt.Println("float")
		}
	}
	w(3.5)
	w("PK")
	w(2)
	w(true)
}
