package helper

import (
    "net/mail"
    "unicode"
)

func Valid(emailUser string) bool {
    _, err := mail.ParseAddress(emailUser)
    return err == nil
}

func IsValidName(name string) bool {
    for _, char := range name {
        if !unicode.IsLetter(char) {
            return false
        }
    }
    return true
}