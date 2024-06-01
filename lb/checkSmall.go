package main

import "fmt"

func checkSmall(ch rune)bool{
	if ch >= '97' && ch <= '122'{
		return true
	}else{
		return false
	}
}

func main(){
	var cValue rune
	var bRet bool

	fmt.Println("enter the character")
	fmt.Scanln(&cValue)

	bRet = checkSmall(cValue)

	if bRet == true{
		fmt.Println("%c is a small case letter",cValue)
	}else{
		fmt.Println("%c is not a Small case letter",cValue)
	}
}

