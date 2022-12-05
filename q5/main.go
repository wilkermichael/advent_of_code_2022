package main

import (
	"advent_of_code_q5/cargo"
	"fmt"
)

func main() {
	m, err := cargo.NewMover("input.txt")
	if err != nil {
		panic(err)
	}
	s := m.PerformMovements()
	fmt.Printf("The movements are: %s\n", s)

	// Playing with linked lists so need to create a brand new object
	m2, err := cargo.NewMover("input.txt")
	s = m2.PerformMovementsBulk()
	fmt.Printf("The bulk movements are: %s", s)
}
