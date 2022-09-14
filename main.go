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
	// const conferenceTickets = 50
	var remainingTickets int = 50

	greetUsers(conferenceName)
	fmt.Println("Get yuor tickets here to attend!")
	fmt.Println(remainingTickets, "tickets are left.")

	for {

		userName, userLastName, userEmail, userWantedTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userName, userLastName, userEmail, userWantedTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			fmt.Println("Hello ", userName, " ", userLastName, "!\nWe sent an email to your address:", userEmail, ". Please verify it.\nWe're processing yuor request for", userWantedTickets, "tickets.")
			remainingTickets = remainingTickets - userWantedTickets

			fmt.Println("Tickets left: ", remainingTickets)

			// Arrays and Slices
			var bookings []string
			bookings = append(bookings, userName+" "+userLastName)

			firstNames := printFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets <= 0 {
				// end program
				fmt.Println("The conference is booked out, come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name are not valid!")
			}
			if !isValidEmail {
				fmt.Println("Email address entered is not valid!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is not valid!")
			}

			fmt.Printf("We only have %v tickets remaining, you can't book %v tickets\n", remainingTickets, userWantedTickets)
		}
	}
}

func greetUsers(confName string) {
	fmt.Println("Welcome to our conference!")
}

func printFirstNames(bookings []string) []string {

	// loop to print only names
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, int) {
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

	return userName, userLastName, userEmail, userWantedTickets
}
