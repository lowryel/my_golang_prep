package main

import (
	"fmt"
)

type dob struct {
	day int
	month int
	year int
}

type people struct{
	name string
	email string
	dob dob
}
// assign members to a map
var members map[int]people

func (c *people) NewPeople(name string, email string, db *dob) {
	newpeople := make(map[int]people)
	// newdob :=make(map[string]int)
	db.day= 2
	db.month = 5
	db.year = 1993
	newpeople[1] = people{
		name: name,
		email: email,
		dob: dob{db.day,db.month, db.year},
	}
	for _, v := range newpeople{
		db := fmt.Sprintln(v)
		fmt.Println(db)
	}
	fmt.Println(newpeople)
	// newpeople[1]= people{
	// 	name: c.name,
	// 	email: c.email,
	// 	dob: dob{
	// 		day: 25,
	// 		month: 4,
	// 		year: 1995,
	// 	},
	// }
}

func main() {
	// initialize members map with make keyword
	members = make(map[int]people)
	members[1] = people{
		name: "Eugene",
		email: "ellowry@gmail.com",
		dob: dob{
			day: 27,
			month: 3,
			year: 1993},
	}
	members[2] = people{
		name: "Lowry",
		email: "low@gmail.com",
		dob: dob{
			day: 22,
			month: 9,
			year: 1993},
	}
	members[3] = people{
		name: "Atsu",
		email: "at&t@gmail.com",
		dob: dob{
			day: 1,
			month: 1,
			year: 1999},
	}
	members[4] = people{
		name: "Thiel",
		email: "thiel@gmail.com",
		dob: dob{
			day: 27,
			month: 3,
			year: 2003},
	}
	for k, v := range members{
		fmt.Println(k, v.name, v.email, v.dob)
	}
	k :=people{}
	k.NewPeople("Eugnee", "eugene@gmail.com", &dob{})
}

// func (kv kvPair) Heights() {
// 	fmt.Printf("%T", kv)
// }

