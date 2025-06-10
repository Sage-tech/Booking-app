package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var userTickets uint
	var fristName string
	var lastName string
	var email string
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&fristName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email: ")
		fmt.Scan(&email)
		fmt.Println("Enter how many tickets you would like to buy: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, fristName+" "+lastName)

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", fristName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fristNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			fristNames = append(fristNames, names[0])
		}
		fmt.Printf("The frist names of bookings are: %v\n", fristNames)
	}
}
