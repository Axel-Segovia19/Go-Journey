// package main

// import "fmt"

// func main() {
// 	x := 26
// 	fmt.Println(&x)
// }

package main 

import "fmt"

type person struct{
	fn string
	ln string
}

func changeMe( p *person){
	x := &person{
		fn: "Axel",
		ln: "Segovia",
	}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func main() {
	p1 := person{
		fn: "angel",
		ln: "Macias",
	}
	changeMe(&p1)
}