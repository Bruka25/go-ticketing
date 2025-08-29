package main

import (
	"fmt"
	"go-ticketing/share"
)

var confName = "Go conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0) // Updated to use userData struct

type userData struct {
    firstName       string
    lastName        string
    email           string
    numberOfTickets uint
}

func main() {
    greetUsers()

    fmt.Printf("conferenceTickets is %T, confName is %T, remainingTickets is %T\n", conferenceTickets, confName, remainingTickets)

    for {
        firstName, lastName, email, userTickets := getUserInput()

        isValidName, isValidEmail, isValidTicketNumber := share.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

        if isValidName && isValidEmail && isValidTicketNumber {
            bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

            // Display the first names of all bookings
            var firstNames = getFirstName(bookings)
            fmt.Printf("The first names of bookings are: %v\n", firstNames)

            if remainingTickets == 0 {
                // End the program
                fmt.Println("The conference is booked out")
                break
            }
        } else {
            if !isValidName {
                fmt.Println("Your first name or last name is too short")
            }
            if !isValidEmail {
                fmt.Println("Your email address is not valid")
            }
            if !isValidTicketNumber {
                fmt.Println("Your ticket number is invalid")
            }
        }
    }
}

func greetUsers() {
    fmt.Printf("Welcome to %v the booking application!\n", confName)
    fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")
}

func getFirstName(bookings []userData) []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking.firstName)
    }
    return firstNames
}

func getUserInput() (string, string, string, uint) {
    var firstName string
    var lastName string
    var email string
    var userTickets uint

    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)

    fmt.Println("Enter your Email:")
    fmt.Scan(&email)

    fmt.Println("Enter number of tickets:")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
    // Update remaining tickets
    remainingTickets = remainingTickets - userTickets

    // Create a userData struct to store user data
    var user = userData{
        firstName:       firstName,
        lastName:        lastName,
        email:           email,
        numberOfTickets: userTickets,
    }
    bookings = append(bookings, user)

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.sleep(10 * time.second)
	var tickets = fmt.Sprintf(" %v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket\n %v to email address %v\n", tickets, email)
	fmt.Println("###################")
}
