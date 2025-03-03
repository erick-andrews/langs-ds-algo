package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Hello!")
	fmt.Printf("Welcome to %v booking application \n", conferenceName)

	fmt.Printf("The number of tickets available is: %v and %v is the total left. \n", remainingTickets, conferenceTickets)
	// fmt.Println("The number of tickets left is:", conferenceTickets)
}
