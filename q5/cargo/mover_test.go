package cargo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newContainers(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	container := newContainers(input)
	fmt.Println(container)

	assert.Equal(t, 3, len(container))
	assert.Equal(t, "[N]", container[1].Back().Value)
	assert.Equal(t, "[D]", container[2].Back().Value)
	assert.Equal(t, "[P]", container[3].Back().Value)
}

func Test_parseOperatorGuide(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1`
	expectedMovePlan := []movePlan{
		{
			numCratesToMove: 1,
			moveFrom:        2,
			moveTo:          1,
		},
	}

	container, mp, err := parseOperatorGuide(input)
	assert.NoError(t, err)

	assert.Equal(t, 3, len(container))
	assert.Equal(t, "[N]", container[1].Back().Value)
	assert.Equal(t, "[D]", container[2].Back().Value)
	assert.Equal(t, "[P]", container[3].Back().Value)

	assert.Equal(t, expectedMovePlan, mp)
}

func Test_newMovePlans(t *testing.T) {
	input := `move 1 from 2 to 1
move 3 from 1 to 3`
	expectedMovePlan := movePlans{
		{
			numCratesToMove: 1,
			moveFrom:        2,
			moveTo:          1,
		},
		{
			numCratesToMove: 3,
			moveFrom:        1,
			moveTo:          3,
		},
	}

	mp, err := newMovePlans(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedMovePlan, mp)
}

func Test_PerformMovements(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name: "elf-data",
			input: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`,
			expectedOutput: "CMZ",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			c, mp, err := parseOperatorGuide(tc.input)
			assert.NoError(t, err)

			mover := Mover{
				containers: c,
				movePlan:   mp,
			}

			actual := mover.PerformMovements()

			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}

func Test_PerformMovementsBulk(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name: "elf-data",
			input: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`,
			expectedOutput: "MCD",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			c, mp, err := parseOperatorGuide(tc.input)
			assert.NoError(t, err)

			mover := Mover{
				containers: c,
				movePlan:   mp,
			}

			actual := mover.PerformMovementsBulk()

			assert.Equal(t, tc.expectedOutput, actual)
		})
	}
}
