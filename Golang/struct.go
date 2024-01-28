// package main

// import (
// 	"fmt"
// )

// // struct is a collection of user defined data
// type department struct {
// 	name     string
// 	roll_no  string
// 	stu_name string
// }

// // embedded struct example
// type user struct {
// 	new_user string
// 	college
// }

// // nested struct example
// type college struct {
// 	name_of_college    string
// 	area               string
// 	dep                string
// 	name_of_department department
// }

// func gotest(nc user) {
// 	fmt.Printf(" name: %v \n roll_no: %v\n Your assigned college %v\n Your department %v\n Area of college %v\n new user %v\n ", nc.name_of_department.stu_name, nc.name_of_department.roll_no, nc.name_of_college, nc.name_of_department.name, nc.area, nc.new_user)
// }

// func main() {
// 	gotest(
// 		user{new_user: "kesavan",
// 			college: college{name_of_college: "tce",
// 				area: "Madurai",
// 				name_of_department: department{
// 					name:     "EEE",
// 					stu_name: "tharun",
// 					roll_no:  "20e099",
// 				}}})
// }