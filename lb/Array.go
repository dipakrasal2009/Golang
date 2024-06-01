package main

import "fmt"

func main(){
	var Arr[5]int
	var iCnt= 0

	fmt.Println("enter the elements")

	for iCnt = 0;iCnt < 5;iCnt ++{
		fmt.Scanln(&Arr[iCnt])
	}

	fmt.Println("elements from array are")
	
	for iCnt = 0;iCnt < 5;iCnt++{
		fmt.Println(Arr[iCnt])
	}
}
