package drone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newGrid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected grid
	}{
		{
			name: "Test Case 1",
			input: `30373
25512
65332
33549
35390`,
			expected: grid{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 0},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g, err := newGrid(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, g)
		})
	}
}

func Test_GetSumOfVisibleTreeHeights(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Test Case 1",
			input: `30373
25512
65332
33549
35390`,
			expected: 21,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			g, err := newGrid(tc.input)
			assert.NoError(t, err)

			d := dimensions{
				x: len(g[0]),
				y: len(g),
			}

			tp := TreeProcessor{
				treeGrid:           g,
				treeGridDimensions: d,
			}
			actual := tp.GetNumberOfVisibleTrees()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_GetBestVisibilityScore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Test Case 1",
			input: `30373
25512
65332
33549
35390`,
			expected: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			g, err := newGrid(tc.input)
			assert.NoError(t, err)

			d := dimensions{
				x: len(g[0]),
				y: len(g),
			}

			tp := TreeProcessor{
				treeGrid:           g,
				treeGridDimensions: d,
			}
			actual := tp.GetBestVisibilityScore()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
