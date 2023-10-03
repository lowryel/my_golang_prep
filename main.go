package main

import "fmt"

func main(){
	conferenceName := "Go Conference"  // := is applied in assigning values to strings. it declares and assign variables
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to the %v booking platform\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and so far %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Make your booking here")
	// var bookings [50]string //It is not appropriate to use an array due to space redundancy use slice instead
	var bookings []string // an open ended list that takes several arguments

	var firstName string
	var lastName string
	var email string
	var noTickets uint
	fmt.Println("Provide your firstName, lastName, Email and Tickets in that order")
	fmt.Scan(&firstName, &lastName, &email, &noTickets)

	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - noTickets
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Printf("The whole slice %v\n", bookings)
	fmt.Printf("First item in slice ==> %v\n", bookings[0])
	if (len(bookings)>0){
		fmt.Println("Hello, there are items in the array")
	}
	

	fmt.Printf("Thank you %v %v, you booked %v tickets. You will receive a confirmation email soon on %v\n", firstName, lastName, noTickets, email)
}
