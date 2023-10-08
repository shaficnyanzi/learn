package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go conference"
	const conferenceTicket int = 60
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("welcome to", conferenceName, "booking system")
	fmt.Println("We are open with", conferenceTicket, "still available,we have only", remainingTickets, "available")
	fmt.Println("feel at home here!")

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

		if userTickets > remainingTickets {

			fmt.Printf("You can't book more than the %v but from 1-%v\n", remainingTickets, remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		//userTicket = 4
		//fmt.Printf("whole slice: %v\n",bookings)
		//fmt.Printf("first value of slice: %v\n",bookings[0])
		//fmt.Printf("type of slice: %T\n",bookings)
		//fmt.Printf("length of the slice: %v\n",len(bookings))

		fmt.Printf("user %v %v with email address %v has bought %v tickets\n", firstName, lastName, userEmail, userTickets)
		fmt.Printf("%v tickets are available now for the %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of all our bookings are %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our tickets are all booked.come back next year")
			break
		}
	}
}
