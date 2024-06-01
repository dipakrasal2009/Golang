//Accept the number from user and add that two number`


package main

import "fmt"

func main(){
	var number1 int = 0
	var number2 int = 0
	var Result = 0
	fmt.Println("Enter the First nummber : ")
	fmt.Scanln(&number1)

	fmt.Println("Enter the second number : ")
	fmt.Scanln(&number2)

	Result = number1 + number2

	fmt.Println("Addition is : ",Result)


}
