package rucksack

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateDuplicatePriorityTotal(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name: "elf-data",
			input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`,
			expectedOutput: 157,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := Optimizer{
				rucksackContents: []byte(tc.input),
			}
			actual := o.CalculateDuplicatePriorityTotal()
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_CalculatePriorityOfBadges(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name: "elf-data",
			input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`,
			expectedOutput: 70,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := Optimizer{
				rucksackContents: []byte(tc.input),
			}
			actual := o.CalculatePriorityOfBadges()
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_decodeRune(t *testing.T) {
	tests := []struct {
		name           string
		input          rune
		expectedOutput int
	}{
		{
			name:           "A",
			input:          'A',
			expectedOutput: 27,
		},
		{
			name:           "B",
			input:          'B',
			expectedOutput: 28,
		},
		{
			name:           "Z",
			input:          'Z',
			expectedOutput: 52,
		},
		{
			name:           "a",
			input:          'a',
			expectedOutput: 1,
		},
		{
			name:           "b",
			input:          'b',
			expectedOutput: 2,
		},
		{
			name:           "z",
			input:          'z',
			expectedOutput: 26,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := decodeRunePriority(tc.input)
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_stringSplitter(t *testing.T) {
	a, b := splitItems("aBcDeZfGrD")
	assert.Equal(t, "aBcDe", a)
	assert.Equal(t, "ZfGrD", b)
}

func Test_findDuplicates(t *testing.T) {
	s1 := "atzbck"
	s2 := "fbcmzlKb"
	expected := []rune{'b', 'z', 'c'}
	actual := findDuplicates(s1, s2)

	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})
	sort.Slice(actual, func(i, j int) bool {
		return actual[i] < actual[j]
	})

	assert.Equal(t, expected, actual)
}

func Test_findSharedDuplicates(t *testing.T) {
	s1 := "abzcde"
	s2 := "fgzhij"
	s3 := "klzmno"
	expected := 'z'
	actual := findSharedDuplicate(s1, s2, s3)
	assert.Equal(t, expected, actual)
}
