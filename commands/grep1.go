package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

const (
	redStart = "\033[31m"
	redEnd   = "\033[0m"
)

func main() {
	// Define command-line flags
	showLineNumber := flag.Bool("n", false, "Show line numbers")
	caseInsensitive := flag.Bool("i", false, "Case insensitive search")
	countMatches := flag.Bool("c", false, "Count matching lines")
	before := flag.Int("B", 0, "Print N lines before each match")
	after := flag.Int("A", 0, "Print N lines after each match")
	flag.Parse()

	// Get the list of arguments
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go [-n] [-i] [-c] [-B n] [-A n] <pattern> <file1> <file2> ... <fileN>")
		os.Exit(1)
	}

	// The first argument is the pattern, the rest are files
	pattern := args[0]
	files := args[1:]

	// Compile the regex pattern
	var re *regexp.Regexp
	var err error
	if *caseInsensitive {
		re, err = regexp.Compile("(?i)" + pattern)
	} else {
		re, err = regexp.Compile(pattern)
	}
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
// with the pattern highlighted in red and optionally the line number
func grepFile(filePath string, re *regexp.Regexp, showLineNumber, countMatches bool, before, after int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
						fmt.Printf("%d:%s\n", j+1, highlightPattern(lines[j], re))
					} else {
						fmt.Println(highlightPattern(lines[j], re))
					}
				}
			}
		}
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
						fmt.Printf("%d:%s\n", j+1, highlightPattern(lines[j], re))
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

// highlightPattern highlights all occurrences of the pattern in the line
func highlightPattern(line string, re *regexp.Regexp) string {
	return re.ReplaceAllStringFunc(line, func(match string) string {
		return redStart + match + redEnd
	})
}
