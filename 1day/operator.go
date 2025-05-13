package main

import "fmt"

func main() {
	var a int 
	var b int
	// Taking input from user
	fmt.Print("Enter two values:-")
	fmt.Scanf("%d %d",&a,&b)


	// Printing of data
	fmt.Println(a,"greater than",b,":",a>b)
	fmt.Println(a,"less than",b,":",a<b)
	fmt.Println(a,"equals to",b,":",a==b)
	fmt.Println("Addition:-",a+b)
	fmt.Println("subtraction:-",a-b)
	fmt.Println("multiplication:-",a*b)
	fmt.Println("Division:-",a/b)
	fmt.Println("remainder:-",a%b)
	a++
	fmt.Println("increment:-", a)
	b--
	fmt.Println("decrement:-",b)

}