package helper

import (
	"strings"
)

// moving things to another subfolder to experiment with package imports!

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) >= 1 && len(lastName) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets > remainingTickets
	if !isValidEmail || !isValidName || !isValidTickets {
		return false
	} else {
		return true
	}
}
