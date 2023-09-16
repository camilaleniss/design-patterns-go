package command

// Command
// Want objects to represent a operation
// Sometimes you cannot undo assignments

// An object that represents an intruction to performn a
// particular action.

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount,
		"\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount,
			"\b, balance is now", b.balance)
		return true
	}
	return false
}

// to actually call or undo the methods
type Command interface {
	Call()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

// to know if a which command will be run and if succedeed
type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}

	// in our case the undo operation it's symetric
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	fmt.Println(ba)
}
