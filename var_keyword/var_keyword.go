package main

import "fmt"
var y bool
var t int
// this is how you declare an empty string later to be assigned something different
var g string
var u float32
var z = 45
func main(){
	x := 65

	fmt.Println(x)
	fmt.Println(z)
// this will add the local var and the global var 
	fmt.Println(x + z)
	fmt.Println(y)
	fmt.Println(t)
	g = "GoooooooooLaooooonnngggg"
	fmt.Println(g)
	fmt.Println(u)
// this will not reassign the variable on the global scope
	foo()
}

func foo() {
	fmt.Println(z)
}