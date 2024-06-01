package main

import "fmt"

func DisplayEvenOddDigits(iNo int){
	var iEvenCnt = 0
	var iDigit = 0
	var iOddCnt = 0

	if iNo == 0{
		iEvenCnt++;
	}

	for iNo != 0{
		iDigit = iNo % 10
		if iDigit % 2 == 0{
			iEvenCnt++
		}
		else
		{
			iOddCnt++
		}
		iNo = iNo / 10;
	}
	fmt.Println("number od even digits are :",iEvenCnt)
	fmt.Println("number of odd Digits are : ",iOddCnt)
}


func main(){
	var iValue  = 0

	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	DisplayEvenOddDigits(iValue)
}
