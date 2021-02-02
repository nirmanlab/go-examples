package main

import (
	"fmt"
	"gopkgdemo/user"
)

func main() {
	u := user.CreateUser("Sagar", "jakatnakat")

	s := user.CreateStudent(*u, 10)
	fmt.Printf("Student name :- %v \n", s.U.Name)
	fmt.Printf("Student name :- %v \n", s.U.Address)
	fmt.Printf("Student name :- %v \n", s.StudentID)
}
