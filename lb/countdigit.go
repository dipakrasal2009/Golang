package main

import "fmt"

func countdigits(iNo int )int{
	var digit int= 0
	var iCnt int = 0

	for iNo != 0{
		digit = iNo % 10
		iNo = iNo / 10
		iCnt++
	}
	return iCnt
}

func main(){

	var iValue = 9
	var iRet = 0

	fmt.Println("Enter the number")
	fmt.Scanln(iValue)

	iRet = countdigits(iValue)

	fmt.Println("number of digits are :",iRet)
}
