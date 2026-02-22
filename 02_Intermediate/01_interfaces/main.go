package main

import "fmt"

type Speaker interface {
	Speak() string
}
type Dog struct {
}
type Cat struct {
}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

// In Go, since Dog and Cat both have a 'Speak() string' method,
// they AUTOMATICALLY fulfill the 'Speaker' interface without you having to announce it!
func MakeSound(s Speaker) {
	fmt.Printf("The Type (%T) Sounds: %v\n", s, s.Speak())
}
func main() {
	// In Go, we initialize a struct using curly braces {}, not parentheses ()
	dog := Dog{}
	cat := Cat{}

	MakeSound(dog)
	MakeSound(cat)
}
