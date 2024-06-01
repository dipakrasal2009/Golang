package main

import "fmt"

func Summation(num int)int {
	var i = 0;
	var sum int = 0
	for i = 0; i <= num; i++{
		sum = sum + i
	}
	return sum
}

func main(){
	var i int = 0
	var ret int = 0 
	fmt.Println("enter the number you want to add")
	fmt.Scanln(&i)

	ret = Summation(i)

	fmt.Println("summation of the number is ",ret)
}


	


