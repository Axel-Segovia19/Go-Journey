// package main

// import "fmt"

// func main() {
// 	x := 42
// 	y := "James Bond"
// 	z := true

// 	// fmt.Println(x, y, z)
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)

// }

// package main

// import "fmt"

// var x int
// var y string
// var z bool

// func main() {
// 	fmt.Println(x, y, z)
// 	fmt.Printf("%T\n",x)
// 	fmt.Printf("%T\n",y)
// 	fmt.Printf("%T\n",z)

// }

// package main

// import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

// func main() {
// 	s := fmt.Sprintf("%v\t%v\t%v",x, y, z)

// 	fmt.Println(s)

package main 

import "fmt"

type Dog int

var x Dog
var y int 

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}