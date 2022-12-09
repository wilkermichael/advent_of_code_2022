package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createBoard(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedSizeJ int
		expectedSizeI int
	}{
		{
			name: "case 1",
			input: `R 4
     U 4
     L 3
     D 1
     R 4
     D 1
     L 5
     R 2`,
			expectedSizeJ: 5,
			expectedSizeI: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := newBoard(tc.input)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedSizeJ, len(actual))
			assert.Equal(t, tc.expectedSizeI, len(actual[0]))
		})
	}
}

func Test_GetNumberTailVisits(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "case 1",
			input: `R 4
     U 4
     L 3
     D 1
     R 4
     D 1
     L 5
     R 2`,
			expected: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := newBoard(tc.input)
			assert.NoError(t, err)
			b.markTailVisits(tc.input)
			rm := RopeMovement{
				instructions: tc.input,
				board:        b,
			}

			actual := rm.GetNumberTailVisits()
			fmt.Println(rm.board)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_determineTailMovement(t *testing.T) {
	testCases := []struct {
		name         string
		tail         pos
		head         pos
		expectedTail pos
	}{
		{
			name:         "horizontal left connected",
			tail:         pos{2, 1},
			head:         pos{2, 2},
			expectedTail: pos{2, 1},
		},
		{
			name:         "horizontal right connected",
			tail:         pos{2, 2},
			head:         pos{2, 1},
			expectedTail: pos{2, 2},
		},
		{
			name:         "vertical top connected",
			tail:         pos{1, 1},
			head:         pos{2, 1},
			expectedTail: pos{1, 1},
		},
		{
			name:         "vertical bottom connected",
			tail:         pos{2, 1},
			head:         pos{1, 1},
			expectedTail: pos{2, 1},
		},
		{
			name:         "vertical top left connected",
			tail:         pos{1, 0},
			head:         pos{2, 1},
			expectedTail: pos{1, 0},
		},
		{
			name:         "vertical top right connected",
			tail:         pos{1, 2},
			head:         pos{2, 1},
			expectedTail: pos{1, 2},
		},
		{
			name:         "vertical bottom right connected",
			tail:         pos{3, 2},
			head:         pos{2, 1},
			expectedTail: pos{3, 2},
		},
		{
			name:         "vertical bottom left connected",
			tail:         pos{3, 0},
			head:         pos{2, 1},
			expectedTail: pos{3, 0},
		},
		{
			name:         "vertical bottom left not connected",
			tail:         pos{2, 1},
			head:         pos{0, 2},
			expectedTail: pos{1, 2},
		},
		{
			name:         "vertical bottom right not connected",
			tail:         pos{2, 3},
			head:         pos{0, 2},
			expectedTail: pos{1, 2},
		},
		{
			name:         "vertical top left not connected",
			tail:         pos{0, 0},
			head:         pos{2, 1},
			expectedTail: pos{1, 1},
		},
		{
			name:         "vertical top right not connected",
			tail:         pos{0, 2},
			head:         pos{2, 1},
			expectedTail: pos{1, 1},
		},
		{
			name:         "same",
			tail:         pos{0, 2},
			head:         pos{0, 2},
			expectedTail: pos{0, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := determineTailMovement(tc.head, tc.tail)

			assert.Equal(t, tc.expectedTail, actual)
		})
	}
}
