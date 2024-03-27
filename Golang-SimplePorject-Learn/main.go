package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/mail"
)

func main() {
    conferenceName := "Go Conference"

    const conferenceTicket int = 50
    var conferenceTicketLeft int = 50

    fmt.Printf("Welcome to %v \n", conferenceName)
    fmt.Printf("Check our this one!, ticket left is %v from total %v \n \n", conferenceTicketLeft, conferenceTicket)

    var firstName string
    var lastName string
    var emailUser string
    var userTicket int

    // Error handling for username
    for {
        fmt.Println("Get your ticket here, and please enter your full name (e.g., First Last):")

        reader := bufio.NewReader(os.Stdin)
        fullName, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input. Please try again.")
            continue
        }

        fullName = strings.TrimSpace(fullName)
        nameParts := strings.SplitN(fullName, " ", 2)
        if len(nameParts) != 2 {
            fmt.Println("Invalid input. Please enter both first and last name separated by a space.")
            continue
        }

        firstName = nameParts[0]
        lastName = nameParts[1]
        break
    }

    // Error handling for email
    for {
        fmt.Println("Please enter your email:")
        _, err := fmt.Scan(&emailUser)
        if err != nil || !valid(emailUser) {
            fmt.Println("Invalid email. Please enter a valid email address.")
            continue
        }
        break
    }

    // Error handling for the number of tickets
    for {
        fmt.Println("Please enter amount of tickets you want to buy:")
        _, err := fmt.Scan(&userTicket)
        if err != nil || userTicket <= 0 {
            fmt.Println("Invalid input. Please enter a valid number of tickets.")
            continue
        }
        if userTicket > conferenceTicketLeft {
            fmt.Println("Insufficient tickets available.")
            continue
        }
        break
    }

    conferenceTicketLeft -= userTicket

    fmt.Printf("Hello %v %v, thank you for booking a ticket. Here is your ticket %v, the remaining tickets left is %v \n", firstName, lastName, userTicket, conferenceTicketLeft)
    fmt.Printf("Booked via this email %v, please check for confirmation\n", emailUser)
}

func valid(emailUser string) bool {
    _, err := mail.ParseAddress(emailUser)
    return err == nil
}
