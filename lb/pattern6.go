package main

import "fmt"
func Display(iRow int,iCol int){
	var i = 0
	var j = 0

	if iRow < 0{
		iRow = -iRow
	}

	if iCol < 0{
		iCol = -iCol
	}

	for i = 1;i<=iRow;i++{
		for j=1;j<=iCol;j++{
			fmt.Print("*\t")
		}
		fmt.Println()
	}
}

func main(){
	var iValue1 = 0
	var iValue2 = 0

	fmt.Println("enter the number of rows")
	fmt.Scanln(&iValue1)

	fmt.Println("enter the number of column")
	fmt.Scanln(&iValue2)

	Display(iValue1,iValue2)
}

