package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategy_EncodeRock(t *testing.T) {
	s := Strategy{
		code: make(map[string]string),
	}
	input1 := "x"
	input2 := "y"
	s.EncodeRock(input1)
	s.EncodeRock(input2)
	assert.Equal(t, s.code[input1], rockType)
	assert.Equal(t, s.code[input2], rockType)
}

func TestStrategy_EncodePaper(t *testing.T) {
	s := Strategy{
		code: make(map[string]string),
	}
	input1 := "x"
	input2 := "y"
	s.EncodePaper(input1)
	s.EncodePaper(input2)
	assert.Equal(t, s.code[input1], paperType)
	assert.Equal(t, s.code[input2], paperType)
}

func TestStrategy_EncodeScissors(t *testing.T) {
	s := Strategy{
		code: make(map[string]string),
	}
	input1 := "x"
	input2 := "y"
	s.EncodeScissors(input1)
	s.EncodeScissors(input2)
	assert.Equal(t, s.code[input1], scissorsType)
	assert.Equal(t, s.code[input2], scissorsType)
}

func TestStrategy_CalculateTotalScore(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		rock        []string
		paper       []string
		scissors    []string
		scoreConfig ScoreConfig
		result      int
	}{
		{
			name: "elf-example",
			input: `A Y
B X
C Z
`,
			rock:     []string{"A", "X"},
			paper:    []string{"B", "Y"},
			scissors: []string{"C", "Z"},
			scoreConfig: ScoreConfig{
				Rock:     1,
				Paper:    2,
				Scissors: 3,
				Loss:     0,
				Draw:     3,
				Win:      6,
			},
			result: 15,
		},
		{
			name: "elf-example-defaults",
			input: `A Y
B X
C Z
`,
			result: 15,
		},
		{
			name: "elf-example-leading space",
			input: `
A Y
B X
C Z
`,
			result: 15,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			strategy := Strategy{
				code:     make(map[string]string),
				solution: []byte(tc.input),
				scoring:  make(map[string]int),
			}

			if tc.scoreConfig == (ScoreConfig{}) {
				strategy.SetScoring(DefaultScoreConfig)
			} else {
				strategy.SetScoring(tc.scoreConfig)
			}

			if len(tc.rock) == 0 && len(tc.paper) == 0 && len(tc.scissors) == 0 {
				strategy.SetDefaultEncoding()
			} else {
				for _, v := range tc.rock {
					strategy.EncodeRock(v)
				}
				for _, v := range tc.scissors {
					strategy.EncodeScissors(v)
				}
				for _, v := range tc.paper {
					strategy.EncodePaper(v)
				}
			}

			actual := strategy.CalculateTotalScore()
			assert.Equal(t, tc.result, actual)
		})
	}
}

func TestStrategy_CalculateTotalScoreWithSuggestedMove(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{
			name: "elf-example-defaults",
			input: `A Y
B X
C Z
`,
			result: 12,
		},
		{
			name: "elf-example-leading space",
			input: `
A Y
B X
C Z
`,
			result: 12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			strategy := Strategy{
				code:     make(map[string]string),
				solution: []byte(tc.input),
				scoring:  make(map[string]int),
			}

			strategy.SetScoring(DefaultScoreConfig)
			strategy.SetDefaultEncodingWithExpectedResults()

			actual := strategy.CalculateTotalScoreWithSuggestedMove()
			assert.Equal(t, tc.result, actual)
		})
	}
}
