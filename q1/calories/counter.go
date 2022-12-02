package calories

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

// CalorieList holds the total calories
type CalorieList struct {
	totals []int
}

// NewCalorieList represents a list of total calories
func NewCalorieList(inputPath string) (CalorieList, error) {
	b, err := os.ReadFile(inputPath)
	if err != nil {
		return CalorieList{}, err
	}

	totals := convertToSliceOfTotals(b)

	return CalorieList{
		totals: totals,
	}, nil
}

// CalculateHighestCalories calculates the highest total calories in the list
func (c CalorieList) CalculateHighestCalories() int {
	sort.Ints(c.totals)
	return c.totals[len(c.totals)-1]
}

// CalculateTopThreeHighestCalories calculates the sum of the top 3 highest total calories
func (c CalorieList) CalculateTopThreeHighestCalories() int {
	sort.Ints(c.totals)

	out := 0
	sub := c.totals[len(c.totals)-3:]
	for _, v := range sub {
		out = out + v
	}

	return out
}

func convertToSliceOfTotals(data []byte) []int {
	input := string(data)

	// Split on the empty lines
	split := strings.Split(input, "\n\n")

	// Now they're grouped by the blank newline, split on the remaining newlines
	list := make([][]string, len(split))
	ls := make([][]int, len(split))
	for i, s := range split {
		t := strings.Split(s, "\n")

		// Deal with leading/lagging newlines
		if t[0] == "" {
			list[i] = t[1:]
		} else if t[len(t)-1] == "" {
			list[i] = t[:len(t)-1]
		} else {
			list[i] = strings.Split(s, "\n")
		}

		for _, s2 := range list[i] {
			n, _ := strconv.Atoi(s2)
			ls[i] = append(ls[i], n)
		}
	}

	out := make([]int, len(ls))
	for i, sn := range ls {
		for _, sv := range sn {
			out[i] = out[i] + sv
		}
	}

	return out
}
