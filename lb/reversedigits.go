package main

import "fmt"

func Reverse(iNo int)int{
	var iDigit = 0
	var iRev = 0
	if iNo < 0{
		iNo = -iNo
	}

	for iNo != 0{
		iDigit = iNo %10
		iRev = (iRev*10) + iDigit
		iNo = iNo /10
	}
	return iRev
}



func main(){
	var iValue = 0
	var iRet = 0

	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	iRet = Reverse(iValue)

	fmt.Println("Reversen number is ",iRet)
}

