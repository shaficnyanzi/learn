package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go conference"
	const conferenceTicket int = 60
	var remainingTickets uint = 60
	var bookings []string

	greetusers(conferenceName, conferenceTicket, remainingTickets)

	for {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

		fmt.Println("Enter your firstname: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastName: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&userEmail)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 4 && len(lastName) >= 4
		isValidEmail := strings.Contains("email", "@gmail")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("user %v %v with email address %v has bought %v tickets\n", firstName, lastName, userEmail, userTickets)
			fmt.Printf("%v tickets are available now for the %v\n", remainingTickets, conferenceName)

			//function that calls out the first names of users who have booked a ticket
			firstNames := (bookings)
			fmt.Printf("The first names of all our bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our tickets are all booked.come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you've entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email misses the @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("check the number of tickets you've enterd")
			}
			//fmt.Printf("Your input is invalid,please try again")

		}
	}
}
func greetusers(confName string, confTicket int, remainingTicket uint) {
	fmt.Println("welcome to", confName, "booking system")
	fmt.Println("We are open with", confTicket, " tickets still available,we have only", remainingTicket, "available")
	fmt.Println("feel at home here!")

}
func getfirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
