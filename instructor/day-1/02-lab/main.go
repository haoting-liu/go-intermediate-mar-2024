package main

import "fmt"

/**
TODO: Find abstractions and compose interfaces to fix the commented code
*/

// A current account can be used to deposit and withdraw
type CurrentAccount struct {
	AccountNumber int
	Balance       float64
}

func (s *CurrentAccount) Deposit(amount float64) {
	s.Balance += amount
}

func (s *CurrentAccount) Withdraw(amount float64) {
	if s.Balance >= amount {
		s.Balance -= amount
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func (s *CurrentAccount) GenerateStatement() {
	fmt.Printf("Current Account Statement\nAccount Number: %d\nBalance: $%.2f\n", s.AccountNumber, s.Balance)
}

// A saving account can be used to deposit and withdraw
type SavingsAccount struct {
	AccountNumber int
	Balance       float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
}

func (s *SavingsAccount) Withdraw(amount float64) {
	if s.Balance >= amount {
		s.Balance -= amount
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func (s *SavingsAccount) GenerateStatement() {
	fmt.Printf("Savings Account Statement\nAccount Number: %d\nBalance: $%.2f\n", s.AccountNumber, s.Balance)
}

// A fixed deposit account is not withdrawable
type FixedDepositAccount struct {
	AccountNumber int
	Balance       float64
}

func (s *FixedDepositAccount) Deposit(amount float64) {
	s.Balance += amount
}

func (s *FixedDepositAccount) GenerateStatement() {
	fmt.Printf("Fixed deposit Account Statement\nAccount Number: %d\nBalance: $%.2f\n", s.AccountNumber, s.Balance)
}
func main() {
	// Create instances of both accounts
	savings := &SavingsAccount{AccountNumber: 123, Balance: 1000}
	current := &CurrentAccount{AccountNumber: 999, Balance: 3000}
	fd := &FixedDepositAccount{AccountNumber: 456, Balance: 2000}

	_ = savings // TODO: remove line, this is only to get rid of compilation error (unused variables)
	_ = fd      // TODO: remove line, this is only to get rid of compilation error (unused variables)
	_ = current

	// manageSavingsAccount(savings) //TODO: fix this
	// manageSavingsAccount(current) //TODO: fix this

	// manageFixedDepositAccount(fd) //TODO: fix this
}

/*
// TODO: should be able to deposit, withdraw and print statements
func manageSavingsAccount(as ?) {
}

// TODO: should be able to deposit and print statements
func manageFixedDepositAccount(as ?) {
}
*/
