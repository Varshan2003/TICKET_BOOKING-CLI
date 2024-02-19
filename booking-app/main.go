package main

import (
	"fmt"
	"time"
)

var appName = "TicketBooking"
var totalTickets int = 50
var availableTickets int = totalTickets
var Users = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	ticketNumber int
}

func main() {
	greeting()
	for {
		firstName, lastName, email, ticketNumber := takeUserInput()
		isValidName, isValidMail, isValidTicket := ValidateUserInput(firstName, lastName, email, ticketNumber, availableTickets)
		booking(isValidName, isValidMail, isValidTicket, firstName, lastName, email, ticketNumber)
		go genSendTicket(firstName, email, ticketNumber)

	}
}

func greeting() {
	fmt.Printf("Welcome to %v app\n", appName)
	fmt.Printf("Out of %v tickets we still have %v tickets\n", totalTickets, availableTickets)
}

func takeUserInput() (string, string, string, int) {
	var firstName string
	fmt.Printf("Enter your firstName: ")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Printf("Enter your lastName: ")
	fmt.Scan(&lastName)

	var email string
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	var ticketNumber int
	fmt.Printf("Enter the number of tickets: ")
	fmt.Scan(&ticketNumber)
	return firstName, lastName, email, ticketNumber
}

func booking(isValidName bool, isValidMail bool, isValidTicket bool, firstName string, lastName string, email string, ticketNumber int) {
	if isValidName && isValidMail && isValidTicket {

		var User = UserData{
			firstName:    firstName,
			lastName:     lastName,
			email:        email,
			ticketNumber: ticketNumber,
		}

		Users = append(Users, User)
		fmt.Printf("list of bookings %v\n", Users)

		availableTickets -= ticketNumber
		firstNames := []string{}
		for _, user := range Users {
			firstNames = append(firstNames, user.firstName)
		}

		fmt.Printf("Thank you %v %v! for booking %v tickets, you will be notified @%v\n", firstName, lastName, ticketNumber, email)
		fmt.Printf("Total users are : %v\n", firstNames)
		fmt.Printf("Total tickets available: %v\n", availableTickets)
	} else {
		if !isValidName {
			fmt.Printf("Please enter a valid Name:)\n")
		} else if !isValidMail {
			fmt.Printf("Please enter a valid mail:)\n")
		} else {
			fmt.Printf("Invalid number of tickets!!\n")
		}
	}
}

func genSendTicket(firstName string, email string, ticketNumber int) {
	time.Sleep(10 * time.Second)
	fmt.Print("--------------------\n")
	fmt.Printf("Sending tickets to %v @%vemail for the tickets %v\n", firstName, email, ticketNumber)
	fmt.Printf("-------------------\n")
}
