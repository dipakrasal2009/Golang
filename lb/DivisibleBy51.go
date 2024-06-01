
//check the given numbe id divisible by five or not 


package main

import "fmt"

func divisiblebyfive(num int)int{
	var ans = 0

	ans = num % 5
	return ans
}

func main(){

	var i int = 0
	var ret = 0;

	fmt.Println("Enter the number ")
	fmt.Scanln(&i)

	ret = divisiblebyfive(i)
	if ret == 0{
		fmt.Println("number is divisible by five")
	}else{
		fmt.Println("number is not Divisible by five")
	}
}
