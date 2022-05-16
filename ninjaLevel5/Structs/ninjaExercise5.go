// package main

// import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	favIC []string
// }

// func main() {
// 	p1 := person{
// 		first: "Axel", last: "Segovia", favIC: []string{"rocky road", "Chocolate", "Napolitano"},
// 	}
// 	p2 := person{
// 		first: "Dayerlin", last: "Carrero", favIC: []string{"rocky road", "Chocolate", "Napolitano"},
// 	}
// 	fmt.Println(p1.first, p1.last)
// 	for _, v := range p1.favIC{
// 		fmt.Println("\t", v)
// 	}
// 	fmt.Println(p2.first, p2.last)
// 	for _, v := range p2.favIC{
// 		fmt.Println("\t", v)
// 	}
// package main

// import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	favIC []string
// }
// // range over a map of the struct.
// func main() {
// 	p1 := person{
// 		first: "Axel", last: "Segovia", favIC: []string{"rocky road", "Chocolate", "Napolitano"},
// 	}
// 	p2 := person{
// 		first: "Dayerlin", last: "Carrero", favIC: []string{"rocky road", "Chocolate", "Napolitano"},
// 	}

// 		x := map[string]person{
// 			p1.last: p1,
// 			p2.last:p2,
// 		}
// 		for k, v := range x{
// 					fmt.Println(k)
// 					fmt.Println(v.first)
// 					for _, val := range v.favIC{
// 						fmt.Println("\t",val)
// 					}
// 		}
// 	}

// package main

// import "fmt"

// type vehicle struct{
// 	doors int
// 	color string
// }
// type truck struct{
// 	vehicle
// 	fourWheel bool
// }
// type sedan struct{
// 	vehicle
// 	luxury bool
// }

// func main() {
// 	dodge := truck{
// 		vehicle:
// 		vehicle{ doors: 2, color: "red",},
// 		fourWheel: true,
// 	}
// 	subaru := sedan{
// 		vehicle: vehicle{ doors: 4, color: "blue"},
// 		luxury: true,
// 	}
// 	fmt.Println(dodge)
// 	fmt.Println(subaru)
// 	fmt.Println(dodge.vehicle.doors )
// 	fmt.Println(subaru.vehicle.color)
// }

package main 

import "fmt"



func main() {
	x := struct{
			first string
			friends map[string]int
			favDrinks []string
	}{
		first: "Axel",
		friends: map[string]int {"angel": 1, "chato": 56},
		favDrinks: []string{"coke, rum, beer"},
	}
	for k, v := range x.friends {
		fmt.Println(k, v)
	}
	for _,v := range x.favDrinks{
		fmt.Println(v)
	}
}