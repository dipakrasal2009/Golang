package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	//pattern        = flag.String("grep", "", "pattern to search in to the file")
	showLineNumber = flag.Bool("n", false, "Show line numbers")
	before         = flag.Int("B", 0, "Print N lines before each match")
	after          = flag.Int("A", 0, "Print N lines after each match")
	countMatches   = flag.Bool("c", false, "Count matching lines")
)

const (
	redStart = "\033[31m"
	redEnd   = "\033[0m"
)

func main() {
	// Define and parse command-line flags
	flag.Parse()

	// Check if pattern is provided

	//Get the list of files from the command-line arguments

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go [-n] <pattern> <file1> <file2> ... <fileN>")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	if pattern == "" {
		fmt.Println("Usage: go run main.go -pattern <pattern> <file1> <file2> ... <fileN>")
		os.Exit(1)
	}

	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Invalid pattern: %v\n", err)
		os.Exit(1)
	}

	// Process each file
	for _, file := range files {
		err := grepFile(file, re, *showLineNumber, *countMatches, *before, *after)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", file, err)
		}
	}
}

// grepFile processes a single file and prints lines that match the pattern
func grepFile(filePath string, re *regexp.Regexp, showLineNumber, countMatches bool, before, after int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //
	var lines []string
	lineNumber := 0
	matchCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for i := 0; i < len(lines); i++ {
		if re.MatchString(lines[i]) {
			matchCount++
			if !countMatches {
				start := i - before
				if start < 0 {
					start = 0
				}
				end := i + after + 1
				if end > len(lines) {
					end = len(lines)
				}

				for j := start; j < end; j++ {
					if showLineNumber {
						fmt.Printf("%d : %s\n", j+1, highlightPattern(lines[j], re))
					} else {
						fmt.Println(highlightPattern(lines[j], re))
					}
				}
			}
		}
	}

	if countMatches {
		fmt.Printf("%s: %d\n", filePath, matchCount)
	}

	return nil
}

func highlightPattern(line string, re *regexp.Regexp) string {
	return re.ReplaceAllStringFunc(line, func(match string) string {
		return redStart + match + redEnd
	})
}
