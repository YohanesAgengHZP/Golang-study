package main

import (
	"bufio"
    "fmt"
    "os"
	"booking-app/helper"
)

func main() {
    const conferenceTicket int = 50
    var conferenceTicketLeft int = 50

    greetUser()
    startInformation()

    for {
        var bookedTicket []int

        for {
            fullName := helper.GetValidFullName()

            email := helper.GetValidEmail()
            userTicket := helper.GetValidTicket(conferenceTicketLeft)

            conferenceTicketLeft -= userTicket
            bookedTicket = append(bookedTicket, userTicket)

            fmt.Printf("Hello %v, thank you for booking a ticket. Here is your ticket %v, the remaining tickets left is %v \n", fullName, userTicket, conferenceTicketLeft)
            fmt.Printf("Booked via this email %v, please check for confirmation\n", email)

            // Clear input buffer
            bufio.NewReader(os.Stdin).ReadBytes('\n')

            // Check if the user wants to book another ticket
            if !askForAnotherBooking() {
                break
            }
        }
        fmt.Println("Thank you for booking tickets. Have a great day!")
        break
    }
}

func greetUser() {
    conferenceName := "Go Conference"
    fmt.Printf("\nHello welcome to this Conference %v \n", conferenceName)
}

func startInformation() {
    const conferenceTicket int = 50
    var conferenceTicketLeft int = 50
    fmt.Println("-------------------------------------------------------")
    fmt.Printf("Check our this one!, ticket left is %v from total %v \n", conferenceTicketLeft, conferenceTicket)
    fmt.Println("-------------------------------------------------------")
}

func askForAnotherBooking() bool {
    fmt.Println("Do you want to book another ticket? (yes/no):")
    var anotherBooking string
    fmt.Scanln(&anotherBooking)
    return anotherBooking == "yes"
}


