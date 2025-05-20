package main 
import "fmt"
type dar interface{
	m1() float64
	m2() float64
}

type rec struct{
	l float64
	w float64
}

func (s rec) m1() float64{
	return s.l*s.w
}

func (s rec) m2() float64{
	return 2*(s.l + s.w)
}

func main (){
	var s dar
	s = rec{l: 5,w:	4}
	// s=a
	fmt.Println(s)
	fmt.Println(s.m1())
	fmt.Println(s.m2())
	
}