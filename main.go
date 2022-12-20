package main

import "fmt"

func main() {
	var confereceName = "Go Conference" // Variável string
	const conferenceTicket = 50         // Valor constante, que não pode ser alterado
	var remaningTickets uint = 50       // Valor uint pois essa variável não poderá conter valor negativo

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferceName is %T\n", conferenceTicket, remaningTickets, confereceName)

	fmt.Printf("Welcome to %v booking application\n", confereceName) // Output formatado. {%v} significa printar no valor padrão
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remaningTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
