// package main

// import "fmt"

// func main() {
// 	x := []int{19, 27, 17, 16, 1} // how to call a slice called a composite literal
// 	for i, v := range x { //loops through each index and value by valling the range
// 		fmt.Println(i,v)
// 	}
// 		fmt.Println(x[0]) //calling the value of the slice by the index
// 		fmt.Println(x[1])
// 		fmt.Println(x[2])
// 		fmt.Println(x[3])
// 		fmt.Println(x[4])
// 		fmt.Println(len(x))// how to find the length of how many items in this slice
// 		fmt.Println(x[1:3])// how to slice a slice by givine a starting point and ending point by the index
// 	for i := 0; i < 5; i++{ // how to loop through a slice without calling the range
// 		fmt.Println(i, x[i])
// 	}

// }

package main 

import "fmt"


func main() {
	x := []int {24, 25, 26}
	fmt.Println(x)
	x = append(x, 27, 28, 29)
	fmt.Println(x)
	y := []int { 300, 400 ,500}
	x = append(x, y...)
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i,v)
	}
	fmt.Println(x[2:6])
	fmt.Println(len(x))

	x = append(x[1:3]) // this is a way to take specific values out of a slice. you are taking away t
	fmt.Println(x) // append lets you modify a slice by deleting or adding into.

	
}