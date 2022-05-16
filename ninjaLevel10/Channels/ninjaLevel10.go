//anonymous function assigning 42 to a channel and making it run
// package main

// import "fmt"

// func main() {
// 	//c := make(chane int, 1) <- this is a buffering channel
// 	c := make(chan int)

// 	go func ()  {// if you use a buffering channel you will not need this go func
// 		c <- 42
// 	}()

// 	fmt.Println(<-c)
// }

// package main

// import "fmt"

// func main(){
// 	cs := make(chan int) // this is how you make a channel

// 	go func ()  { // this is how you can initiate a channel without using a buffer
// 		cs <- 42 // remember a channel blocks so if you hvae code that wont run first make sure the channel has the ports open for it to run
// 	}()

// 	fmt.Println(<-cs)
// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)// this is now a type chan int
// }

//ranging through values from another channel
// package main

// import "fmt"

// func main(){
// 	//send to channel
// 	c := gen()
// 	// receive from send and assign it to a variable
// 	receive(c)

// 	fmt.Println("about to exit ")

// }
// //receive func that is ranging over the value
// func receive(c <-chan int){
// 	for v := range c{
// 		fmt.Println("value", v)
// 	}
// }

// //send func running a goroutine
// func gen() <-chan int { // func w/ no params but returns a receive chan int type
// 	c := make(chan int)
// 	go func () {
// 		for i := 0; i<100;i++{
// 			c <-i
// 		}
// 		close(c)// closes the channel after the value so it doesnt keep runing in the bg
// 		}()
// 		return c // returning a chan type in
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	q := make(chan int)
// 	c := gen(q) // c is calling the gen func with q as a param

// 	receive(c, q)//receive func is passing 2 param

// 	fmt.Println("about to exit")
// }

// func receive(c, q <-chan int){// taking c which is = to 1-99 and q = 0 as a param
// 	for { // for loop that goes on forever unless broken or retuned out of
// 		select{ // gets values by case goes through each one until loop ends and conditions meet
// 		case v:= <- c:// assigning v to c
// 			fmt.Println("value", v)
// 		case <- q: // only need to return q since it only has 1 value
// 			return
// 		}
// 	}
// }
// func gen(q chan<- int) <-chan int{// this func is taking q as a param but also using its channel to return a value to that channel pf type chan in
// 	c := make(chan int)

// 	go func(){

// 		for i:= 0; i<100; i++{
// 			c <- i
// 		}
// 		q <- 0 // assigning 0 to q and returning 1 - 99 to variable c
// 		close (c)
// 		}()
// 		return c
// }

//comma ok idiom to verify channels are closed

// package main

// import "fmt"

// func main(){
// 	c := make(chan int)

// 	go func ()  { // literall or anonymous function
// 		c <- 42
// 	}()// rmemeber to call it after by addng () at the end

// 	v, ok := <- c// returns value and boolean letting you know if channel is open for true and closed for false
// 	fmt.Println(v, ok)
// 	close(c)
// 	v, ok = <- c
// 	fmt.Println(v, ok) // returns 0, false due to channel closed and no value passed through it
// }

//write a program that prints 100 numbers to a channel pull the numbers off that channel and print them

//example 1
// package main

// import "fmt"
// func main(){
// 	c := make(chan int)

// 	go func ()  {
// 		for i :=0; i < 10; i++{
// 			c <- i
// 		}
// 		close(c)
// 	}()
// 	for v := range c{
// 		fmt.Println("numbers in ch c ", v)
// 	}
// }

//example 2

// package main

// import "fmt"

// func main(){
// 	c := make(chan int)
// 	g := send(c)
// 	close(c)// closing c to not allow this channel to remain open to reduce usage
// 	receive(g) // passing g through as an argument since g is passing c and running it through its function

// 	v, ok := <-c
// 	fmt.Println(v,ok)// okay idiom to check if channel is open throws error and a block if not closed

// }

// func send(c chan<- int) <-chan int{
// 	g := make(chan int)
// 	go func ()  {
// 		for i := 0; i <10; i++{
// 			g <- i
// 		}
// 		close(g)
// 	}()
// 	return g
// }
// func receive(g <-chan int){
// 	for v := range g{
// 		fmt.Println(v)
// 	}
// }

// write a program that launches 10 go routines each go routine adds 10 numbers
// to a channel pull the numbers and print them

// package main 

// import "fmt"


// func main(){
// 	c := make(chan int)

// 	go func ()  {
// 		for i:=0;i<10;i++{ // nested for loop and go function allowing to send multiple go routines per loop 
// 			go func ()  {
// 				for j:=0;j<10;j++{
// 					c <- j
// 				}
// 			}()
// 		}
// 		}()
// 		for k:=0;k<100;k++{ // another for loop verifying there are 10 go routines sending 10 numbers
// 			fmt.Println(k, <-c)
// 		}
// 		close(c)
// }