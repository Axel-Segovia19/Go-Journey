// package main

// import "fmt"

// func main() {
// 	for x := 1; x <= 10000; x++ {
// 		fmt.Println(x)
// 	}
// }

// package main

// import "fmt"

// func main () {
// 	for i := 65; i <= 90; i++{
// 		fmt.Println(i)
// 		for n := 1; n <= 3; n++ {
// 			fmt.Printf("\t%#U\n", i)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main () {
// 	x := 1995
// 	for x <= 2022 {
// 		fmt.Println(x)
// 		x++
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	x := 1995
// 	for {
// 		if x > 2022 {
// 			break
// 		}
// 			fmt.Println(x)
// 			x++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// x := 10
// 	// for {
// 	// 	x++
// 	// 	if x > 100 {
// 	// 		break
// 	// 	}
// 	// 	if x % 4 == 0 {
// 	// 		fmt.Println(x)
// 	// 		continue
// 	// 	}
// 	// }
// 	for x := 10; x <=100; x++ {
// 			fmt.Printf("when %v is divided by 4 the remainder is %v\n",x, x%4)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	x := "Bob"
// 	if x == "BOB" {
// 		fmt.Println("You sure?")
// 	} else if x == "Bob" {
// 		fmt.Println(" Correct!")
// 	} else {
// 		fmt.Println("Whatver ")
// 	}

// }

package main 

import "fmt"

func main() {
	favSport := "Soccer"
	switch favSport {
	case "Soccer": fmt.Println("Soccer")
	case "Golf": fmt.Println("golf")
	default: fmt.Println("Whatver ")
	}
}