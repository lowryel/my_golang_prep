package main

import (
	// "errors"
	"fmt"

)

type SoundMaker interface {
	MakeSound()
}

type Dog struct {
	sound string
	name string
	color string
}

type Cat struct {
	sound string
	name string
}

func (d Dog) MakeSound() string{
	// Dsound := "woo woo"
	fmt.Printf("Dog Sound: %v, Dog Name: %v, Dog Sound: %v\n\n", d.sound, d.name, d.sound)
	return ""
}

func (c Cat) MakeSound() string{
	// Dsound := "woo woo"
	fmt.Printf("Cat Sound: %v, Cat Name: %v\n", c.sound, c.name)
	return ""
}

type Error struct{
	created string
	ok string
	not_found string
	redirect string
}



func main(){
	error := Error{

	}
	error.created = "201"
	error.ok = "200"
	error.not_found = "400"
	error.redirect = "302"

	dogs := Dog{
		name: "Gameli",
		sound: "Wooo Wo, Woowo",
		color: "cool white and black",
	}
	cats := Cat{
		name: "Peace",
		sound: "Meow, Meow",
	}
	if dogs.name == "Gameli"{
		fmt.Println(error.ok)
	}
	fmt.Println(dogs, cats)
	fmt.Println(dogs.MakeSound(), cats.MakeSound())
}


