package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// Define ANSI color codes
	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
)

func main() {
	// parse command line argument
	flag.Parse()
	root := "."

	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	var fileCount, dirCount int

	err := tree(root, "", &fileCount, &dirCount)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)

}

// tree function is recursivaly call the function and print the structure of tree
func tree(path string, prefix string, fileCount *int, dirCount *int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	files, err := file.Readdir(0)
	if err != nil {
		return err
	}

	for i, f := range files {
		isLast := i == len(files)-1
		printNode(f.Name(), prefix, isLast, f.IsDir())
		if f.IsDir() {
			*dirCount++
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			err := tree(filepath.Join(path, f.Name()), newPrefix, fileCount, dirCount)
			if err != nil {
				return err
			}
		} else {
			*fileCount++
		}
	}
	return nil
}

// print the single file or directory node
func printNode(name string, prefix string, isLast bool, isDir bool) {
	if isDir {
		name = colorBlue + name + colorReset
	}

	if isLast {
		fmt.Printf("%s└── %s\n", prefix, name)
	} else {
		fmt.Printf("%s├── %s\n", prefix, name)
	}
}
