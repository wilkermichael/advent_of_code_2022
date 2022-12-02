package calories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalorieList_CalculateHighestCalories(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{
			name: "elf-example",
			input: `
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
`,
			result: 24000,
		},
		{
			name: "single-val",
			input: `
10000
`,
			result: 10000,
		},
		{
			name: "multi-single-val",
			input: `
12000

10000
`,
			result: 12000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			totals := convertToSliceOfTotals([]byte(tc.input))
			c := CalorieList{
				totals: totals,
			}
			s := c.CalculateHighestCalories()
			assert.Equal(t, tc.result, s)
		})
	}
}

func Test_CalorieList_CalculateTopThreeHighestCalories(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{
			name: "elf-example",
			input: `
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
`,
			result: 45000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			totals := convertToSliceOfTotals([]byte(tc.input))
			c := CalorieList{
				totals: totals,
			}
			s := c.CalculateTopThreeHighestCalories()
			assert.Equal(t, tc.result, s)
		})
	}
}
