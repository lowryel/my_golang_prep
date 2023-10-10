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

type arrayStruct struct {
	key string
	value int
}

type kvPair []arrayStruct
var heights map[string]int


func main() {
	p()	
	areaofa :=areaOfSquare(21, 20)
	fmt.Println(areaofa.instance())
	areaofa.add1(4, 6)
	fmt.Println(*areaofa)
	compareStructs()
	mapMethods()
	structure()
}


func structure() {
	heights := make(map[string]int)
	heights["Janice"] = 174
	heights["James"] = 180
	heights["Jane"] = 172
	fmt.Println(heights)
	p := make(kvPair, len(heights))
	fmt.Println(p)
	h := kvPair{}
	h.Heights()
}

func (kv kvPair) Heights() {
	fmt.Printf("%T\n", kv)
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
func compareStructs() {
	area1 :=area{
		a:8,
		b:5,
	}
	area2 :=area{
		a:8,
		b:3,
	}
	fmt.Println(area1==area2)
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

func (c *area) add1(d, e int){
	c.a += d
	c.b += e
}


// when working with maps, you first have to define and initialze with a variable
// e.g
func mapMethods() {
	heights := make(map[string]int)
	heights["Eugene"] = 174
	heights["Abena"] = 171
	heights["Edem"] = 172
	heights["Thiel"] = 174
	fmt.Println(heights)
	var list []int
	for _, ok := range heights {
		list = append(list, ok)
		// fmt.Println("Height in meters(m) is:", list)
	}
	fmt.Println(list)
}





