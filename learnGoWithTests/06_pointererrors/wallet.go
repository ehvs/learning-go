package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount

}

var ErrInsufficientFunds = errors.New("cannot Withdraw, no plata")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Define um método chamado String que pertence ao tipo Bitcoin.

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
