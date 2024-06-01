package main

import "fmt"

func Summation(Data[] int, iSize int)int{
	var iCnt = 0
	var iSum = 0

	for iCnt = 0; iCnt < iSize;iCnt++{
		iSum = iSum + Data[iCnt]
	}
	return iSum
}


func main(){
	var Arr[5] int
	var iCnt = 0
	var iRet = 0

	fmt.Println("enter the element")

	for iCnt = 0;iCnt < 5;iCnt++{
		fmt.Scanln(&Arr[iCnt])
	}

	fmt.Println("element from the array are")

	for iCnt = 0;iCnt <5; iCnt++{
		fmt.Println(Arr[iCnt])
	}

	iRet = Summation(Arr[5] int,5)

	fmt.Println("Addition of all element is ; ",iRet)
}

