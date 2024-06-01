package main

import  "fmt"

func countDigits(iNo int )int{
	var iDigits = 0
	var iCnt = 0

	if iNo < 0{
		iNo = -iNo
	}

	for iNo < 0{
		iDigits = iNo%10

		iNo = iNo/10
		iCnt++
	}

	return iCnt
}


func main(){

	var iValue = 0
	var iRet = 0

	fmt.Println("please enter the number that you wannt to search")
	fmt.Scanln(&iValue)

	iRet = countDigits(iValue)

	fmt.Println("number of digits are ",iRet)
}
