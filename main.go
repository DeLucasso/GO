// Video is here: https://www.youtube.com/watch?v=yyUHQIec83I&t=171s

package main

import "fmt"

func main() {

	// var conferenceName string = "Go Conference"
	// here is the alternative syntax (syntactic sugar - make it nicer for human)
	conferenceName := "Go Conference" // we cant declare "const", only "var"

	const conferenceTickets = 50

	// with uint32 we set only the positive numbers
	var remainingTickets uint32 = 50

	// let's find out the data types we have declared so far.
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets,and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("You can book a ticket easily here.")

	// we have to chose the data type, when the value is not assigned.
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	// fmt.Printf(&userTickets)
	fmt.Println(remainingTickets)

	// with the pointer "&" let's see the memory address for the variable remainingTickets
	// fmt.Println(&remainingTickets)
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

}
