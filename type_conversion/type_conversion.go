package main

import "fmt"

var g int // gow to assign a 0 value to g that you can later assign it a different value
var a = 45

type dog int // this is how to call your own type and assign it a second type

var c dog //this will produce a 0 value until it is assigned one.

func main() {
	c = 43
	// a = c will not  work since its type dog and not a true int
	// needs to be converted using int() conversion syntax
	//then you can do a = int(c) since both are int's and a will now equal 43
	a = int(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",c)
	// fmt.Printf("%T\n", c) is to print the type if data type it is. there are more codes 
	//to find out what they mean go to the documentation and understand the difference
	// there is also fmt.Sprint, fmt.Fprint etc.
}