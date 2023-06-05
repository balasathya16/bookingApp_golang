package main

import (
	"fmt"
	"strings"
)

const confTickets = 100

var confName = "IPL CSK tickets"
var remainingTickets uint = 100
var booking = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTicket(userTickets, firstName, lastName, email)
			// call print first name function here

			firstNames := getFirstNames()
			fmt.Printf("The first names for the bookings are %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("Welcome to %v booking\n", confName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("GET YOUR TICKETS HERE")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, bookings := range booking {
		firstNames = append(firstNames, bookings.firstName)

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// creating a map

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	booking = append(booking, userData)
	fmt.Printf("List of bookings: %v\n", booking)

	fmt.Printf("Thanks %v %v for booking %v tickets. Confirmation email has been sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v %v tickets are currently remaining\n", remainingTickets, confName)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}
