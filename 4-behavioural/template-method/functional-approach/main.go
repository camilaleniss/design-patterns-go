package main

import "fmt"

// notice that here we don't have interfaces or structs
// just funcs
func PlayGame(start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}

func main() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	// define the functions at runtime
	start := func() {
		fmt.Println("Starting a game of chess.")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n",
			turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	//Let's play!
	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
