package main


import (
	"fmt"
	"log"

	"ex/greetings"
	"loop.go/loop"
)

type Avoid struct {
	name int
	age int
}


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	computation := greetings.Sum(5897,5)
	fmt.Println(computation)
	avoid(47,15)
	call_func()
	fibbo()
	loopl := loop.Loops()
	fmt.Println(loopl)
}

func avoid(a, b int) int {
	var r Avoid
	r.age = a
	r.name = b
	fmt.Printf("%d\n", r.name + r.age)
	return 1
}

func call_func() {
	var input, addvar int
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%d", &input)
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%d", &addvar)
	fmt.Printf("The sum of input and addvar: %d\n", input + addvar)
}

func fibbo () int {
	max := 100
	a := 0
	b := 1
	for ;b <= max; {
		fmt.Println(b)
		a, b = b, a+b
	}
	return b
}

