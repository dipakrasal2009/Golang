package main

import (
	"bufio"
	//"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	fileName = flag.String("file", " ", "name of the file")

	disolayNumber = flag.Bool("n", false, "line number of the file")
)

func DisplayData(fileName string) string {
	Data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("unable to ope the file", err)
		os.Exit(1)
	}
	return string(Data)
	//fmt.Println(Data(string))
}

func printFileDetails(fileName string) {
	fmt.Println("there is thesimple cat utility")

	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Could not get the file Details")
		return
	}
	fmt.Println("Filse Name : %s", info)

}

func process(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		//return errors.Wrap(err, "could not open File")
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()

		if *disolayNumber {
			fmt.Printf("%d : %s\n", lineNumber, line)
			lineNumber++
			continue
		}
		fmt.Println(line)
	}
	return nil

}

func main() {
	flag.Parse()
	var d string
	d = DisplayData(*fileName)
	fmt.Println(d)

	printFileDetails(*fileName)

	process(*fileName)
	fmt.Println("Dipak Rasal")

}
