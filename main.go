package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print("Hello Go!")

	// variables
	var conferenceName = "Go Conference"

	// constants
	//const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("Get yuor tickets here to attend!")
	fmt.Println(remainingTickets, "tickets are left.")

	for {

		// input: ask user for their name
		var userName string
		var userLastName string
		var userEmail string
		var userWantedTickets int

		fmt.Println("What is your name?")
		fmt.Scan(&userName)

		fmt.Println("What is your last name?")
		fmt.Scan(&userLastName)

		fmt.Println("What is your email?")
		fmt.Scan(&userEmail)

		fmt.Println("How many tickets do you want?")
		fmt.Scan(&userWantedTickets)

		if userWantedTickets <= remainingTickets {
			fmt.Println("Hello ", userName, " ", userLastName, "!\nWe sent an email to your address:", userEmail, ". Please verify it.\nWe're processing yuor request for", userWantedTickets, "tickets.")
			remainingTickets = remainingTickets - userWantedTickets

			fmt.Println("Tickets left: ", remainingTickets)

			// Arrays and Slices

			var bookings []string
			bookings = append(bookings, userName+" "+userLastName)

			// loop to print only names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets <= 0 {
				// end program
				fmt.Println("The conference is booked out, come back next year.")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, you can't book %v tickets\n", remainingTickets, userWantedTickets)
		}
	}
}
