package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Println("Hello!")
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	//Let's see the types of the above vars:
	fmt.Printf("conferenceTickets is %T, conferenceName is %T \n", conferenceTickets, conferenceName)

	askUserNameTicketsNum(remainingTickets)

	//fmt.Printf("The number of tickets available is: %v and %v is the total left. \n", remainingTickets, conferenceTickets)
	// fmt.Println("The number of tickets left is:", conferenceTickets)
}

// Data types supported:
// Strings, Integers, Float, Booleans, Maps, Arrays. Go can infer the type of a var.

// ints: uint8-64, and int8-64. uint prohibits neg, int is neg to pos.

func askUserNameTicketsNum(remainingTickets uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// need a pointer! it's a ... special variable!
	// fmt.Println(&userTickets) --> allow printing place in memory.
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your # of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)

	fmt.Printf("Remaining tickets: %v", remainingTickets)
}
