package main

import (
	"advent_of_code_q6/processor"
	"fmt"
)

func main() {
	ep, err := processor.NewElfPacket("input.txt", 4, 14)
	if err != nil {
		panic(err)
	}
	startIndex := ep.GetStartSequenceIndex()
	fmt.Printf("The start sequence index is %d\n", startIndex)

	startMessageIndex := ep.GetStartMessageIndex()
	fmt.Printf("The start message index is %d\n", startMessageIndex)
}
