package main

import (
	"fmt"
	"strings"
)

// example for a Game where you can have +1M of users with similar names
type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

// more efficient user
type User2 struct {
	names []uint8
}

// uploading to the slice
func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {
	// without using flyweight
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsoJane := NewUser("Jane Smith")

	fmt.Println(john.FullName)
	fmt.Println(jane.FullName)
	fmt.Println(alsoJane.FullName)

	fmt.Println("Memory taken by users:",
		len([]byte(john.FullName))+
			len([]byte(alsoJane.FullName))+
			len([]byte(jane.FullName)))

	john2 := NewUser2("John Doe")
	jane2 := NewUser2("Jane Doe")
	alsoJane2 := NewUser2("Jane Smith")

	fmt.Println(john2.FullName())
	fmt.Println(jane2.FullName())
	fmt.Println(alsoJane2.FullName())

	totalMem := 0
	for _, a := range allNames {
		totalMem += len([]byte(a))
	}
	totalMem += len(john2.names)
	totalMem += len(jane2.names)
	totalMem += len(alsoJane2.names)

	// flywight!
	fmt.Println("Memory taken by users2:", totalMem)
}
