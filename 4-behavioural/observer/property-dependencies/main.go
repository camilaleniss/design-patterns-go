package propertydependencies

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{Observable{new(list.List)}, age}
}

type PropertyChanged struct {
	Name  string
	Value interface{}
}

// the fact that a person can vote depends on the age
func (p *Person) CanVote() bool {
	return p.age >= 18
}

func (p *Person) Age() int { return p.age }

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	} // no change

	// record the previous state
	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChanged{"Age", p.age})

	// review if change
	if oldCanVote != p.CanVote() {
		// notify them all
		p.Fire(PropertyChanged{"CanVote", p.CanVote()})
	}
}

type ElectrocalRoll struct {
}

func (e *ElectrocalRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChanged); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectrocalRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
