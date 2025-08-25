package main

import "fmt"

func main() {
var confName = "Go conference"
const conferenceTickets = 50
var remainingTickets = 50
fmt.Printf("conferenceTickets is %T, confName is %T, remainingTickets is %T\n", conferenceTickets, confName, remainingTickets)
fmt.Printf("welcome to %v booking application\n", confName)
fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
fmt.Println("Get your tickets here to attend")
	var userName string
	var userTickets int
	// ask a user for their name

	userName = "Bruka"
	userTickets = 2

	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)
	fmt.Println("Thank you for booking, see you there!")
}