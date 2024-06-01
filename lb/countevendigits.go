package main

import "fmt"

func countevendigits(iNo int)int{
	var iEvenCnt = 0
	var iDigits = 0

	if iNo != 0{
		return 1
	}

	for iNo != 0{
		iDigits = iNo %10
		if iDigits % 2 == 0{
			iEvenCnt++
		}
		iNo = iNo/10
	}
	return iEvenCnt
}


func main(){
	var iValue = 0
	var iRet = 0;

	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	iRet = countevendigits(iValue)

	fmt.Println("number of even digits are :",iRet)
}
