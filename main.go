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

	// fmt.Println("Welcome to", conferenceName, "booking appllication")
	// fmt.Println("Get your tickets here to attend")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{"Ravindu", "Senal"}
	// var bookings [50]string

	// slice
	bookings := []string{"test"}

	for {

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
		// fmt.Scan(&userTickets)

		// userName = "Tom"

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[1])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice length: %v\v", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will recieve a confirmation email at %v\n",
				firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remainig for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings: %v\n\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			// break
			// continue

		}

	}
}
