package main

import "fmt"

func MaxDigit(iNo int)int{
	var iDigit = 0
	var iMax = 0

	if iNo < 0{
		iNo = -iNo
	}

	for iNo != 0{
		iDigit = iNo %10
		if iDigit > iMax{
			iMax = iDigit
		}
		iNo = iNo / 10
	}
	return iMax
}

func main(){
	var iValue = 0
	var iRet = 0

	fmt.Println("enter the number ")
	fmt.Scanln(&iValue)

	iRet = MaxDigit(iValue)

	fmt.Println("largest digit is ",iRet)
}
