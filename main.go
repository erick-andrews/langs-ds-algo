package main

import (
	"fmt"
	"strings"
)

const conferenceTickets uint = 50 // variable scope is package wide
var conferenceName = "Go Conference"

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
	// isOptedInForNewsletter bool, etc.
}

func main() {
	// const ConferenceTickets uint = 50 Global variable scope... (exported)
	var remainingTickets uint = 50

	fmt.Println("Hello!")
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	//Let's see the types of the above vars:
	fmt.Printf("conferenceTickets is %T, conferenceName is %T \n", conferenceTickets, conferenceName)

	// empty slice of maps
	var bookings = make([]UserData, 0)
	// working with structs

	for {
		// NoTicketsRemaining := remainingTickets == 0 if referencing
		// this statement more than once, can be saved as var
		if remainingTickets == 0 {
			// end program
			fmt.Println("Oops, our conference is totally booked. Come back next year.")
			break
		}
		var shouldExit bool

		remainingTickets, shouldExit = askUserNameTicketsNum(&bookings, remainingTickets)

		if shouldExit {
			fmt.Println("Exiting program...")
			break
		}
	}

	//fmt.Printf("The number of tickets available is: %v and %v is the total left. \n", remainingTickets, conferenceTickets)
	fmt.Println("The number of tickets left is:", remainingTickets)
}

// Data types supported:
// Strings, Integers, Float, Booleans, Maps, Arrays. Go can infer the type of a var.

// ints: uint8-64, and int8-64. uint prohibits neg, int is neg to pos.

func askUserNameTicketsNum(bookings *[]UserData, remainingTickets uint) (uint, bool) {
	var firstName string // defined vars in function, local scope.
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

	continueFlag := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if !continueFlag {
		fmt.Println("One of the inputs failed.")
		return remainingTickets, continueFlag
	}

	firstNames := []string{}
	for _, booking := range *bookings {
		firstNames = append(firstNames, booking.firstName) // Extract "firstName" property of struct
	}

	fmt.Printf("The first names of bookings are %v\n", firstNames)

	// example syntax for a switch statement
	// city := "London"
	// switch city {
	// 	case "New York", "Boston":
	// 		execute code for new york or boston tickets
	// 	case "Singapore", "Shanghai":
	// 		execute code for singapore or shanghai
	// 	case "etc etc":
	// 		and so on
	// 	default:
	//		fmt.println("No city selected!")
	//		break
	// }

	createArrays(bookings, firstName, lastName, email, userTickets)
	remainingTickets -= userTickets

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)

	// fmt.Printf("Remaining tickets: %v", remainingTickets)
	return remainingTickets, continueFlag // Return the updated value
}

func createArrays(bookings *[]UserData, firstName string, lastName string, email string, userTickets uint) {
	// Arrays in Go are of a fixed size
	// What kind of value do we want to store here? A list of names of users who booked tickets!
	// Can be empty or already have things in it.
	// No mixed types allowed in an array.

	// var bookings = [50]string{"Erick", "Tom", "Pete"}
	// var bookings [50]string

	// But what about dynamic size list? A slice - array type under hood, but is dynamic in size.
	// create a mapo for a user:
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	} // map keys and values can only have same type.
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	*bookings = append(*bookings, userData)
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(*bookings))

	fmt.Printf("These are all our bookings: %v\n", *bookings)

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) >= 1 && len(lastName) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets > remainingTickets
	if !isValidEmail || !isValidName || !isValidTickets {
		return false
	} else {
		return true
	}
}
