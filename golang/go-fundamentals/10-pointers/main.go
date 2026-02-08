package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//exercise1
	// myString := "English, motherfubber, do you speak it?"
	// myStringPtr := &myString
	// excercise1(myStringPtr)

	//exercise2
	excercise2()
}

func excercise1(message *string) {
	if message == nil {
		return
	}

	words := strings.Fields(*message)
	forbiddenWords := []string{
		"fubb",
		"shiz",
		"witch",
	}

	for idx, word := range words {
		for _, forbiddenWord := range forbiddenWords {
			if strings.Contains(word, forbiddenWord) {
				mask := strings.Repeat("*", len(forbiddenWord))
				words[idx] = strings.ReplaceAll(word, forbiddenWord, mask)
			}
		}
	}

	*message = strings.Join(words, " ")
	fmt.Println(*message)
}

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func excercise2() {
	customer := &customer{
		id:      1,
		balance: 1000.0,
	}

	transaction := transaction{
		customerID:      1,
		amount:          500.0,
		transactionType: transactionWithdrawal,
	}

	updateBalance(customer, transaction)

	fmt.Println(*customer)
}

func updateBalance(customer *customer, transaction transaction) error {
	switch transaction.transactionType {
	case transactionWithdrawal:
		if customer.balance < transaction.amount {
			return errors.New("insufficient funds")
		}
		customer.balance -= transaction.amount
	case transactionDeposit:
		customer.balance += transaction.amount
	default:
		return errors.New("unknown transaction type")
	}

	return nil
}
