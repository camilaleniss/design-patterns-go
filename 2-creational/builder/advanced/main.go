package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person // needs to be inited
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}} // initialize the person
}

// for building the entire person
func (it *PersonBuilder) Build() *Person {
	return it.person
}

type PersonJobBuilder struct {
	// aggregates the person builder
	// so it has access to person too
	PersonBuilder
}

// returns the job builder
func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}

// methods to set up the job
func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

// aggregates the person builder
// so it has access to person too
type PersonAddressBuilder struct {
	PersonBuilder
}

// method to set up the location
func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}

func (it *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(
	postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)

	person := pb.Build() //sum all together
	fmt.Println(*person) // initialize all params :D
}
