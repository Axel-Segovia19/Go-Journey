//turning go code into json
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	First string
// 	Age   int
// }

// func main(){
// 	u1 := User{
// 		First: "James",
// 		Age:   32,
// 	}

// 	u2 := User{
// 		First: "Moneypenny",
// 		Age:   27,
// 	}

// 	u3 := User{
// 		First: "M",
// 		Age:   54,
// 	}

// 	users := []User{u1, u2, u3}

// 	fmt.Println(users)

// 	bs, err := json.Marshal(users)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(string(bs))

// }

// how to receive code from json and convert it to go
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type person struct {
// 	First   string   `json:"First"`
// 	Last    string   `json:"Last"`
// 	Age     int      `json:"Age"`
// 	Sayings []string `json:"Sayings"`
// }

// func main(){
// 	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
// 	bs := []byte(s)

// 	people := []person{}
// 	err := json.Unmarshal(bs, &people)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for i, person := range people {
// 		fmt.Println("person #", i+1)
// 		fmt.Println(person.First, person.Last, person.Age)
// 		for _, saying := range person.Sayings{
// 			fmt.Println("\t\t", saying)
// 		}
// 	}
// }

// encode to json the []user sending the results to stdout //hint use json.newnecoder

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type user struct {
// 	First   string
// 	Last    string
// 	Age     int
// 	Sayings []string
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Last:  "Bond",
// 		Age:   32,
// 		Sayings: []string{
// 			"Shaken, not stirred",
// 			"Youth is no guarantee of innovation",
// 			"In his majesty's royal service",
// 		},
// 	}

// 	u2 := user{
// 		First: "Miss",
// 		Last:  "Moneypenny",
// 		Age:   27,
// 		Sayings: []string{
// 			"James, it is soo good to see you",
// 			"Would you like me to take care of that for you, James?",
// 			"I would really prefer to be a secret agent myself.",
// 		},
// 	}

// 	u3 := user{
// 		First: "M",
// 		Last:  "Hmmmm",
// 		Age:   54,
// 		Sayings: []string{
// 			"Oh, James. You didn't.",
// 			"Dear God, what has James done now?",
// 			"Can someone please tell me where James Bond is?",
// 		},
// 	}

// 	users := []user{u1, u2, u3}

// 	//fmt.Println(users)

// 	err := json.NewEncoder(os.Stdout).Encode(users)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }

// sort this code by int, string for each

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
// 	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
	
// 	sort.Ints(xi)
// 	sort.Strings(xs)
// 	fmt.Println(xi)
// 	// sort xi
// 	fmt.Println(xi)

// 	fmt.Println(xs)
// 	// sort xs
// 	fmt.Println(xs)

// }

// sort by age, name and saying

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	sort.Sort(ByAge(users))
	for _ , u := range users{
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings{
				fmt.Println("\t",v)
		}
	}


	sort.Sort(ByLast(users))
	

	for _ , u := range users{
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings{
				fmt.Println("\t",v)
		}
	}

}
