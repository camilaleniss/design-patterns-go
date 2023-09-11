package factories

// Factories
// When the struct has too many fields and
// you want to initialize them correctly.
// Not each time you want to create them.
// Wholesale object creation != builder (piecewise)

import "fmt"

type Person struct {
	Name string
	Age  int
}

// factory function: just a single function handles all the creation
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func main_() {
	// initialize directly
	p := Person{"John", 22}
	fmt.Println(p)

	// use a constructor
	// this gives us the ability to check values and do some business logic
	p2 := NewPerson("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)

	// also we can use interfaces as the object returned in the constructor
	// just for tha sake of encapsulating the actual Person
	// and also allows us to use other underlying objects
}
