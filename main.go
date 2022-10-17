package main

import (
	"fmt"
	"sync"
	"test/helper"
	"time"
)

// var conferenceName string = "Go conference"
// or
const conferenceTickets int = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // slice of struct
// var bookings = make([]map[string]string, 0) // slice of map
// var bookings = []string{} // slice in golang
// var bookings []string
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	fmt.Printf("types %T,%T,%T\n", conferenceName, conferenceTickets, remainingTickets)

	//	for {

	firstName, lastName, email, userTickets := userInput()

	isValidName, isValidemail, isValidTicketNumber := helper.UserInputValidation(firstName, lastName, email, uint(userTickets), uint(remainingTickets))

	if isValidName && isValidemail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		////////
		firstNames := getFirstName()
		fmt.Printf("The frist names of booking are: %v\n", firstNames)
		////////

		if remainingTickets == 0 {
			fmt.Printf("Our conferience ticktes booked out. Come back next year\n")
			//break
		}
	} else {
		if !isValidName {
			fmt.Printf("First name or last name you entred is too short\n")
		} else if !isValidemail {
			fmt.Printf("Email address you entered doesn't contain '@' sign\n")
		} else if !isValidTicketNumber {
			fmt.Printf("Number of tickets you entered is invalid\n")
		}

	}
	wg.Wait()
	//}
}
func greetUsers() {

	fmt.Printf("Welcome to %v booking applicattion\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
}

func getFirstName() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

///////

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// var userData = make(map[string]string)             //----Map----//
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	fmt.Printf("List of booking is %v\n", userData)

	bookings = append(bookings, userData)
	//bookings = append(bookings, firstName+" "+lastName) //when slice is used

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("Booking list %v\n", bookings)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()

}
