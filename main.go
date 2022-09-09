package main

import "fmt"

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
	fmt.Println("Hello ", userName, " ", userLastName, "!\nWe sent an email to your address:", userEmail, ". Please verify it.\nWe're processing yuor request for", userWantedTickets, "tickets.")
	remainingTickets = remainingTickets - userWantedTickets
	fmt.Println("Tickets left: ", remainingTickets)

	// Arrays and Slices

	var bookings []string
	bookings = append(bookings, userName+" "+userLastName)
	fmt.Printf("The whole slice: %v\n", bookings)

}
