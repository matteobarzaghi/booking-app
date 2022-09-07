package main

import "fmt"

func main() {
	fmt.Print("Hello Go!")

	// variables
	var conferenceName = "Go Conference"

	// constants
	const conferenceTickets = 50
	//var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("Get yuor tickets here to attend!")
	fmt.Println(conferenceTickets, "tickets are left.")

}
