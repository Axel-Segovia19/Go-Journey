// package main

// import "fmt"

// func main() {
// 	x := foo()
// 	n, s := bar()

// 	fmt.Println(x, n, s)

// }
// func foo() int {
// 	return 42
// }

// func bar() (int, string) {
// 	// x := 42
// 	// y := "Hello World"
// 	return 42, "Hello World"
// }

// package main

// import "fmt"

// func main() {
// 	t := []int{ 1, 2, 3, 4 }
// 	g := foo(t ...)
// 	defer fmt.Println(g)

// 	r := []int{ 3, 6, 3, 4 }
// 	v := bar(r)
// 	fmt.Println(v)
// }
// func foo(x ...int) int {
// 	b := 0
// 	for _ , v := range x {
// 		b += v
// 	}
// 	return b
// }

// func bar(h []int) int{
// 	b := 0
// 	for _ , v := range h {
// 		b += v
// 	}
// 	return b
// }
// package main

// import "fmt"

// type person struct{
// 	first string
// 	last string
// 	age int
// }

// func (s person) speak(){
// 	fmt.Println("Hello my name is", s.first, s.last, "I am ", s.age, "years old.")
// }

// func main(){
// 	p1 := person{
// 		first: "Axel",
// 		last: "Segovia",
// 		age: 26,
// 	}
// 	fmt.Println(p1)
// 	p1.speak()
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type square struct{
// 	L float64
// }

// type circle struct{
// 	radius float64
// }

// type Shape interface{
// 	Area() float64
// }

// func (a square) Area() float64{
// 	return a.L * a.L
// }
// func (a circle) Area() float64{
// 	return  math.Pi * a.radius * a.radius
// }

// func info(s Shape){
// 	fmt.Println(s.Area())
// }
// func main() {
//  a := square{
// 	 L: 12,
//  }
//  b := circle{
// 	 radius: 12.546,
//  }

//  info(a)
//  info(b)
// }

// package main

// import "fmt"

// var g func()
// func main(){
// 	x := func() {
// 		fmt.Println("Hello I am an anonymous function")
// 	}
// 	x()

// 	g = x
// 	g()
// }
// package main

// import "fmt"

// func foo() func() string {
// 	return func() string{
// 		return "Hellow World"
// 	}

// }
// func main() {
// 	z := foo()
// 	fmt.Println(z())
// }
// package main

// import "fmt"

// func main(){

// 	x := []int {2, 3 , 4 ,5 ,6, }
// 	z := bar(x ...)
// fmt.Println(z)
// 	c := foo(bar, x ...)

// 	fmt.Println(c)
// }

// func bar(s ...int) int {
// 	x := 0
// 	for _, v := range s{
// 		x += v
// 	}
// 	return x
// }

// func foo(g func(s ...int) int, y ...int) int {
// 	total := []int{}
// 	for _, v := range y{
// 		if v % 2 == 0{
// 			total = append(total, v)
// 		}
// 	}
// 	return g(total...)
// }

// closure keeping a variable to a certian scope of a block
// package main

// import "fmt"

// func foo() func() int{
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }
// func main(){
// 	x := foo()
// 	g := foo()
// 	fmt.Print(x())
// 	fmt.Print(x())
// 	fmt.Print(x())
// 	fmt.Print(x())
// 	fmt.Print(x())

// 	fmt.Println(g())
// 	fmt.Println(g())
// 	fmt.Println(g())
// 	fmt.Println(g())
// 	fmt.Println(g())
// }

package main 

import "fmt"



func main() {
	g := func (x []int) int  {
		b := 0
		for _ , v := range x{
			b += v 
		}
		return b
	}

	j := foo(g, []int{2,4,5,6,7,8})
	fmt.Println(j)
}

func foo( h func (x []int) int, y []int ) int {
	 total := h(y)
	 total*=total
	 return total
}
