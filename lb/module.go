package greeting
package main

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func main(){
	var c string = "Dipak"

	var ret string = ""

	ret = Hello(c)
	fPrintln(ret)

}
