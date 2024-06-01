package main

import "fmt"


func main(){
	var numbers [3]string

	numbers[0]  = "Dipak"
	numbers[1] = "Pratik"
	numbers[2] = "omkar"

	fmt.Println(numbers)

	var colours = [4]string{"Red","Yellow","Orange","green"}

	fmt.Println(colours)
	fmt.Println(colours[2])
	fmt.Println(len(colours))
}

