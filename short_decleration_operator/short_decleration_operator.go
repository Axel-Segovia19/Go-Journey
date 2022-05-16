package main

import "fmt"


func main() {
	x := 365
	y := "Disco Stew"
	g := 50 + 50
	z := x + g
	c := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(g)
	fmt.Println(z)
	fmt.Println(c)
}
// strings that are later reassigned get reassined throughout the whole code. since code is read top
// to bottom if assigned at top then reassigned at bottom it will change for the top as well. be careful