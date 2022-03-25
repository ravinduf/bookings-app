package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string= "Go conference"
	conferenceName := "Go conference" // cannot do for consts
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// var bookings = [50]string{"Ravindu", "Senal"}
	// var bookings [50]string

	// slice
	bookings := []string{"test"}

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, conferenceName)
			
			firstNames := getFirstNames(bookings)

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

func greetUsers(confName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to our conference %v booking application", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("The first names of bookings: %v\n\n", firstNames)

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
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

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will recieve a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainig for %v\n", remainingTickets, conferenceName)

}
