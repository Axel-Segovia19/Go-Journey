// package main

// import "fmt"

// func main() {
// 	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Println(x[1:6])
// 	fmt.Println(x[6:])
// 	fmt.Println(x[3:8])
// 	fmt.Println(x[2:7])
// }
// package main

// import "fmt"

// func main(){
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	x = append(x, 52, 53, 54, 55)
// 	y := []int{56, 57, 58, 59, 60}
// 	x = append(x, y...)
// 	fmt.Println(x)
// }

// package main

// import "fmt"

// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

// 	y := append(x[:3], x[6:] ... )
// 	fmt.Println(y)
// }

// package main

// import "fmt"

// func main() {
// x := make([]string, 0, 50)

// x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,)

// fmt.Println(len(x))
// fmt.Println(cap(x))
// for i := 1; i < 50; i++{
// 	fmt.Println(i, x[i])
// }

// }

// package main

// import "fmt"

// func main() {
// 	x := []string {"James", "Bond", "Shaken, not stirred"}
// 	y := []string {"Miss", "Moneypenny", "Helloooooo, James."}

// 	z := [][]string {x, y}

// 	for i, v := range z {
// 		fmt.Println("Record #", i)
// 		for _, val := range v {
// 			fmt.Println(val)
// 		}
// 	}
// }

package main 

import "fmt"


func main(){
	x := map[string][]string {

`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`}, 
`no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	x["segovia_axel"] = []string{"Soccer", "dayerlin", "dogs"}
	delete(x, "segovia_axel")
	for k, v := range x {
		fmt.Println(k)
		for i , val := range v {
			fmt.Println("\t", i, val)
		}
	}
}