package main

import (
	"fmt"
	"strings"
	"time"
)

var bookings = make([]UserData, 0) // an open ended list that takes several arguments
type UserData struct{
	firstName	string
	lastName	string
	email 		string
	noTickets 	uint
}

func main(){
	var conferenceTicket int= 50
	var remainingTickets uint= 50
	conferenceName := "Go Conference"  // := is applied in assigning values to strings. it declares and assign variable
	
	welcomeMessage(conferenceName, conferenceTicket, remainingTickets)
	// var bookings [50]string //It is not appropriate to use an array due to space redundancy. use slice instead

	for {

		var firstName string
		var lastName string
		var email string
		var noTickets uint



		fmt.Println("Provide your firstName")
		fmt.Scan(&firstName)

		fmt.Println("Provide your lastName")
		fmt.Scan(&lastName)

		fmt.Println("Provide your Email")
		fmt.Scan(&email)

		fmt.Println("Provide your Tickets")
		fmt.Scan(&noTickets)

		userData := UserData{
			firstName: firstName,
			lastName: lastName,
			email: email,
			noTickets: noTickets,
		}

		isValidEmail :=strings.Contains(email, "@")
		validTicketsNo := noTickets > remainingTickets || noTickets <=0
		if  !isValidEmail || validTicketsNo  || remainingTickets <=0{
			fmt.Println("Invalid inputs, try again")
			continue

		} else {
			// Order summary and email to viewers
			go sendEmailToViewer(userData.firstName, userData.lastName, userData.email, userData.noTickets)

			bookings = append(bookings, userData)
			remainingTickets = remainingTickets - noTickets
			fmt.Printf("%v tickets remaining\n", remainingTickets)
			fmt.Printf("List of bookings %v\n", bookings)
			// returning viewers first names
			firstNames := returnFirstName()
			fmt.Printf("Returning first name of viewers %v\n", firstNames)
		}
	}
}

func sendEmailToViewer(fN, lN, email string, NoTickets uint) {

	fmt.Printf("Thank you %v %v, you booked %v tickets. You will receive a confirmation email soon on %v\n", fN, lN, NoTickets, email)
	time.Sleep(5*time.Second)
	fmt.Println("############################################")
	Email(email,fN, lN, NoTickets)
	fmt.Println("############################################")
}


func returnFirstName() []string {
		firstNames := []string{}
		for _, booking := range bookings{
			firstNames= append(firstNames, booking.firstName)
	}
	return firstNames
}


func Email(email, firstName, lastName string, noTickets uint) {
	fmt.Printf("Email sent to %v:\nContent: %v tickets to %v %v\n", email, noTickets, firstName, lastName)

}

func welcomeMessage(conferenceName string, conferenceTicket int, remainingTickets uint) {
	fmt.Printf("Welcome to the %v booking platform\n", conferenceName)

	fmt.Printf("We have a total of %v tickets and so far %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Make your booking here")
}
