package main

import (
	"advent_of_code_q7/filesystem"
	"fmt"
)

const (
	maxSize      = 100000
	requiredSize = 30000000
)

func main() {
	ft, err := filesystem.NewFileTree("input.txt")
	if err != nil {
		panic(err)
	}

	totalSize := ft.GetSumOfDirectoriesUnderSpaceValue(maxSize)
	fmt.Printf("The total size of directories under %d is %d\n", maxSize, totalSize)

	fileSizeToDelete := ft.GetMinimumSizeDirectoryToClearSpace(requiredSize)
	fmt.Printf("The minimum total filesize to delete is %d", fileSizeToDelete)
}
