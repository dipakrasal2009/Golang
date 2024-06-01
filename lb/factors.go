package main

import "fmt"

func DisplayEvenFactors(iNo int){
	var iCnt = 0

	for iCnt = 2;iCnt <= iNo;iCnt++{
		if iNo %iCnt ==0{
			fmt.Println(iCnt)
		}
	}
}

func main(){
	var  iValue = 0;
	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	DisplayEvenFactors(iValue)
}
