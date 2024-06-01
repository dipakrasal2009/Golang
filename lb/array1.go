package main

import "fmt"

func main(){

	var iLength = 0
	var arr[100] int
	var iCnt = 0

	fmt.Println("enter the number that youn want to enter")

	fmt.Scanln(&iLength)

	fmt.Println("eneter the element in the array ")

	for iCnt  = 0;iCnt < iLength;iCnt++{
		fmt.Scanln(&arr[iCnt])
	}

	fmt.Println("your enterd number is ")

	for iCnt = 0; iCnt < iLength; iCnt++{
		fmt.Println(arr[iCnt])
	}
}
