package main

import (
	"advent_of_code_q1/calories"
	"fmt"
)

// List of calories, each elf separates the calories by a blank line
// eg.
/*
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
*/
//This list represents the Calories of the food carried by five Elves:
//
//The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
//The second Elf is carrying one food item with 4000 Calories.
//The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
//The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
//The fifth Elf is carrying one food item with 10000 Calories.

// How many total calories is the elf carrying with the most calories

func main() {
	cl, err := calories.NewCalorieList("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("The highest calories held by an elf is %d\n", cl.CalculateHighestCalories())
	fmt.Printf("The sum of the top three highest calories held by the elves is  %d\n", cl.CalculateTopThreeHighestCalories())
	return
}
