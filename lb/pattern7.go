package main

import "fmt"

func Display(iRow int ,iCol int){
	var i = 0
	var j = 0

	for i = 0;i <= iRow;i++{
		for j = 0;j <= iCol;j++{
			if i % 2 == 0{
				fmt.Print("#\t")
			}else{
				fmt.Print("*\t")
			}
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

