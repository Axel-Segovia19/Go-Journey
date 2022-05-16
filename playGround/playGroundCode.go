// package main

// import "fmt"

// type dog struct {
// 	First     string
// 	FavToy    []string
// 	FavPerson map[int]string
// }

// func main() {
// 	d1 := dog{
// 		First:     "Rosie",
// 		FavToy:    []string{"Bone", "Bed", "Carrots"},
// 		FavPerson: map[int]string{1: "Dayerlin", 2: "Axel", 3: "Simba", 4: "Ziggy"},
// 	}
// 	d2 := dog{
// 		First:     "Ziggy",
// 		FavToy:    []string{"Tennis Ball", "String", "Carrots"},
// 		FavPerson: map[int]string{1: "Axel", 2: "Dayerlin", 3: "Rosie", 4: "Simba"},
// 	}
// 	displayInfo(d1)
// 	fmt.Println("--------")
// 	displayInfo(d2)

// }

// func displayInfo(g dog) {
// 	fmt.Println(g.First)
// 	for i, v := range g.FavToy{
// 		fmt.Println("Favorite toy", i+1, v )
// 	}
// 	for k, v := range g.FavPerson{
// 		fmt.Println("Favorite person by rank",k, v)
// 	}

// }

package main

import "fmt"

// you need to make the slice of int return the target indexes of the number that sum to the target number
func main (){
	b := []int{2, 7, 10, 11}
	target := 9
	twoNum(b, target)
}

func twoNum(n []int,target int) []int{
	var x []int
	for _, v := range n{
		for j, v2 := range n{
			if v + v2 == target{
				x = append(x, j )
			}
		}
	}
	fmt.Println(x)
		return x
}