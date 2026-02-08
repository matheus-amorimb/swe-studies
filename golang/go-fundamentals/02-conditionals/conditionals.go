package main

import "fmt"

func Conditionals() {
	fmt.Println("Hello Conditionals")
	// BasicConditionals()
	// Switch()
	ConditionalExercise()
}

func BasicConditionals() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

func InitialStatementIfBlock() {
	/*
		An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.
		if INITIAL_STATEMENT; CONDITION {
		}

		instead of:
			length := getLength(email)
			if length < 1 {
					fmt.Println("Email is invalid")
			}
	*/
	if length := 10; length < 1 {
		fmt.Println("Email is invalid")
	}
}

func Switch() {
	innerFunc := func(os string) string {
		var creator string

		switch os {
		case "linux":
			creator = "Linus Torvalds"
		case "windows":
			creator = "Bill Gates"

		// all three of these cases will set creator to "A Steve"
		case "macOS":
			fallthrough
		case "Mac OS X":
			fallthrough
		case "mac":
			creator = "A Steve"

		default:
			creator = "Unknown"
		}
		return creator
	}

	fmt.Println(innerFunc("linux"))
}

func ConditionalExercise() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

	finalCost = bulkMessageCost
	if isPremiumUser {
		finalCost = finalCost * (1 - discountRate)
	}

	if accountBalance >= finalCost {
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}
