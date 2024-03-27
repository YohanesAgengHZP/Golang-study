package main

import (
    "bufio"
    "fmt"
    "net/mail"
    "os"
    "strings"
    "unicode"
)

func main() {
    const conferenceTicket int = 50
    var conferenceTicketLeft int = 50

    greetUser()
    startInformation()

    for {
        var bookedTicket []int

        for {
            fullName := getValidFullName()

            email := getValidEmail()
            userTicket := getValidTicket(conferenceTicketLeft)

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

func getValidFullName() string {
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
        firstName := strings.TrimSpace(nameParts[0])
        lastName := strings.TrimSpace(nameParts[1])
        if !isValidName(firstName) || !isValidName(lastName) {
            fmt.Println("Invalid input. Please enter a valid first and last name (only letters allowed).")
            continue
        }
        return fullName
    }
}

func getValidEmail() string {
    var emailUser string
    for {
        fmt.Println("Please enter your email:")
        _, err := fmt.Scan(&emailUser)
        if err != nil || !valid(emailUser) {
            fmt.Println("Invalid email. Please enter a valid email address.")
            continue
        }
        break
    }
    return emailUser
}

func getValidTicket(conferenceTicketLeft int) int {
    var userTicket int
    for {
        fmt.Println("Please enter the number of tickets you want to buy:")
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
    return userTicket
}

func askForAnotherBooking() bool {
    fmt.Println("Do you want to book another ticket? (yes/no):")
    var anotherBooking string
    fmt.Scanln(&anotherBooking)
    return anotherBooking == "yes"
}

func valid(emailUser string) bool {
    _, err := mail.ParseAddress(emailUser)
    return err == nil
}

func isValidName(name string) bool {
    for _, char := range name {
        if !unicode.IsLetter(char) {
            return false
        }
    }
    return true
}
