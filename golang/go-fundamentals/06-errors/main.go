package main

import (
	"errors"
	"fmt"
)

// Error Interface
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	mCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	mSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return (mCustomer + mSpouse), nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

// Custom Error
type useError struct {
	name string
}

func (e useError) Error() string {
	return fmt.Sprintf("$v has a problem with their account", e.name)
}

func sendMessage(msg, userName string) error {
	canSendToUser := true
	if !canSendToUser {
		return useError{name: "matheus"}
	}
	// ...
	return nil
}

// The Errors Package
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

// Panic
func panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	// this panics, but the defer/recover block catches it
	// a truly astonishingly bad way to handle errors
	enrichUser("123")
}

func enrichUser(userID string) {
	_, err := 0, sendMessage(userID, "matheus")
	if err != nil {
		panic()
	}
}

func validateStatus(status string) error {
	sLen := len(status)
	if sLen == 0 {
		return StatusError{message: "status cannot be empty"}
	}

	if sLen > 140 {
		return StatusError{message: "sstatus exceeds 140 characters"}
	}

	return nil
}

type StatusError struct {
	message string
}

func (e StatusError) Error() string {
	return e.message
}

func main() {}
