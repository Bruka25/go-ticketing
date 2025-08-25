package main

import "fmt"

func main() {
var confName = "Go conference"
const conferenceTickets = 50
var remainingTickets = 50
//fmt.Println("welcome to our booking application")
//fmt.Println("we are happy to have you here") 
fmt.Printf("welcome to %v booking application\n", confName)
fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
fmt.Println("Get your tickets here to attend")
} 