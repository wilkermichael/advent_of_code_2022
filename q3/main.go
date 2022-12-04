package main

import (
	"advent_of_code_q3/rucksack"
	"fmt"
)

func main() {
	o, err := rucksack.NewOptimizer("input.txt")
	if err != nil {
		panic(err)
	}
	priorityTotal := o.CalculateDuplicatePriorityTotal()
	badgePriorityTotal := o.CalculatePriorityOfBadges()
	fmt.Printf("The priority total for duplicates is %d\n", priorityTotal)
	fmt.Printf("The priority total for badges is %d\n", badgePriorityTotal)
}
