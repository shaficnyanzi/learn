package main

import (
	"fmt"
	"strings"
)
    const conferenceTicket int = 60
    var conferenceName = "Go conference"
	var remainingTickets uint = 60
	var bookings []string

func main() {

	greetusers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		//function that contains user validation data
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//function to book the ticket
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName )  

			//function that calls out the first names of users who have booked a ticket
			firstNames := getfirstNames()
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
func greetusers() {
	fmt.Println("welcome to", conferenceName, "booking system")
	fmt.Println("We are open with", conferenceTicket, " tickets still available,we have only", remainingTickets, "available")
	fmt.Println("feel at home here!")

}
func getfirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func getUserInput()(string,string, string, uint){
	var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your firstname: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastName: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets

}
func bookTicket(remainingTickets uint,userTickets uint,bookings [] string, firstName string, lastName string,email string, conferenceName string){
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName +" "+ lastName)

	fmt.Printf("user %v %v with email address %v has bought %v tickets\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets are available now for the %v\n", remainingTickets, conferenceName)
}
