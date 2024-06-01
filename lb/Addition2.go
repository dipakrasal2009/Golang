
//perform the arithmetic operation using function and get the value from user

package main

import "fmt"

func Addition(num1 int,num2 int)int{
	var Answer int  = 0;

	Answer = num1 + num2;

	return Answer
}

func multiplication(num1 int ,num2 int)int{
	var Answer int = 0

	Answer = num1 * num2

	return Answer
}

func main(){

	var number1 int = 0;
	var number2 int = 0;
	var result int = 0;
	var result1 int = 0;
	fmt.Println("enter the first number ")
	fmt.Scanln(&number1)

	fmt.Println("enter the second number ")
	fmt.Scanln(&number2)

	result = Addition(number1,number2)
	result1 = multiplication(number1,number2)

	fmt.Println("addition is :",result)
	fmt.Println("multiplication is :",result1)
}
