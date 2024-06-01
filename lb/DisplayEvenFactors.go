package main

import "fmt"

func DisplayEvenFactors(iNo int){
	var iCnt = 0;

	for iCnt = 1;iCnt <= iNo;iCnt++{
		if iNo % iCnt == 0 && iCnt % 2 != 0{
			fmt.Println(iCnt)
		}
	}
}

func main(){
	var iValue = 0

	fmt.Println("Enter the number")
	fmt.Scanln(&iValue)

	DisplayEvenFactors(iValue)
}
