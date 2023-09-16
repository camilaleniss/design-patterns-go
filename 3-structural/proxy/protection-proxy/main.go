package main

import "fmt"

// base interface
type Driven interface {
	Drive()
}

type Car struct{}

// implements the interface
func (c *Car) Drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

// this can also be called 'secureCar' or something like that
type CarProxy struct {
	// notice the private of the fields
	car    Car
	driver *Driver
}

// also implements the interface
func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

// constructor
func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}
