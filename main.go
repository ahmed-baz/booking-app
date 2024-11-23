package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	remainingTickets := 30
	fmt.Println("Hello World from GO")
	//fmt.Print("Welcome to our Conference \"", conferenceName, "\" booking application\n")
	//fmt.Print("We had total of ", conferenceTickets, " tickets, and the available are ", remainingTickets, " tickets\n")
	fmt.Printf("Welcome to our Conference \"%v\" for booking application\n", conferenceName)
	fmt.Printf("We had total of %v tickets, and the available are %v tickets\n", conferenceTickets, remainingTickets)

	var userName string
	var userTicket int

	fmt.Print("Your name : ")
	fmt.Scan(&userName)
	fmt.Print("Ticket No : ")
	fmt.Scan(&userTicket)

	fmt.Printf("Type of conferenceName is %T and ,conferenceTickets is %T ,and remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("The user %v booked %v tickets. \n", userName, userTicket)

	// GO   	JAVA
	// int8     byte
	// int16    short
	// int32    int
	// int64    long

}
