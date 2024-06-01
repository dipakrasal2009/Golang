//4
//a	b	c	d

package main

import "fmt"

func Display(iNo int){
	var iCnt = 0
	var ch string ="A"

	if iNo < 0{
		iNo = -iNo
	}

	for iCnt = 1; iCnt <= iNo;iCnt++{
		fmt.Printf(ch,"\t")
		ch++
	}
	fmt.Println()
}


func main(){
	var iValue = 0

	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	Display(iValue)
}
