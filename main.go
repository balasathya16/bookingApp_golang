package main

import "fmt"

func main() {

	var confName = "IPL CSK tickets"

	const confTickets = 40000
	var remainingTickets uint = 40000
	var booking [50]string

	fmt.Printf("Welcome to %v booking\n", confName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("GET YOUR TICKETS HERE")

	fmt.Printf("%v %v\n", booking[0], booking[10])

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// get user name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	booking[0] = firstName + " " + lastName

	fmt.Printf("whole array: %v\n", booking)
	fmt.Printf("first val: %v\n", booking[0])
	fmt.Printf("Type: %T\n", len(booking))
	tt

	fmt.Printf("Thanks %v %v for booking %v tickets. Confirmation email has been sent to %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v %v tickets are currently remaining\n", remainingTickets, confName)

}
