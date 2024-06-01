package main

import "fmt"

func strlen(*str string)int{
	var iCnt int = 0
	for *str != '\0'{
		iCnt++
		str++
	}
	return iCnt
}


func main(){
	var arr[10] string
	var iRet = 0

	fmt.Println("please enter the string ")
	fmt.Scanln(&arr)

	iRet = strlen(arr)

	fmt.Println("number of character are :,"iRet)
}
