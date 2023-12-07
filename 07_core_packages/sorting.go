package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type Persons []Person

func (this Persons) Len() int {
	return len(this)
}

func (this Persons) Less(i, j int) bool {
	return this[i].name < this[j].name
}

func (this Persons) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type PersonsByAgeInverted []Person

func (this PersonsByAgeInverted) Len() int {
	return len(this)
}

func (this PersonsByAgeInverted) Less(i, j int) bool {
	return this[i].age > this[j].age
}

func (this PersonsByAgeInverted) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	whatever := Persons{
		{"test", 20},
		{"teen", 15},
		{"oskol", 25},
		{"ahmaq", 10},
	}

	sort.Sort(whatever)
	fmt.Println(whatever)

	sort.Sort(PersonsByAgeInverted(whatever))

	fmt.Println(whatever)

	test2 := []Person{
		{"t1", 10},
		{"t2", 8},
		{"a3", 25},
		{"a1", 10},
	}

	sort.Sort(PersonsByAgeInverted(test2))
	fmt.Println(test2)

	sort.Sort(Persons(test2))

	fmt.Println(test2)
}
