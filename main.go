package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	remainingTickets := conferenceTickets
	var bookings [50]string
	fmt.Println("Hello World from GO")
	//fmt.Print("Welcome to our Conference \"", conferenceName, "\" booking application\n")
	//fmt.Print("We had total of ", conferenceTickets, " tickets, and the available are ", remainingTickets, " tickets\n")

	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Print("Your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Your email address : ")
	fmt.Scan(&email)

	fmt.Print("Ticket No : ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Array all data %v\n", bookings)
	fmt.Printf("Array first item %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array size %v\n", len(bookings))

	//fmt.Printf("Type of conferenceName is %T and ,conferenceTickets is %T ,and remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our Conference \"%v\" for booking application\n", conferenceName)
	fmt.Printf("We had total of %v tickets, and the available are %v tickets\n", conferenceTickets, remainingTickets)
	fmt.Printf("Thanks %v for booking %v tickets, you will receive a confirmation email on %v . \n", bookings[0], userTicket, email)

	// GO   	JAVA
	// int8     byte
	// int16    short
	// int32    int
	// int64    long

}
