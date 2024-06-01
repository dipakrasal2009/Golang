package main

import "fmt"

func pointer(p *string){
	fmt.Println("value of p is :",*p)

	var m = *p
	fmt.Println("name of pointer function is :",m)
}

func main(){
	var name string = "Dipak"

	var ptr = &name;

	fmt.Println("Address of name is : ",&name)
	fmt.Println("Value of ptr is :",ptr)

	fmt.Println("name is",name)
	pointer(ptr)

}

