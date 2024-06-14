package main

import (
	//"bytes"
	//	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"syscall"
	"text/tabwriter"
	"time"

	//"strconv"

	//"os"
	"os/user"
)

// Struct to hold command line arguments
type Config struct {
	dir           []string
	reverse       bool
	list          bool
	humanReadable bool
}

const (
	colorBlue  = "\033[34m"
	colorReset = "\033[0m"
)

// Function to parse command line arguments
func parseArgs() Config {
	var config Config

	// Get non-flag argument (directory)
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: ls <directory> [-r] [-h]")
		os.Exit(1)
	}

	//config.dir = args[0]
	//args = args[1:] // Remove directory argument from args

	// Parse flags
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-r":
			config.reverse = true
		case "-l":
			config.list = true
		case "-h":
			config.humanReadable = true
		default:
			config.dir = append(config.dir, args[i])
		}
	}

	return config
}

// Function to read and sort directory contents
func readDir(dir string, reverse bool) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Sort the files
	sort.Slice(files, func(i, j int) bool {
		if reverse {
			return files[i].Name() > files[j].Name()
		}
		return files[i].Name() < files[j].Name()
	})

	return files, nil
}

// Function to get owner of a file
func getFileOwner(file os.FileInfo) (string, error) {
	uid := file.Sys().(*syscall.Stat_t).Uid
	u, err := user.LookupId(fmt.Sprintf("%d", uid))
	if err != nil {
		return "", err
	}
	return u.Username, nil
}

func getAccessCount(file os.FileInfo) int {
	stat := file.Sys().(*syscall.Stat_t)
	return int(stat.Nlink)
}

// Function to print directory contents
func printFiles(files []os.FileInfo, list, humanReadable bool) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for _, file := range files {
		if list {
			modTime := file.ModTime().Format(time.RFC822)

			var size string

			if humanReadable {
				const (
					KB = 1024
					MB = KB * 1024
					GB = MB * 1024
					TB = GB * 1024
				)

				if file.Size() < KB {
					size1 := file.Size()
					size = fmt.Sprintf("%d", size1)
				} else if file.Size() < MB {
					size = fmt.Sprintf("%.0fK", float64(file.Size())/KB)
				} else if file.Size() < GB {
					size = fmt.Sprintf("%.1fM", float64(file.Size())/MB)
				} else if file.Size() < TB {
					size = fmt.Sprintf("%.2fGB", float64(file.Size())/GB)
				} else {
					size = fmt.Sprintf("%.2fTB", float64(file.Size())/TB)
				}

			} else {
				size = fmt.Sprint(file.Size())
			}
			modTime = file.ModTime().Format(time.RFC822)
			permissions := file.Mode().String()
			name := file.Name()
			owner, err := getFileOwner(file)
			//grp, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))

			accessCount := getAccessCount(file)

			if err != nil {
				fmt.Printf("Error getting owner for file %s: %v\n", name, err)
				return
			}

			if file.IsDir() {
				name = colorBlue + name + colorReset
			}
			fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%s\n", permissions, accessCount, owner, size, modTime, name)
		} else {
			fmt.Fprintf(w, "%s\t", file.Name())
		}
	}
	w.Flush()

}

func main() {
	config := parseArgs()

	for _, dir := range config.dir {
		files, err := readDir(dir, config.reverse)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}
		if len(config.dir) > 1 {
			fmt.Printf("%s : \n", dir)
		}
		printFiles(files, config.list, config.humanReadable)
		fmt.Println()
	}

}
