package command

import "fmt"

type BankAccount struct {
	Balance int
}

func Deposit(ba *BankAccount, amount int) {
	fmt.Println("Depositing", amount)
	ba.Balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
	if ba.Balance >= amount {
		fmt.Println("Withdrawing", amount)
		ba.Balance -= amount
	}
}

func main() {
	ba := &BankAccount{0}
	var commands []func()

	// here we are just adding commands to be implemented in order
	// in some cases can be practical
	// where we don't need to have all control just perform actions
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 100)
	})

	for _, cmd := range commands {
		cmd()
	}
}
