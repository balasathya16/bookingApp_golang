package main

import "fmt"

func main() {

	var confName = "IPL CSK tickets"

	const confTickets = 40000

	var remainingTickets = 40000

	fmt.Printf("Welcome to %v booking\n", confName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("GET YOUR TICKETS HERE")

	var firstName string
	var lastName string
	var email string

	//var userTickets int

	userTickets := 2
	// get user name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thanks! User %v booked %v tickets\n", firstName, userTickets)

}
