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
	var bookings []string
	for {
		remainingTickets = askUserNameTicketsNum(&bookings, remainingTickets)
	}

	//fmt.Printf("The number of tickets available is: %v and %v is the total left. \n", remainingTickets, conferenceTickets)
	fmt.Println("The number of tickets left is:", remainingTickets)
}

// Data types supported:
// Strings, Integers, Float, Booleans, Maps, Arrays. Go can infer the type of a var.

// ints: uint8-64, and int8-64. uint prohibits neg, int is neg to pos.

func askUserNameTicketsNum(bookings *[]string, remainingTickets uint) uint {
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

	createArrays(bookings, firstName, lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your # of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)

	// fmt.Printf("Remaining tickets: %v", remainingTickets)
	return remainingTickets // Return the updated value
}

func createArrays(bookings *[]string, firstName string, lastName string) {
	// Arrays in Go are of a fixed size
	// What kind of value do we want to store here? A list of names of users who booked tickets!
	// Can be empty or already have things in it.
	// No mixed types allowed in an array.

	// var bookings = [50]string{"Erick", "Tom", "Pete"}
	// var bookings [50]string

	// But what about dynamic size list? A slice - array type under hood, but is dynamic in size.
	*bookings = append(*bookings, firstName+" "+lastName)
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(*bookings))

	fmt.Printf("These are all our bookings: %v\n", bookings)

}
