package main

import (
	"advent_of_code_q4/cleanup"
	"fmt"
)

func main() {
	a, err := cleanup.NewAssigner("input.txt")
	if err != nil {
		panic(err)
	}
	c, err := a.CountAssignmentOverlaps()
	if err != nil {
		panic(err)
	}
	intersections, err := a.CountAssignmentIntersections()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The number of overlapping pairs is: %d\n", c)
	fmt.Printf("The number of intersecting pairs is: %d\n", intersections)
}
