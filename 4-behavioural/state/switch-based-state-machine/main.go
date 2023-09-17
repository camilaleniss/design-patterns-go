package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State int

const (
	Locked State = iota
	Failed
	Unlocked
)

func main() {
	// the code we have to guess
	code := "1234"
	state := Locked

	// strings builder to manipulate the entry
	entry := strings.Builder{}

	for {
		switch state {
		case Locked:
			// only reads input when you press Return
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune() //read rune read only one char

			// saves it in the strings builder
			entry.WriteRune(r)

			if entry.String() == code {
				state = Unlocked
				break
			}

			// change state in failed
			if strings.Index(code, entry.String()) != 0 {
				// code is wrong
				state = Failed
			}
		case Failed:
			fmt.Println("FAILED")
			// resets, start again
			entry.Reset()
			// locked again
			state = Locked
		case Unlocked:
			fmt.Println("UNLOCKED")
			// finish condition
			return
		}
	}
}
