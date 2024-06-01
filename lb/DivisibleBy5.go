//chcek wether the given number is divisible  by five or not

package main

import "fmt"

func divisiblebyfive(num int){
	if num % 5 == 0{
		fmt.Println("the number is divisible by five")
	}else{
		fmt.Println("number is not divisible by five")
	}
}

func main(){

	var number int = 0;
	fmt.Println("enter the number ffrom user ")
	fmt.Scanln(&number)

	divisiblebyfive(number)
}
