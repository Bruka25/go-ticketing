package main

import (
	"fmt"
	"strings"
)

func main() {
var confName = "Go conference"
const conferenceTickets uint = 50
var remainingTickets uint= 50
var bookings = []string{}

greetUsers(confName, int(conferenceTickets), int(remainingTickets))

fmt.Printf("conferenceTickets is %T, confName is %T, remainingTickets is %T\n", conferenceTickets, confName, remainingTickets)

for {


	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask a user for their name
	
	

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)
	fmt.Println("Enter your first name:")
	 fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)


	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	if isValidName && isValidEmail && isValidTicketNumber {
		remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	//userName = "Bruka 
	//userTickets = 2
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The First names of bookings are: %v\n", firstNames)
	if  remainingTickets == 0 {
		//end the program
		fmt.Println("The conference is booked out")
		break
	}
	
	} else {
		if !isValidName {
			fmt.Println("your first name or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("your email address is not valid")
		}
		if !isValidTicketNumber {
			fmt.Println("your ticket number is invalid")
		}
	}
	
}	
}

func greetUsers(confName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v the booking application!\n", confName)

    fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")

}
