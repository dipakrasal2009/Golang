package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	//"os/user"
	"time"
)

var (
	longFormat = flag.Bool("ls", false, "use to display the list os files anf dir")
)

func parseArgs() (string, bool) {
	flag.Parse()

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	return dir, *longFormat

}

// Function to read directory contents
func readDir(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Function to print directory contents
func printFilesInfo(files []os.FileInfo, longFormat bool) {

	for _, file := range files {
		if longFormat {

			modTime := file.ModTime().Format(time.RFC822)

			size := file.Size()

			name := file.Name()

			mode := file.Mode()
			//a := file.IsDir()
			fmt.Printf("%s \t%d \t%s\t %s \n", modTime, size, mode, name)
		} else {
			fmt.Println(file.Name())
		}
	}
}

func main() {
	dir, longFormat := parseArgs()

	files, err := readDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	printFilesInfo(files, longFormat)
}
