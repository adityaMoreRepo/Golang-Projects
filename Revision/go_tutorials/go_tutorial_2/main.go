package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	phoneNo string
	zip     int
}

func main() {
	Aditya := person{
		firstName: "Aditya",
		lastName:  "More",
		contactInfo: contactInfo{
			email:   "adimore@gmail.com",
			phoneNo: "3858392051",
			zip:     41511,
		},
	}

	fmt.Printf("Person : %+v\n", Aditya)
	Aditya.setContact(contactInfo{
		email:   "adimore@gmail.com",
		phoneNo: "8984212007",
		zip:     411010,
	})
	fmt.Printf("Person : %+v\n", Aditya)
}

func (p *person) setContact(contact contactInfo) {
	p.contactInfo = contact
}
