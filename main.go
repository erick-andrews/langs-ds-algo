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

	askUserNameTicketsNum()

	fmt.Printf("The number of tickets available is: %v and %v is the total left. \n", remainingTickets, conferenceTickets)
	// fmt.Println("The number of tickets left is:", conferenceTickets)
}

// Data types supported:
// Strings, Integers, Float, Booleans, Maps, Arrays. Go can infer the type of a var.

// ints: uint8-64, and int8-64. uint prohibits neg, int is neg to pos.

func askUserNameTicketsNum() {
	var userName string
	var userTickets uint

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
