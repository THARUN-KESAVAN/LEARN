package main

import (
	"fmt"
	"strings"
)

func main() {

	concertname := "GO CONFERENCE"
	var totalTicketsAvailable uint = 50
	var RemainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var usertickets uint
	var bookings []string
	for {
		greet(concertname, totalTicketsAvailable)
		firstName, lastName, email, usertickets := userinp(firstName, lastName, email, usertickets)

		isvalidfirstName, isvalidLastName, isvalidemail, isvalidusertickets := isvalidinput(firstName, lastName, email, usertickets, RemainingTickets)

		if isvalidfirstName || isvalidLastName && isvalidemail && isvalidusertickets {

			RemainingTickets -= usertickets
			bookings = append(bookings, firstName)
			fmt.Printf("The users booked are%v\n", bookings)
			message(firstName, lastName, email, usertickets, RemainingTickets, totalTicketsAvailable, concertname)
			totalTicketsAvailable = RemainingTickets
			if RemainingTickets == 0 {

				fmt.Println("Our concert tickets are over, so come back next year")
				break
			}

		}
	}
}
func greet(concertname string, totalTicketsAvailable uint) {
	fmt.Printf("Hey  Welcome to our  %v\n", concertname)
	fmt.Printf("Hey Total TIckets available for concert is  %v\n", totalTicketsAvailable)

}
func userinp(firstName string, lastName string, email string, usertickets uint) (string, string, string, uint) {
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter the total no of tickets needed")
	fmt.Scan(&usertickets)
	return firstName, lastName, email, usertickets
}
func isvalidinput(firstName string, lastName string, email string, usertickets uint, RemainingTickets uint) (bool, bool, bool, bool) {
	isvalidFirstName := len(firstName) >= 2 //here
	isvalidLastName := len(lastName) >= 2
	isvalidemail := strings.Contains(email, "@")
	isvaliduserTicket := usertickets <= RemainingTickets && usertickets > 0
	return isvalidFirstName, isvalidLastName, isvalidemail, isvaliduserTicket
}
func message(firstName string, lastName string, email string, usertickets uint, RemainingTickets uint, totalTicketsAvailable uint, concertname string) {

	fmt.Printf("Hey %v %v, Welcome to our  %v\n", firstName, lastName, concertname)
	fmt.Printf("Total tickets you have booked %v,remaing tickets availabe for %v is %v\n", usertickets, concertname, RemainingTickets)
	fmt.Printf("Confirmation email has been sent to %v\n", email)

}
