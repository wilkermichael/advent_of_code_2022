package main

import (
	"advent_of_code_q2/rps"
	"fmt"
)

func main() {
	s, err := rps.NewStrategy("input.txt", rps.DefaultScoreConfig)
	if err != nil {
		panic(err)
	}

	s.SetDefaultEncoding()
	fmt.Printf("The total score with oponent player r p s mappings is: %d\n", s.CalculateTotalScore())

	s.SetDefaultEncodingWithExpectedResults()
	fmt.Printf("The total score with suggested results is : %d\n", s.CalculateTotalScoreWithSuggestedMove())
}
