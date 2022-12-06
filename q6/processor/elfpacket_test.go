package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStartSequenceIndex(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
		numUnique      int
	}{
		{
			name:           "input 1 4 unique",
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expectedOutput: 7,
			numUnique:      4,
		},
		{
			name:           "input 2 4 unique",
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expectedOutput: 5,
			numUnique:      4,
		},
		{
			name:           "input 3 4 unique",
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			expectedOutput: 6,
			numUnique:      4,
		},
		{
			name:           "input 4 4 unique",
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expectedOutput: 10,
			numUnique:      4,
		},
		{
			name:           "input 5 4 unique",
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expectedOutput: 11,
			numUnique:      4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ep := ElfPacket{
				packet:               []byte(tc.input),
				startPacketNumUnique: tc.numUnique,
			}

			actual := ep.GetStartSequenceIndex()
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_GetStartMessageIndex(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
		numUnique      int
	}{
		{
			name:           "input 1 14 unique",
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expectedOutput: 19,
			numUnique:      14,
		},
		{
			name:           "input 2 14 unique",
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expectedOutput: 23,
			numUnique:      14,
		},
		{
			name:           "input 3 14 unique",
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			expectedOutput: 23,
			numUnique:      14,
		},
		{
			name:           "input 4 14 unique",
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expectedOutput: 29,
			numUnique:      14,
		},
		{
			name:           "input 5 14 unique",
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expectedOutput: 26,
			numUnique:      14,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ep := ElfPacket{
				packet:               []byte(tc.input),
				startPacketNumUnique: 0,
				messageNumUnique:     tc.numUnique,
			}

			actual := ep.GetStartMessageIndex()
			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}
