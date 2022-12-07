package filesystem

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	fileSystemSize = 70000000
)

type directory struct {
	name           string
	files          map[string]int
	subDirectories map[string]*directory
	previous       *directory
	root           *directory
}

func newDirectory(name string, previous *directory) *directory {
	d := &directory{
		name:           name,
		files:          make(map[string]int),
		subDirectories: make(map[string]*directory),
		previous:       previous,
	}
	if previous != nil {
		d.previous.addSubDirectory(d)
		d.root = previous.root
	} else {
		d.root = d
	}
	return d
}

func (d *directory) addFile(fileName string, size int) {
	d.files[fileName] = size
}

func (d *directory) addSubDirectory(dir *directory) {
	d.subDirectories[dir.name] = dir
}

func (d *directory) changeInto(dir string) *directory {
	return d.subDirectories[dir]
}

func (d *directory) changeBack() *directory {
	return d.previous
}

func (d *directory) getRoot() *directory {
	return d.root
}

func (d *directory) calcDirectoryFileSizes() int {
	out := 0
	for _, v := range d.files {
		out = out + v
	}

	return out
}

func (d *directory) calculateDirectorySizes() (currSize int, allSizes []int) {
	currSize = d.calcDirectoryFileSizes()
	allSizes = make([]int, 0)
	for _, v := range d.subDirectories {
		o, o2 := v.calculateDirectorySizes()
		currSize = currSize + o
		allSizes = append(allSizes, o2...)
	}
	allSizes = append(allSizes, currSize)
	return currSize, allSizes
}

type FileTree struct {
	root *directory
}

func NewFileTree(fileName string) (FileTree, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return FileTree{}, err
	}

	ft := FileTree{}
	err = ft.constructTree(string(b))
	if err != nil {
		return FileTree{}, err
	}
	return ft, nil
}

func (ft *FileTree) GetSumOfDirectoriesUnderSpaceValue(maxSize int) int {
	dir := ft.root
	_, sizes := dir.calculateDirectorySizes()
	out := 0
	for _, v := range sizes {
		if v <= maxSize {
			out = out + v
		}
	}

	return out
}

func (ft *FileTree) GetMinimumSizeDirectoryToClearSpace(requiredSpace int) int {
	dir := ft.root
	size, sizes := dir.calculateDirectorySizes()
	freeSpace := fileSystemSize - size
	sort.Ints(sizes)
	for _, v := range sizes {
		freeSpaceAfterDelete := freeSpace + v
		if requiredSpace <= freeSpaceAfterDelete {
			return v
		}
	}
	return 0
}

func (ft *FileTree) constructTree(s string) error {
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)
	dir := newDirectory("/", nil)
	for scanner.Scan() {
		l := scanner.Text()
		switch {
		case strings.Contains(l, "$ cd .."):
			if dir.previous != nil {
				dir = dir.changeBack()
			}
		case strings.Contains(l, "$ cd /"):
			if dir.root != nil {
				dir = dir.getRoot()
			}
		case strings.Contains(l, "$ cd"):
			// Follows convention "$ cd dirname" so [2] will contain the dirname
			ss := strings.Fields(l)
			dir = dir.changeInto(ss[2])
		case strings.Contains(l, "$ ls"):
		default:
			// Making some assumptions to simplify checking, 0 index is either a number or a "dir"
			ss := strings.Fields(l)
			if ss[0] == "dir" {
				// This is a directory
				dir.addSubDirectory(newDirectory(ss[1], dir))
			} else {
				// This must be a file
				size, err := strconv.Atoi(ss[0])
				if err != nil {
					return err
				}
				dir.addFile(ss[1], size)
			}
		}
	}

	ft.root = dir.root
	return nil
}
