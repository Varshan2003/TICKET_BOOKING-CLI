package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, ticketNumber int, availableTickets int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) >= 2
	isValidMail := strings.Contains(email, "@")
	isValidTicket := ticketNumber <= availableTickets
	return isValidName, isValidMail, isValidTicket
}
