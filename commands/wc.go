package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

var (
	fileName = flag.String("file", "", "file to process")
)

func main() {
	// Define a flag for the file name
	//fileName := flag.String("file", "", "File to process")
	flag.Parse()

	// Ensure a file name is provided
	if *fileName == "" {
		fmt.Println("Please provide a file name using the -file flag.")
		os.Exit(1)
	}

	// Process the file
	lines, words, bytes, err := processFile(*fileName)
	if err != nil {
		fmt.Println("Error processing file:", err)
		os.Exit(1)
	}

	// Print the results
	fmt.Printf("Lines: %d\nWords: %d\nBytes: %d\n", lines, words, bytes)
}

// processFile processes the file and returns the counts of lines, words, and bytes
func processFile(fileName string) (int, int, int, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	// Initialize counters
	var lines, words, bytes int

	// Create a new reader
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				lines++
				words += countWords(line)
				bytes += len(line)
			}
			break
		}
		if err != nil {
			return 0, 0, 0, err
		}

		lines++
		words += countWords(line)
		bytes += len(line)
	}

	return lines, words, bytes, nil
}

// countWords counts the number of words in a given string
func countWords(s string) int {
	inWord := false
	wordCount := 0
	for _, r := range s {
		if unicode.IsSpace(r) {
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}
	if inWord {
		wordCount++
	}
	return wordCount
}

//git remote set -url origin https://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@github.com/dipakrasal2009/Golang

//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
