package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Hello!")
	fmt.Println("Welcome to", conferenceName, "booking application")

	fmt.Println("The number of tickets available is:", remainingTickets)
	fmt.Println("The number of tickets left is:", conferenceTickets)
}
