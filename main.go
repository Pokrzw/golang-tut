package main

import (
	"fmt"
	"strings"
)

func main() {

	//write a simple ticket booking app
	var confName = "Go conference"
	const confTicks = 50
	var remainingTicks uint = 50
	var bookings []string

	fmt.Println("Book your tickets for", confName, "there")
	fmt.Printf("We have a total of %v tickets, %v remain\n", confTicks, remainingTicks)
	for {

		var firstName string
		var lastName string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("How many tickets u want to book?: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTicks {
			fmt.Printf("Only %v tickets left", remainingTicks)
			continue //sprawia, ze pomija caly ten kod ponizej i zaczyna
			//kolejna iteracje petli
		}

		email := firstName + "." + lastName + "@gmail.com"
		fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
		fmt.Printf("%v's email: %v\n", firstName, email)
		bookings = append(bookings, firstName+" "+lastName)

		remainingTicks = remainingTicks - userTickets
		fmt.Printf("%v tickets left\n", remainingTicks)
		fmt.Printf("All reservations: \n")

		for index, booking := range bookings {
			fmt.Printf("%v. %v\n", index+1, strings.Fields(booking)[0])
		}

		noBitches := remainingTicks == 0
		if noBitches {
			//end program
			fmt.Println("no mo tickets")
			break
		}
	}
}
