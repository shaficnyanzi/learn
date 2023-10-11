package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 4 && len(lastName) >= 4
	isValidEmail := strings.Contains(email, "@gmail")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}