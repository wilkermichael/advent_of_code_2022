package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_changeInto(t *testing.T) {
	d1 := &directory{
		name:           "/",
		files:          make(map[string]int),
		subDirectories: make(map[string]*directory),
		previous:       nil,
	}
	d2 := &directory{
		name:           "a",
		files:          make(map[string]int),
		subDirectories: make(map[string]*directory),
		previous:       d1,
	}
	d1.addSubDirectory(d2)

	actual := d1.changeInto("a")
	assert.Equal(t, "a", actual.name)
}

func Test_changeBack(t *testing.T) {
	d1 := &directory{
		name:           "/",
		files:          make(map[string]int),
		subDirectories: make(map[string]*directory),
		previous:       nil,
	}
	d2 := &directory{
		name:           "a",
		files:          make(map[string]int),
		subDirectories: make(map[string]*directory),
		previous:       d1,
	}
	d1.addSubDirectory(d2)

	actual := d1.changeInto("a").changeBack()

	assert.Equal(t, "/", actual.name)
}

func Test_calculateDirectoryFileSize(t *testing.T) {
	d1 := &directory{
		name: "/",
		files: map[string]int{
			"a.txt": 2241,
			"b":     5534,
		},
		subDirectories: make(map[string]*directory),
		previous:       nil,
	}
	actual := d1.calcDirectoryFileSizes()
	expected := 2241 + 5534

	assert.Equal(t, expected, actual)
}

func Test_calcDirectorySize(t *testing.T) {
	root := newDirectory("/", nil)
	a := newDirectory("a", root)
	c := newDirectory("c", root)
	b := newDirectory("b", a)

	a.addFile("f.txt", 3)
	a.addFile("z.txt", 1)
	b.addFile("f.txt", 3)
	b.addFile("z.txt", 1)
	c.addFile("f.txt", 3)
	c.addFile("z.txt", 1)
	root.addFile("f.txt", 3)
	root.addFile("z.txt", 1)

	expectedFileSizeRoot := 16
	expecteFileSizedB := 4

	rootSize, rootSizes := root.calculateDirectorySizes()
	assert.Equal(t, expectedFileSizeRoot, rootSize)
	assert.Equal(t, []int{4, 8, 4, 16}, rootSizes)

	bSize, bSizes := b.calculateDirectorySizes()
	assert.Equal(t, expecteFileSizedB, bSize)
	assert.Equal(t, []int{4}, bSizes)
}

func Test_constructTree(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	ft := FileTree{}
	ft.constructTree(input)
	assert.Equal(t, "/", ft.root.name)
	dir := ft.root.subDirectories["a"].subDirectories["e"]
	assert.Equal(t, "e", dir.name)
	assert.Equal(t, 584, dir.files["i"])
}

func Test_GetSumOfDirectoriesUnderSpaceValue(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	maxSize := 100000
	ft := FileTree{}
	ft.constructTree(input)
	actual := ft.GetSumOfDirectoriesUnderSpaceValue(maxSize)
	expected := 95437
	assert.Equal(t, expected, actual)
}

func Test_GetMinimumSizeDirectoryToClearSpace(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	desiredSpace := 30000000
	ft := FileTree{}
	ft.constructTree(input)
	actual := ft.GetMinimumSizeDirectoryToClearSpace(desiredSpace)
	expected := 24933642
	assert.Equal(t, expected, actual)
}
