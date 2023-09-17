package observer

// Observer
// When we need to be informed when certain things happen
// Kinda a trigger

// We want to listen to events and notified when they occur
// We have an observer and an observable

import (
	"container/list"
	"fmt"
)

// the object that will have the subscribers and will notify them
type Observable struct {
	subs *list.List
}

// provide methods to subscribe and unsubscribe as well
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

// notify method for all the subscribers
func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

// the object that will be notified
type Observer interface {
	Notify(data interface{})
}

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	// notify them all!
	p.Fire(p.Name)
}

type DoctorService struct{}

// actually what happens when it gets notified
func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}

func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	// let's test it!
	p.CatchACold()
}
