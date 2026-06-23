package main


import "fmt"

import (
	"engineering_playbook/fundamentals/lund"
)


func main() {

	var programName = "Get a girlfriend program"
	const totalTickets = 50
	var remainingTickets uint = 50


	//as we have provided the uint type explicitly so remainingTickets cannot be negative at all
	// remainingTickets = -2


	fmt.Println("Welcome to our", programName)
	fmt.Printf("We have limited number of ticket : %v\n", totalTickets)
	fmt.Println("Available tickets:", remainingTickets)
	LundChota()
	chut()
	Lund()
	lund.Hello()
	//a new user comes must have a username

	var userName string
	

	userName = "Tanishk"
	userTickets := 50

	fmt.Println(userName)
	fmt.Println(userTickets)

	//supppose the user has booked some tickets

	fmt.Printf("%T is the type of username and %T is the type of tickets\n" , userName , userTickets)
	fmt.Printf("%v has booked %v tickets\n", userName, userTickets)

}
