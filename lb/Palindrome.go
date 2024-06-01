package main

import "fmt"

func checkPalindrome(iNo int)bool{
	var iDigit = 0
	var iRev = 0
	var iTemp = iNo

	for iNo != 0{
		iDigit = iNo %10
		iRev = (iRev * 10) + iDigit
		iNo = iNo/10
	}
	if iRev == iTemp{
		return true
	}else{
		return false
	}
}



func main(){

	var iValue = 0
	var bRet = false

	fmt.Println("please enter the number")
	fmt.Scanln(&iValue)

	bRet = checkPalindrome(iValue)

	if bRet == true{
		fmt.Println("the number is a palindrome number",iValue)
	}else{
		fmt.Println("the number is not the palindrome number",iValue)
	}
}
