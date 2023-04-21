package main

import "fmt"

func main() {

	var confName = "IPL CSK tickets"

	const confTickets = 40000

	var remainingTickets = 40000

	fmt.Printf("Welcome to %v booking\n", confName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("GET YOUR TICKETS HERE")

}
