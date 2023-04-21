package main

import "fmt"

func main() {

	var confName = "IPL CSK tickets"

	const confTickets = 40000

	var remainingTickets = 40000

	fmt.Println("Welcome to", confName, "booking")
	fmt.Println("we have a total of", confTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("GET YOUR TICKETS HERE")

}
