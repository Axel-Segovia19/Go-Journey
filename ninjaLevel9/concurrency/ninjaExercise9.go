// allowing go code to be waited on gor go routines to finish and close. allows to run go concurrently
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wg sync.WaitGroup
// func main(){
// 	fmt.Println("begining ",runtime.NumCPU())
// 	fmt.Println("begining go routines",runtime.NumGoroutine())
// 	wg.Add(2)
// 	go foo()
// 	go bar()
// 	mer()
// 	fmt.Println("mid ",runtime.NumCPU())
// 	fmt.Println("mid go routines",runtime.NumGoroutine())
// 	wg.Wait()
// 	fmt.Println("end ",runtime.NumCPU())
// 	fmt.Println("end go routines",runtime.NumGoroutine())
// }

// func foo(){
// 	fmt.Println("This is Foo")
// 	wg.Done()
// }

// func bar(){
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// 	wg.Done()
// }

// func mer(){
// 	fmt.Println("This is Mer")
// }

// remember a pointer receiver will only allow a pointer value to be passed through
// package main

// import "fmt"

// type person struct{
// 	fn string
// 	ln string
// }

// type human interface{
// 	speak()
// }

// func (p *person) speak() {
// 	fmt.Println("hello" , p.fn, p.ln)
// }

// func saySomething(h human) {
// 	h.speak()
// }

// func main(){
// 	p1 := person{ fn: "Axel", ln: "Segovia",}

// 	saySomething(&p1)

// }

// confirming race condition and doing fnctions to avoide that from happenning mutex, atomic
// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// 	//"runtime"
// 	"sync"
// )

// func main() {
// 	// var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	var counter int64
// 	//counter := 0
// 	 gs := 100

// 		wg.Add(gs)

// 	 for i := 0; i < gs; i++ {
// 			go func ()  {
// 				// mu.Lock()
// 				atomic.AddInt64(&counter, 1)

// 				//runtime.Gosched() remove runtime schedule when using mutex snce you are locking and unlocking the routine

// 				fmt.Println(atomic.LoadInt64(&counter))
// 				// mu.Unlock()
// 				wg.Done()
// 			}()
// 		}
// 			wg.Wait()

// 		fmt.Println("end", counter)

// }

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}