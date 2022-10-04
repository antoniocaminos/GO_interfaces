package main

import "fmt"

//interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

//Dog is a type of Dog
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

//say is required t the animal interface
func (d *Dog) Says() string {
	return d.Sound
}

//howmanylegs isrequired to the animal interface
func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

//Cats is the type for Cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

//Says is required to the animal interface
func (c *Cat) Says() string {
	return c.Sound
}

//howmanylegs is required to the animal interface
func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}
func main() {
	//ask a riddle
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	Riddle(&dog)

	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`this animal say %s and has %d legs. what animal is it `, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
