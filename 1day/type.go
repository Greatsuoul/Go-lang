package main

import (
	"fmt"
	"reflect"
)
func main() {
	// Taking input from user
	var age int = 20 
	name := "Prasad"
	ff:=12.33
	c:=true
	// using %T method
	fmt.Printf("data:%v type:%T\n",age,age)
	fmt.Printf("data:%v type:%T\n",name,name)
	fmt.Printf("data:%v type:%T\n",ff,ff)
	fmt.Printf("data:%v type:%T\n",c,c)
	fmt.Println()

	// using reflect.Typeof method
	
	fmt.Printf("data:%v type:%v\n",age,reflect.TypeOf(age))
	fmt.Printf("data:%v type:%v\n",name,reflect.TypeOf(name))
	fmt.Printf("data:%v type:%v\n",ff,reflect.TypeOf(ff))
	fmt.Printf("data:%v type:%v\n",c,reflect.TypeOf(c))

}
