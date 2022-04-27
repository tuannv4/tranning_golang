package main

import "fmt"

type person struct {
	firstName string
	latsName  string
	contact   contactInfor
}

type contactInfor struct {
	email   string
	zipCode int
}

func main() {
	// 	var alex person
	// 	alex.firstName = "xxx"
	// 	alex.latsName = "yyy"
	// 	fmt.Println(alex)
	// 	fmt.Printf("%+v", alex)
	// }

	jim := person{
		firstName: "jim",
		latsName:  "party",
		contact: contactInfor{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimPointer := &jim
	// jimPointer.updateName("tuan")
	fmt.Println(jimPointer)
	// }

	// func (p person) print() {
	// 	fmt.Printf("%+v", p)
	// }

	// func (pointerToPerson *person) updateName(newFirstName string) {
	// 	pointerToPerson.firstName = newFirstName
	// }

	// mySlice := []string{"hi", "There", "How", "are", "you"}

	// updateSlice(mySlice)
	// fmt.Println(mySlice)
}

// func updateSlice(s []string) {
// 	s[0] = "Bye"
// }
