package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 30
	fmt.Println("Hello World from GO")
	//fmt.Print("Welcome to our Conference \"", conferenceName, "\" booking application\n")
	//fmt.Print("We had total of ", conferenceTickets, " tickets, and the available are ", remainingTickets, " tickets\n")
	fmt.Printf("Welcome to our Conference \"%v\" for booking application\n", conferenceName)
	fmt.Printf("We had total of %v tickets, and the available are %v tickets\n", conferenceTickets, remainingTickets)
}
