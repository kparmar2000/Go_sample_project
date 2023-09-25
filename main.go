package main

import (
	"fmt"
	"strconv"
	"time"
)

var fname string
var lname string
var email string
var remainingticket uint
var ticketbought uint
var fullname = make([]map[string]string, 0)

func main() {
	remainingticket = 50

	for {

		var myMap = make(map[string]string) // Create a new map in each iteration

		fmt.Println("Enter your first name:")
		fmt.Scan(&fname)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lname)
		fmt.Println("Enter your email id:")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets to buy:")
		fmt.Scan(&ticketbought)
		go sendTicket(fname, lname, ticketbought)
		if ticketbought > remainingticket {
			fmt.Printf("Sorry, we only have %v ticket(s) left. Please enter a lower amount.\n", remainingticket)
			continue
		}

		myMap["firstName"] = fname // Update map with user input
		myMap["lastName"] = lname
		myMap["email"] = email
		myMap["ticketsBought"] = strconv.FormatUint(uint64(ticketbought), 10)

		greeting(fname, email, remainingticket, ticketbought)
		remainingticket = remainingticket - ticketbought

		fullname = append(fullname, myMap)

		fmt.Println(fullname)
		fmt.Println(len(fullname))

		printfname()

		if remainingticket == 0 {
			fmt.Println("Sorry, we are all sold out. Please try reducing the ticket amount or come back later.")
			break
		}
	}
}

func greeting(fname string, email string, remainingticket uint, ticketbought uint) {
	fmt.Printf("Hello, %v\n", fname)
	fmt.Println("Welcome to GoTheater")
	fmt.Printf("Thanks %v for buying %v ticket(s). Confirm your tickets on %v \n", fname, ticketbought, email)
	fmt.Printf("We have %v tickets left\n", remainingticket)
}

func printfname() {
	firstnames := []string{}
	for _, fname := range fullname {
		firstnames = append(firstnames, fname["firstName"])
	}
	fmt.Println("First names:", firstnames)
}
func sendTicket(fname string, lname string, ticketbought uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v %v ticket %v ", fname, lname, ticketbought)
	fmt.Printf("==============================")
	fmt.Printf("%v", ticket)
	fmt.Printf("==============================")

}
