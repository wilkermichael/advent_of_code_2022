package cleanup

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	start  int
	finish int
}

type assignmentList []int

type Assigner struct {
	assignmentList []byte
}

func NewAssigner(fileName string) (Assigner, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return Assigner{}, err
	}
	return Assigner{
		assignmentList: b,
	}, nil
}

func (a Assigner) CountAssignmentOverlaps() (int, error) {
	r := bytes.NewReader(a.assignmentList)
	scanner := bufio.NewScanner(r)
	out := 0
	for scanner.Scan() {
		// Parse the pairs
		assignments, err := parsePairs(scanner.Text())
		if err != nil {
			return 0, err
		}
		// Check for overlap
		out = out + checkFullyContainsPairs(assignments[0].start, assignments[0].finish, assignments[1].start, assignments[1].finish)
	}

	return out, nil
}

func (a Assigner) CountAssignmentIntersections() (int, error) {
	r := bytes.NewReader(a.assignmentList)
	scanner := bufio.NewScanner(r)
	out := 0
	for scanner.Scan() {
		// Parse the pairs
		al, err := parsePairsToLists(scanner.Text())
		if err != nil {
			return 0, err
		}
		// Check for overlap
		out = out + checkContainsPair(al[0], al[1])
	}

	return out, nil
}

// parsePairs parses a pair '2-4,6-8' into two arrays
func parsePairs(pair string) ([]assignment, error) {
	// Split on comma to get two pairs
	ss := strings.Split(pair, ",")

	assignments := make([]assignment, 0, 2)
	for _, s := range ss {
		// Split on '-' to get first and final value of range of pair
		r := strings.Split(s, "-")
		start, err := strconv.Atoi(r[0])
		finish, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, nil
		}
		a := assignment{
			start:  start,
			finish: finish,
		}
		assignments = append(assignments, a)
	}

	return assignments, nil
}

func parsePairsToLists(pair string) ([]assignmentList, error) {
	// Split on comma to get two pairs
	ss := strings.Split(pair, ",")

	a := make([]assignmentList, 2)

	for i, s := range ss {
		a[i] = make(assignmentList, 0)
		// Split on '-' to get first and final value of range of pair
		r := strings.Split(s, "-")
		start, err := strconv.Atoi(r[0])
		if err != nil {
			return nil, nil
		}
		finish, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, nil
		}
		for j := start; j <= finish; j++ {
			a[i] = append(a[i], j)
		}
	}

	return a, nil
}

// checkFullyContainsPairs returns 1 if 1 member of the pair contains the other
// 0 otherwise
func checkFullyContainsPairs(a1, a2, b1, b2 int) int {
	if a1 <= b1 && a2 >= b2 {
		return 1
	} else if b1 <= a1 && b2 >= a2 {
		return 1
	} else {
		return 0
	}
}

// checkFullyContainsPairs returns 1 if 1 member of the pair contains the other
// 0 otherwise
func checkContainsPair(a, b []int) int {
	// Check for an intersection
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			return 1
		}
	}

	return 0
}
