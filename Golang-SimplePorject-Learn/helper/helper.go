package helper

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type UserData struct {
    EmailUser   string
    UserTicket  int
}

func GetValidFullName() string {
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
        if !IsValidName(firstName) || !IsValidName(lastName) {
            fmt.Println("Invalid input. Please enter a valid first and last name (only letters allowed).")
            continue
        }
        return fullName
    }
}

func GetValidEmail() string {
    for {
        fmt.Println("Please enter your email:")
        var emailUser string
        _, err := fmt.Scan(&emailUser)
        if err != nil || !Valid(emailUser) {
            fmt.Println("Invalid email. Please enter a valid email address.")
            continue
        }
        return emailUser
    }
}

func GetValidTicket(conferenceTicketLeft int) int {
    for {
        fmt.Println("Please enter the number of tickets you want to buy:")
        var userTicket int
        _, err := fmt.Scan(&userTicket)
        if err != nil || userTicket <= 0 {
            fmt.Println("Invalid input. Please enter a valid number of tickets.")
            continue
        }
        if userTicket > conferenceTicketLeft {
            fmt.Println("Insufficient tickets available.")
            continue
        }
        return userTicket
    }
}
