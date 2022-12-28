package main

import "fmt"

func main() {
	var confereceName = "Go Conference" // Variável string
	const conferenceTicket = 50         // Valor constante, que não pode ser alterado
	var remaningTickets uint = 50       // Valor uint pois essa variável não poderá conter valor negativo

	fmt.Printf("Welcome to %v booking application\n", confereceName) // Output formatado. {%v} significa printar no valor padrão
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remaningTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// declaração array
		// var bookings [50]string

		// declaração slice
		//var bookings []string
		bookings := []string{}

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

		remaningTickets = remaningTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remainend for %v\n", remaningTickets, confereceName)
	}
}
