// check the given number is divisible by 5 or 3

package main

import "fmt"

func DivisibleByFiveOrThree(n int)int{
	if n % 5 == 0 && n %3 == 0{
		return 1
	}else{
		return 0
	}
}


func main(){
	var number int = 0
	var ret = 0

	fmt.Println("enter the number ")
	fmt.Scanln(&number)

	ret = DivisibleByFiveOrThree(number)

	if ret == 1{
		fmt.Println("number is divisible by five or three")
	}else{
		fmt.Println("number is not Divisible by five or three")
	}
}
