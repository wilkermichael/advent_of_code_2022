package cargo

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Mover represents a mover of cargo containers
type Mover struct {
	containers containers
	movePlan   []movePlan
}

// NewMover is a constructor for the Mover type.
func NewMover(fileName string) (Mover, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return Mover{}, err
	}

	c, mp, err := parseOperatorGuide(string(b))
	if err != nil {
		return Mover{}, err
	}

	return Mover{
		containers: c,
		movePlan:   mp,
	}, nil
}

// PerformMovements implement the move plans stored in the Mover object by moving and removing
// elements from the containers stored in the Mover object.
func (m Mover) PerformMovements() string {
	// Move and remove list elements according to the move plan
	// Would have been nice to just move around the nodes, but the `list` package
	// seems to only allow copying from one list to another
	for _, mp := range m.movePlan {
		stoppoint := m.containers[mp.moveFrom].Len() - mp.numCratesToMove
		for i := m.containers[mp.moveFrom].Len(); i > stoppoint; i-- {
			m.containers[mp.moveTo].PushBack(m.containers[mp.moveFrom].Back().Value)
			m.containers[mp.moveFrom].Remove(m.containers[mp.moveFrom].Back())
		}
	}

	// Output the top of each column
	var out string
	for i := 1; i <= len(m.containers); i++ {
		var s string
		if m.containers[i].Back() != nil {
			s = fmt.Sprint(m.containers[i].Back().Value)
			out = out + string(s[1])
		}
	}

	return out
}

// PerformMovementsBulk method is similar to the PerformMovements method, but it moves elements in bulk rather
// than individually. This results in a different ordering of containers as the order of bulk moves is preserved.
func (m Mover) PerformMovementsBulk() string {
	// Create a temporary list to hold the bulk movement
	l := list.New()
	for _, mp := range m.movePlan {
		stoppoint := m.containers[mp.moveFrom].Len() - mp.numCratesToMove
		for i := m.containers[mp.moveFrom].Len(); i > stoppoint; i-- {
			l.PushFront(m.containers[mp.moveFrom].Back().Value)
			m.containers[mp.moveFrom].Remove(m.containers[mp.moveFrom].Back())
		}

		m.containers[mp.moveTo].PushBackList(l)

		// Clear the temporary list
		l.Init()
	}

	// Output the top of each column
	var out string
	for i := 1; i <= len(m.containers); i++ {
		var s string
		if m.containers[i].Back() != nil {
			s = fmt.Sprint(m.containers[i].Back().Value)
			out = out + string(s[1])
		}
	}

	return out
}

func parseOperatorGuide(operatorGuide string) (containers, []movePlan, error) {
	// Parse out the containers and move plan
	ss := strings.Split(operatorGuide, "\n\n")

	// Containers are in the first section
	c := newContainers(ss[0])

	// move plan will be in the second section
	mp, err := newMovePlans(ss[1])
	if err != nil {
		return nil, nil, err
	}

	return c, mp, nil
}

type containers map[int]*list.List

func newContainers(containerGuide string) containers {
	// Subtract the space from the end of the line
	a := containerGuide[len(containerGuide)-2]
	numContainers := a - '0'

	// create the containers
	c := make(containers)
	for i := 1; i <= int(numContainers); i++ {
		c[i] = list.New()
	}

	// Split on new line and throw away last row
	ss := strings.Split(containerGuide, "\n")
	containerCount := 1
	for i := 0; i < len(ss)-1; i++ {
		s := ss[i]
		// Scan through the lines in chunks of 4 characters: 0:3, 4:7, etc. to get container names
		for j := 0; j <= len(s); j = j + 4 {
			v := s[j : j+3]
			if s[j] == '[' {
				c[containerCount].PushFront(v)
			}
			containerCount++
		}
		containerCount = 1
	}

	return c
}

type movePlan struct {
	numCratesToMove int
	moveFrom        int
	moveTo          int
}

type movePlans []movePlan

func newMovePlans(moveGuide string) (movePlans, error) {
	r := bytes.NewReader([]byte(moveGuide))
	scanner := bufio.NewScanner(r)

	re := regexp.MustCompile(`\d+`)
	mp := make([]movePlan, 0)
	for scanner.Scan() {
		s := scanner.Text()

		m := re.FindAllString(s, -1) // -1 indicates to find all
		numCratesToMove, err := strconv.Atoi(m[0])
		if err != nil {
			return nil, err
		}
		moveFrom, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		moveTo, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
		mp = append(mp, movePlan{
			numCratesToMove: numCratesToMove,
			moveFrom:        moveFrom,
			moveTo:          moveTo,
		})
	}

	return mp, nil
}
