/*
Row 5
Col 5

1	2	3	4	5
1	2	3	4	5
1	2	3	4	5
1	2	3	4	5
1	2	3	4	5
*/



package main

import "fmt"

func Display(iRow int,iCol int){
        var i = 0
        var j = 0

        for i = 1;i<=iRow;i++{
                for j = 1;j<=iCol;j++{
                        fmt.Print(i,"\t")
                }
                fmt.Println()
        }
}


func main(){
        var iValue1 = 0
        var iValue2 = 0

        fmt.Println("enter the first number")
        fmt.Scanln(&iValue1)

        fmt.Println("enter the second number")
        fmt.Scanln(&iValue2)

        Display(iValue1,iValue2)
}
                      
