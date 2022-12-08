package main

import (
	"advent_of_code_q8/drone"
	"fmt"
)

func main() {
	d, err := drone.NewTreeProcessor("input.txt")
	if err != nil {
		panic(err)
	}
	numVisibleTrees := d.GetNumberOfVisibleTrees()
	fmt.Printf("The number of visible trees is %d\n", numVisibleTrees)

	bestVisibilityScore := d.GetBestVisibilityScore()
	fmt.Printf("The best visibility score is %d\n", bestVisibilityScore)
}
