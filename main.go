package main

import (
	"fmt"
	"strings"
)

func main() {

	var confName = "IPL CSK tickets"

	const confTickets = 100
	var remainingTickets uint = 100
	booking := []string{}

	greetUsers(confName, confTickets, remainingTickets)

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets = remainingTickets - userTickets
			booking = append(booking, firstName+" "+lastName)

			fmt.Printf("Thanks %v %v for booking %v tickets. Confirmation email has been sent to %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v %v tickets are currently remaining\n", remainingTickets, confName)

			// call print first name function here

			printFirstNames(booking)

			if remainingTickets == 0 {
				// end for loop
				fmt.Println("Tickets are sold out")
				break

			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("invalid email")
			}
			if !isValidTicket {
				fmt.Println("invalid number of tickets")
			}

		}
	}
}

func greetUsers(conferenceName string, confTicks int, remainTicks uint) {
	fmt.Printf("Welcome to %v booking\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", confTicks, remainTicks)
	fmt.Println("GET YOUR TICKETS HERE")

}

func printFirstNames(booking []string) []string {
	firstNames := []string{}
	for _, bookings := range booking {
		var names = strings.Fields(bookings)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}
