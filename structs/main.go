package main

import (
	"fmt"
)


type Gender struct{
	A string
	B string
}

type person struct {
	name string
	age int
	gender Gender
}

type points struct{
	x float32
	y float32
	z float32
}


func main() {
	p()	
	areaofa :=areaOfSquare(21, 20)
	fmt.Println(areaofa.instance())
}

func p() *person {
	// new_p := make([]int, 5)
	per := person{"Eugene", 25, Gender{A:"Male", B: "Female"}}
	fmt.Println(per.name, per.age, per.gender.A)
	per1 := &per
	fmt.Println(per1.name)
	per1.name = "Thomas"
	fmt.Println(per1)
	fmt.Println(per)

	// p := points{
	// 	x: 8,
	// 	y: 6,
	// 	z: 9,
	// }


	call_points(6,5,8)

	// fmt.Println(p)
	fmt.Println("#####################################################")
	// for k, v := range new_p {
	// 	fmt.Println(k,v)
	// }

	// for i := 0; i<=len(new_p); i++ {
	// 	fmt.Println(i)
	// }
	return nil
}

func call_points(x, y, z float32) *points{
	p:=points{x:x, y:y, z:z}
	fmt.Printf("%v\n", p)
	p1 :=&p
	p1.x=20.254
	fmt.Println(p)
	fmt.Println(p1)
	return nil
}



type area struct{
	a int
	b int
}

func (p area) instance() int{
	square:=p.a * p.b
	return square
}

func areaOfSquare(a, b int) *area{
	area_of_square :=area{ a:a, b:b }
	fmt.Println(area_of_square)
	return &area_of_square
}
