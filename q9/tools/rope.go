package tools

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

type RopeMovement struct {
	instructions string
	board        board
}

func newMovement(fileName string) (RopeMovement, error) {
	in, err := os.ReadFile(fileName)
	if err != nil {
		return RopeMovement{}, err
	}

	b, err := newBoard(string(in))
	if err != nil {
		return RopeMovement{}, err
	}
	b.markTailVisits(string(in))
	return RopeMovement{
		instructions: string(in),
		board:        b,
	}, nil
}

func (rm RopeMovement) GetNumberTailVisits() int {
	out := 0
	for j, row := range rm.board {
		for i, _ := range row {
			if rm.board[j][i] >= 1 {
				out++
			}
		}
	}

	return out
}

type pos struct {
	j int
	i int
}

type board [][]int

func (b board) markTailVisits(s string) error {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(s)))
	posHead := pos{
		j: len(b) - 1,
		i: 0,
	}
	posTail := pos{
		j: len(b) - 1,
		i: 0,
	}
	b[posTail.j][posTail.i] = 1
	for scanner.Scan() {
		ss := strings.Fields(scanner.Text())
		instruction := ss[0]
		movement, err := strconv.Atoi(ss[1])
		if err != nil {
			return err
		}

		switch instruction {
		case "R":
			for i := 0; i < movement; i++ {
				// Move head
				posHead.i++

				// Move tail
				posTail = determineTailMovement(posHead, posTail)

				// Mark position
				b[posTail.j][posTail.i]++
			}
		case "L":
			for i := 0; i < movement; i++ {
				// Move head
				posHead.i--

				// Move tail
				posTail = determineTailMovement(posHead, posTail)

				// Mark position
				b[posTail.j][posTail.i]++
			}
		case "U":
			for j := 0; j < movement; j++ {
				// Move head
				posHead.j--

				// Move tail
				posTail = determineTailMovement(posHead, posTail)

				// Mark position
				b[posTail.j][posTail.i]++
			}
		case "D":
			for j := 0; j < movement; j++ {
				// Move head
				posHead.j++

				// Move tail
				posTail = determineTailMovement(posHead, posTail)

				// Mark position
				b[posTail.j][posTail.i]++
			}
		}
	}

	return nil
}

func determineTailMovement(h, t pos) pos {
	// Check horizontal
	if (h.j == t.j && (h.i-t.i) == 1) ||
		(h.j == t.j && (h.i-t.i) == -1) {
		return t
	}

	// Check vertical
	if (h.i == t.i && (h.j-t.j) == 1) ||
		(h.i == t.i && (h.j-t.j) == -1) {
		return t
	}

	// Check Diagonal
	if (h.j-t.j == 1 && h.i-t.i == 1) ||
		(h.j-t.j == 1 && h.i-t.i == -1) ||
		(h.j-t.j == -1 && h.i-t.i == -1) ||
		(h.j-t.j == -1 && h.i-t.i == 1) {
		return t
	}

	// Check Same
	if h == t {
		return t
	}

	if h.i == t.i {
		// t.j above
		if t.j < h.j {
			t.j = h.j - 1
		} else {
			// t.j below
			t.j = h.j + 1
		}
	} else if h.j == t.j {
		// t.i right of h.i
		if t.i < h.i {
			t.i = h.i - 1
		} else {
			t.i = h.i + 1
		}
	} else {
		t.i = h.i
		// t.j is above h
		if t.j < h.j {
			t.j = h.j - 1
		} else {
			t.j = h.j + 1
		}
	}

	return t
}

func newBoard(s string) (board, error) {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(s)))
	out := make(board, 0)
	currPos := pos{
		j: 0,
		i: 0,
	}
	out = append(out, make([]int, 0))
	for scanner.Scan() {
		ss := strings.Fields(scanner.Text())
		instruction := ss[0]
		movement, err := strconv.Atoi(ss[1])
		if err != nil {
			return nil, err
		}

		switch instruction {
		case "R":
			currPos.i = currPos.i + movement

			// Add columns to the right
			if currPos.i > len(out[currPos.j])-1 {
				add := make([]int, currPos.i-len(out[currPos.j])+1)
				for j, _ := range out {
					out[j] = append(out[j], add...)
				}
			}
		case "L":
			currPos.i = currPos.i - movement

			// Add columns to the left
			if currPos.i < 0 {
				add := make([]int, -currPos.i+1)
				for j, _ := range out {
					out[j] = append(add, out[j]...)
				}
				currPos.i = 0
			}
		case "U":
			prevPos := currPos.j
			currPos.j = currPos.j - movement
			// Add row to the top
			if currPos.j < 0 {
				add := make(board, -currPos.j)
				for j, _ := range add {
					add[j] = append(add[j], make([]int, len(out[prevPos]))...)
				}
				out = append(add, out...)
				currPos.j = 0
			}
		case "D":
			prevPos := currPos.j
			currPos.j = currPos.j + movement

			// Add row to the bottom
			if currPos.j > len(out)-1 {
				add := make(board, currPos.j-len(out)+1)
				for j, _ := range add {
					add[j] = append(add[j], make([]int, len(out[prevPos]))...)
				}
				out = append(out, add...)
				currPos.j = 0
			}
		}
	}

	return out, nil
}
