//4

//1	2	3	4	*	*	*	*




package main

import "fmt"


func Display(iNo int){
	var iCnt = 0
	if iNo < 0{
		iNo = -iNo
	}

	for iCnt = 0;iCnt <= iNo;iCnt++{
		fmt.Print(iCnt,"\t")
	}
	for iCnt = 1;iCnt <=iNo;iCnt++{
		fmt.Print("*\t")
	}
	fmt.Println()
}

func main(){
	var iValue = 0
	fmt.Println("enter the number")
	fmt.Scanln(&iValue)

	Display(iValue)
}


