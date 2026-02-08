package main

import (
	"fmt"
)

func Variables() {
	// BasicVariabless()
	// ShortVariableDeclaration()
	FormattingStrings()
}

func BasicVariabless() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func ShortVariableDeclaration() {
	//The walrus operator, :=, declares a new variable and assigns a value to it in one line.
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)
}

func Comments() {
	// This is a single line comment

	/*
		This is a multi-line comment neither of these comments will execute as code
	*/
}

func TypeSizes() {
	/*
		- Whole Numbers (No Decimal)
		int  int8  int16  int32  int64

		- Positive Whole Numbers (No Decimal)
		uint uint8 uint16 uint32 uint64 uintptr

		- Signed Decimal Numbers
		float32 float64

		- Complex Numbers
		complex64 complex128
	*/

	accountAgeFloat := 2.6
	//Converting Between Types
	accountAgeInt := int32(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

func Constants() {
	const premiumPlanName = "Premium Plan"
	fmt.Println("plan:", premiumPlanName)

	//Constants must be known at compile time.

	// The current time can only be known when the program is running
	// const currentTime = time.Now()
}

func FormattingStrings() {
	//fmt.Printf - Prints a formatted string to standard out.
	//fmt.Sprintf() - Returns the formatted string

	//The %v variant prints the Go syntax representation of a value
	s := fmt.Sprintf("I am %v years old", 10)
	fmt.Println(s)

	/*
		string: %s
		int: %d
		float: %f
			- ".2" rounds the number to 2 decimal places
			- s := fmt.Sprintf("I am %.2f years old", 10.523)
	*/

	const name = "Saul Goodman"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %s, your open rate is %.2f percent \n", name, openRate)
	fmt.Print(msg)
}

func TypeInference() {
	// penniesPerText(int) := 2
	penniesPerText := 2.0

	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}

func FormatPractice() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %f, Is Subscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)
	// Don't touch below this line

	fmt.Println(userLog)
}
