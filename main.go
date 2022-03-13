package main

import "fmt"

func main() {
	var confName = "Go conference"
	const confTicks = 50
	var remainingTicks = 50

	fmt.Println("Book your tickets for", confName, "there")
	fmt.Printf("We have a total of %v tickets, %v remain\n", confTicks, remainingTicks)

	var userName string
	var userTickets int
	//komentarz

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
