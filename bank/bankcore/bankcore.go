package bankcore

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// TransformTo ...
func (a *Account) TransformTo(amount float64, ta *Account) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("Balance can not less than 0 when transform ")
	}

	a.Withdraw(amount)

	ta.Deposit(amount)

	return nil
}

// Bank ...
type Bank interface {
	Statement() string
}

// Statement ...
func Statement(b Bank) string {
	return b.Statement()
}
