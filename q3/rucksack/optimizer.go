package rucksack

import (
	"bufio"
	"bytes"
	"os"
	"unicode"
)

type Optimizer struct {
	rucksackContents []byte
}

func NewOptimizer(fileName string) (Optimizer, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return Optimizer{}, err
	}
	return Optimizer{
		rucksackContents: b,
	}, nil
}

func (o Optimizer) CalculateDuplicatePriorityTotal() int {
	r := bytes.NewReader(o.rucksackContents)
	scanner := bufio.NewScanner(r)
	out := 0
	for scanner.Scan() {
		// Split the string
		t := scanner.Text()
		s1, s2 := splitItems(t)

		// Find the Duplicates
		dupes := findDuplicates(s1, s2)

		// Calculate the priority
		for _, v := range dupes {
			out = out + decodeRunePriority(v)
		}
	}

	return out
}

func (o Optimizer) CalculatePriorityOfBadges() int {
	r := bytes.NewReader(o.rucksackContents)
	scanner := bufio.NewScanner(r)
	out := 0

	// Scan elf groups Create the elf groups
	i := 0
	elfGroupSize := 3
	buf := make([]string, 3)
	for scanner.Scan() {
		buf[i] = scanner.Text()
		i++
		if i > elfGroupSize-1 {
			v := findSharedDuplicate(buf[0], buf[1], buf[2])
			out = out + decodeRunePriority(v)
			i = 0
		}
	}

	return out
}

func decodeRunePriority(r rune) int {
	var out int
	if unicode.IsLower(r) {
		out = int(r - 'a' + 1)
	} else {
		out = int(r - 'A' + 27)
	}
	return out
}

func splitItems(s string) (string, string) {
	l := len(s)
	hl := l / 2
	return s[0:hl], s[hl:l]
}

func findDuplicates(s1, s2 string) []rune {
	m := make(map[rune]bool)
	for _, v := range s2 {
		m[v] = true
	}

	out := make([]rune, 0)
	set := make(map[rune]bool)
	for _, v := range s1 {
		if _, ok := m[v]; ok {
			if _, ok2 := set[v]; !ok2 {
				set[v] = true
				out = append(out, v)
			}
		}
	}

	return out
}

func findSharedDuplicate(s1, s2, s3 string) rune {
	m2 := make(map[rune]bool)
	for _, v := range s2 {
		m2[v] = true
	}

	m3 := make(map[rune]bool)
	for _, v := range s3 {
		m3[v] = true
	}

	var out rune
	for _, v := range s1 {
		if _, ok := m2[v]; ok {
			if _, ok2 := m3[v]; ok2 {
				out = v
				break
			}
		}
	}

	return out
}
