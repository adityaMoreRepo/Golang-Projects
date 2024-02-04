package main

import "fmt"

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }
// type contactInfo struct {
// 	email   string
// 	phoneNo string
// 	zip     int
// }

type petorlEngine struct {
	KMPL int
	fuel int
}

type electricEngine struct {
	KMPC   int
	charge int
}

type engine interface {
	KMsToGO() int
}

func (pe petorlEngine) KMsToGO() int {
	return pe.fuel * pe.KMPL
}

func (ee electricEngine) KMsToGO() int {
	return ee.KMPC * ee.charge
}

func willMakeIT(e engine, kms int) {
	if e.KMsToGO() <= kms {
		fmt.Println("You won't be able to make it")
	} else {
		fmt.Println("You will make it")
	}

}

func main() {

	//We can init the struct variable:
	// var Alex person
	//if we have embedded another struct we can also access it's members if it is directly embedded.
	// Alex.firstName = "Alex"
	// Alex.lastName = "Ford"
	// Alex.email = "alex@gmail.com"
	//We can also assign like this
	// Alex.contactInfo = contactInfo{
	//      email:   "alex@gmail.com",
	// 		phoneNo: "3858392051",
	// 		zip:     41511,
	// }
	// Aditya := person{
	// 	firstName: "Aditya",
	// 	lastName:  "More",
	// 	contactInfo: contactInfo{
	// 		email:   "adimore@gmail.com",
	// 		phoneNo: "3858392051",
	// 		zip:     41511,
	// 	},
	// }

	// fmt.Printf("Person : %+v\n", Aditya)
	// fmt.Printf("Email : %v\n", Aditya.email) // when Person is embedded with contact struct directly
	// fmt.Printf("Email : %v\n", Aditya.email)
	// Aditya.setContact(contactInfo{
	// 	email:   "adimore@gmail.com",
	// 	phoneNo: "8984212007",
	// 	zip:     411010,
	// })

	// fmt.Printf("Person : %+v\n", Aditya)

	//test for subset of string
	// str := "ABC"
	// curr := ""
	// getSubString(str, curr, 0)
	myPetrolEngine := petorlEngine{
		KMPL: 30,
		fuel: 12,
	}
	myElectricEngine := electricEngine{
		KMPC:   40,
		charge: 15,
	}
	willMakeIT(myPetrolEngine, 500)
	willMakeIT(myElectricEngine, 500)
}

// func (p *person) setContact(contact contactInfo) {
// 	p.contactInfo = contact
// }

// func getSubString(str string, curr string, index int) {
// 	if index == len(str) {
// 		fmt.Printf("%v ", curr)
// 		return
// 	}
// 	getSubString(str, curr, index+1)
// 	getSubString(str, curr+string(str[index]), index+1)
// }
