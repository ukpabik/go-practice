package main

import "errors"

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

// Don't touch above this line

// ?

func updateBalance(user *customer, trans transaction) error {
	if user.balance < trans.amount {
		return errors.New("insufficient funds")
	}

	if trans.transactionType != transactionDeposit && trans.transactionType != transactionWithdrawal {
		return errors.New("unknown transaction type")
	}

	if trans.transactionType == transactionDeposit {
		user.balance += trans.amount
	} else {
		user.balance -= trans.amount
	}
	return nil
}
