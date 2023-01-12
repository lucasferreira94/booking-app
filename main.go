package main

import (
	"fmt"
	"strings"
)

func greetUsers(confName string, confTickets int, remaningTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets, remaningTickets uint) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remaningTickets
}

func main() {
	var confereceName = "Go Conference" // Variável string
	const conferenceTicket = 50         // Valor constante, que não pode ser alterado
	var remaningTickets uint = 50       // Valor uint pois essa variável não poderá conter valor negativo
	bookings := []string{}

	greetUsers(confereceName, conferenceTicket, remaningTickets)

	for {

		//userName := ""
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		// ask user for their last name
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		// ask user for their email
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		// ask user for their tickets
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		validateUserInput(firstName)

		if isValidName && isValidEmail && isValidTicketNumber {
			remaningTickets = remaningTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remainend for %v\n", remaningTickets, confereceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remaningTickets == 0 {
				// end program
				fmt.Println("Our conferece is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contian @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}
