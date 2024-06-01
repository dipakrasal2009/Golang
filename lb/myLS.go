package main

import (
	"fmt"
	"io/ioutil"
//	"flag"
)
func main(){
	
//	LS := flag.String(".","","to display the list of files")
//	flag.Parse()


	files, err := ioutil.ReadDir(".")

	if err != nil{
		fmt.Println("unable to read the directory ",err)
		return
	}

	for _, file := range files {
      		fmt.Println(file.Name())
   	}
}
