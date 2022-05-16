// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type person struct{
// 	First string
// 	Last string
// 	Sayings []string
// }
// func main(){
// 	p1 := person{
// 		First: "Axel",
// 		Last: "Segovia",
// 		Sayings: []string{"Hello", "I love Dayerlin", "This is pretty cool"},
// 	}
// 	// checking that err is not getting missed and launches an error if its caught
// 		bs, err := json.Marshal(p1)
// 		if err != nil{
// 		 log.Fatalln("Error with code converting to Json")
// 		}
// 		fmt.Println(string(bs))
// }

// package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"log"
// )

// type person struct {
// 	First   string
// 	Last    string
// 	Sayings []string
// }

// func main() {
// 	p1 := person{
// 		First:   "James",
// 		Last:    "Bond",
// 		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
// 	}

// 	bs, err := toJSON(p1)
// 	if err != nil{
// 		log.Fatalln("Error with converting to json")
// 	}
// 	fmt.Println(string(bs))

// }

// // toJSON needs to return an error also
// func toJSON(a interface{}) ([]byte, error) {
// 	bs, err := json.Marshal(a)
// 	if err != nil{
// 	 return []byte{}, errors.New("There is an error with json please check")
// 	}
// 	return bs, nil // make sure to return nil if its asking to return an error if none are found.
// }

// package main

// import "fmt"


// type customErr struct{
// 	info string // this will allow us to save errors to structs and pull the info to print
// }

// func (cs customErr) Error()string{
// 	return fmt.Sprintf("Error has occured: %v", cs.info)//sprintF coverts into a string that then helps get returned as an error
// }

// func main(){
// 	x := customErr{
// 		info: "Hey there! this is your error from customErr",
// 	}
// 	foo(x)
// }

// func foo(e error){
// 	fmt.Println("Foo ran and also ", e)
// }
// package main

// import (
// 	"fmt"
// 	"log"
// )

// type sqrtError struct {
// 	lat  string
// 	long string
// 	err  error
// }

// func (se sqrtError) Error() string {
// 	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
// }

// func main() {
// 	_, err := sqrt(-10.23)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		e := fmt.Errorf("There is an error with %v", f)
// 		return f, e
// 	}
// 	return 42, nil
// }