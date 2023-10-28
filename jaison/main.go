package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	firstname string
	lastname string
}

func JSON(){
	var person Person
	jsonString := `{
		"firstname": "Eugene",
		"lastname": "Lowry",
	}`

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(person.firstname, person.lastname)
}


func main() {
	JSON()
}