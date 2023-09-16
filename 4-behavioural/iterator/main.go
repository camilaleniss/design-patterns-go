package iterator

import "fmt"

// Iterator
// Core functionality of data structures
// Also called transversal

// Keeps a pointer to the current element
// Knows how to move to a different element

// Object that facilities the transversal for a data structure

// Also it's kinda built it in Go when using range()

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
	return []string{p.FirstName, p.MiddleName, p.LastName}
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}
}
