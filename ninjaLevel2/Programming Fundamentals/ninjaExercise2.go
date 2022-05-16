// print a number to a decimal, binary or hex
// package main

// import "fmt"

// func main() {
// 	x := 19
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// 	fmt.Printf("%d", x)
// 	fmt.Printf("%b", x)
// 	fmt.Printf("%#x", x)
// or can be done this way as well for less code.
// fmt.Printf("%d\t%b\t%#x", x, x, x) for a single liner.
// }

// print a numer using operators
// package main

// import "fmt"

// func main() {
// 	x := 4
// 	y := 5
// 	z := 4

// 	fmt.Println( x == z)
// 	fmt.Println( x <= y )
// 	fmt.Println( x >= y)
// 	fmt.Println( x != y)
// 	fmt.Println( x < y)
// 	fmt.Println( x > y)
// another way of doing it
// 	a := (42 == 42)
// 	b := (42 <= 43)
// 	c := (42 >= 43)
// 	d := (42 != 43)
// 	e := (42 < 43)
// 	f := (42 > 43)
// 	fmt.Println(a, b, c, d, e, f)
// }

// print type and untyped constants and their values

// package main

// import "fmt"

// const(
// 	a = 42
// 	b = "Axel Segovia"
// 	c = 4.625
// 	u//ntyped
// )

// const (
// 	//typed
// 	d int = 55
// 	e string = "James Bond"
// 	f float64 = 5.625
// )
// func main() {
// 	fmt.Println(a, b, c, d, e , f)
// }

// use bit shifting to move a varable 1 side to the left
// package main

// import "fmt"

// func main() {
// 	a := 27
// 	fmt.Printf("%d\t%b\t%#x", a, a , a)
// 	d := a << 1
// 	fmt.Printf("%d\t%b\t%#x", d, d , d)
// }

// use a variable of a string using raw string literal
// package main

// import "fmt"

// func main() {
// 	b := ` La Liga is dead! 
// 	"The Dream team is DEAD" - Chris Tucker
// 	`
// 	fmt.Println(b)
// }

// using iota create 4 constant for the next 4 years!

package main

const (
	year1 = iota
	
)
func main() {

}