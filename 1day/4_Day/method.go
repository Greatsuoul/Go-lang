package main
import "fmt"

type car struct{
	name string
	// mn int
	yr int 
	// fbf bool
}

func (s car) bb(){
fmt.Println(s.name," made in year",s.yr)
}

func main(){
	// Suzuki:= car{name: "Suzuki",mn: 232, yr :2014,fbf: false}
	// fmt.Println("Name of the car:",Suzuki.name,"\n Model number is:",Suzuki.mn,"\n Year of manufacturing:",Suzuki.yr,"\n Is it 4x4 Vehicle",Suzuki.fbf)
	// s2 := car{name: "Thar",mn:11,yr: 2018,fbf: true}
	s1:= car{name: "Bolero",yr: 2024}
	s1.bb()
	// if s2.name != Suzuki.name{
	// 	fmt.Println("Not Same")
	// }else{
	// 	fmt.Println("Same")
	// }

}