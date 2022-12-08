package drone

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strconv"
)

type grid [][]int

func newGrid(s string) (grid, error) {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(s)))
	out := make(grid, 0)
	i := 0
	for scanner.Scan() {
		out = append(out, make([]int, 0))
		for _, r := range scanner.Text() {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			out[i] = append(out[i], v)
		}
		i++
	}

	return out, nil
}

type dimensions struct {
	x int
	y int
}

type TreeProcessor struct {
	treeGrid           grid
	treeGridDimensions dimensions
}

func NewTreeProcessor(fileName string) (TreeProcessor, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return TreeProcessor{}, err
	}

	g, err := newGrid(string(b))
	if err != nil {
		return TreeProcessor{}, err
	}

	d := dimensions{
		x: len(g[0]),
		y: len(g),
	}

	return TreeProcessor{
		treeGrid:           g,
		treeGridDimensions: d,
	}, nil
}

func (tp TreeProcessor) GetNumberOfVisibleTrees() int {
	return len(tp.getVisibleTreeHeights())
}

func (tp TreeProcessor) GetBestVisibilityScore() int {
	scores := tp.getVisibilityScores()
	sort.Ints(scores)
	return scores[len(scores)-1]
}

func (tp TreeProcessor) getVisibleTreeHeights() []int {
	out := make([]int, 0)
	for i, _ := range tp.treeGrid {
		for j, _ := range tp.treeGrid[i] {
			if v, ok := tp.checkVisibility(j, i); ok {
				out = append(out, v)
			}
		}
	}

	return out
}

func (tp TreeProcessor) getVisibilityScores() []int {
	// If i is vertical, and j is horizontal
	// Everything on the outside is visible, so anything at position [0,j] or [i,0] is automatically visible
	out := make([]int, 0)
	for i, _ := range tp.treeGrid {
		for j, _ := range tp.treeGrid[i] {
			v := tp.getNumVisibleTrees(j, i)
			score := v[0] * v[1] * v[2] * v[3]
			out = append(out, score)
		}
	}

	return out
}

func (tp TreeProcessor) checkVisibility(x, y int) (int, bool) {
	g := tp.treeGrid

	// If trees are on edges, they are automatically visible
	if x == 0 || y == 0 || x == tp.treeGridDimensions.x-1 || y == tp.treeGridDimensions.y-1 {
		return g[y][x], true
	}

	// Check left
	isHigher := true
	for i := x - 1; i >= 0; i-- {
		if g[y][x] <= g[y][i] {
			isHigher = false
			break
		}
	}
	if isHigher {
		return g[y][x], true
	}

	// Check right
	isHigher = true
	for i := x + 1; i < tp.treeGridDimensions.x; i++ {
		if g[y][x] <= g[y][i] {
			isHigher = false
			break
		}
	}
	if isHigher {
		return g[y][x], true
	}

	// Check up
	isHigher = true
	for j := y - 1; j >= 0; j-- {
		if g[y][x] <= g[j][x] {
			isHigher = false
			break
		}
	}
	if isHigher {
		return g[y][x], true
	}

	// Check down
	isHigher = true
	for j := y + 1; j < tp.treeGridDimensions.y; j++ {
		if g[y][x] <= g[j][x] {
			isHigher = false
			break
		}
	}
	if isHigher {
		return g[y][x], true
	}

	return 0, false
}

func (tp TreeProcessor) getNumVisibleTrees(x, y int) []int {
	g := tp.treeGrid

	// Check left
	out := make([]int, 4)

	for i := x - 1; i >= 0; i-- {
		out[0]++
		if g[y][x] <= g[y][i] {
			break
		}
	}

	// Check right
	for i := x + 1; i < tp.treeGridDimensions.x; i++ {
		out[1]++
		if g[y][x] <= g[y][i] {
			break
		}
	}

	// Check up
	for j := y - 1; j >= 0; j-- {
		out[2]++
		if g[y][x] <= g[j][x] {
			break
		}
	}

	// Check down
	for j := y + 1; j < tp.treeGridDimensions.y; j++ {
		out[3]++
		if g[y][x] <= g[j][x] {
			break
		}
	}

	return out
}
