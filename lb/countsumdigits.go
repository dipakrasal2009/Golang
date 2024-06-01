package main

import "fmt"

func SumDigits(iNo int)int{
	var Digits = 0
	var iSum = 0

	if iNo < 0{
		iNo = -iNo
	}
	for iNo != 0{
		Digits = iNo % 10
		iSum = iSum + Digits
		iNo = iNo/10
	}
	return iSum
}


func main(){
	var iValue = 0
	var iRet = 0

	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	iRet = SumDigits(iValue)

	fmt.Println("summation of digists are :",iRet)
}
