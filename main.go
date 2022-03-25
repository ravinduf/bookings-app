package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

// Package level variables
var conferenceName = "Go conference"
const conferenceTickets uint = 50
var remainingTickets uint = conferenceTickets
var bookings = []string{"test"}

func main() {
	// var conferenceName string= "Go conference"
	// conferenceName := "Go conference" // cannot do for consts
	// const conferenceTickets uint = 50
	// var remainingTickets uint = conferenceTickets

	greetUsers()

	// var bookings = [50]string{"Ravindu", "Senal"}
	// var bookings [50]string

	// slice
	// bookings := []string{"test"}

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			
			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings: %v\n\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number is not valid")
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our conference %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("The first names of bookings: %v\n\n", firstNames)

	return firstNames
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will recieve a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainig for %v\n", remainingTickets, conferenceName)

}
