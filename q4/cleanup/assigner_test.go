package cleanup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountAssignmentOverlaps(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name: "elf-data",
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			expectedOutput: 2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := Assigner{
				assignmentList: []byte(tc.input),
			}
			actual, err := a.CountAssignmentOverlaps()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_CountAssignmentIntersections(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name: "elf-data",
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			expectedOutput: 4,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := Assigner{
				assignmentList: []byte(tc.input),
			}
			actual, err := a.CountAssignmentIntersections()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_checkFullyContainsPair(t *testing.T) {
	tests := []struct {
		name           string
		a              []int
		b              []int
		expectedOutput int
	}{
		{
			name:           "a contains b",
			a:              []int{1, 6},
			b:              []int{2, 5},
			expectedOutput: 1,
		},
		{
			name:           "b contains a",
			a:              []int{2, 5},
			b:              []int{1, 6},
			expectedOutput: 1,
		},
		{
			name:           "a equals b",
			a:              []int{2, 5},
			b:              []int{1, 6},
			expectedOutput: 1,
		},
		{
			name:           "a does not contain b",
			a:              []int{2, 5},
			b:              []int{6, 9},
			expectedOutput: 0,
		},
		{
			name:           "b does not contain a",
			a:              []int{6, 9},
			b:              []int{2, 5},
			expectedOutput: 0,
		},
		{
			name:           "a does not fully contain b",
			a:              []int{6, 9},
			b:              []int{8, 10},
			expectedOutput: 0,
		},
		{
			name:           "b does not fully contain a",
			a:              []int{8, 10},
			b:              []int{6, 9},
			expectedOutput: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := checkFullyContainsPairs(tc.a[0], tc.a[1], tc.b[0], tc.b[1])
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_parsePairs(t *testing.T) {
	input := "2-4,6-8"
	expected := []assignment{
		{
			start:  2,
			finish: 4,
		},
		{
			start:  6,
			finish: 8,
		},
	}
	actual, err := parsePairs(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_parsePairsToLists(t *testing.T) {
	input := "2-4,6-8"
	expected := []assignmentList{
		{
			2, 3, 4,
		},
		{
			6, 7, 8,
		},
	}
	actual, err := parsePairsToLists(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_checkContainsPairs(t *testing.T) {
	tests := []struct {
		name           string
		a              []int
		b              []int
		expectedOutput int
	}{
		{
			name:           "a contains b",
			a:              []int{1, 2, 3, 4, 5, 6},
			b:              []int{2, 3, 4, 5},
			expectedOutput: 1,
		},
		{
			name:           "b contains a",
			a:              []int{2, 3, 4, 5},
			b:              []int{1, 2, 3, 4, 5, 6},
			expectedOutput: 1,
		},
		{
			name:           "a equals b",
			a:              []int{2, 3, 4, 5},
			b:              []int{1, 2, 3, 4, 5, 6},
			expectedOutput: 1,
		},
		{
			name:           "a does not contain b",
			a:              []int{2, 3, 4, 5},
			b:              []int{6, 7, 8, 9},
			expectedOutput: 0,
		},
		{
			name:           "b does not contain a",
			a:              []int{6, 7, 8, 9},
			b:              []int{2, 3, 4, 5},
			expectedOutput: 0,
		},
		{
			name:           "a does not fully contain b",
			a:              []int{6, 7, 8, 9},
			b:              []int{8, 9, 10},
			expectedOutput: 1,
		},
		{
			name:           "b does not fully contain a",
			a:              []int{8, 9, 10},
			b:              []int{6, 7, 8, 9},
			expectedOutput: 1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := checkContainsPair(tc.a, tc.b)
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}
