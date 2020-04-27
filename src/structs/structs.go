package main

import (
	"fmt"
	"hr"
)

func main() {
	a := hr.Employee{
		Person:      hr.Person{Name: "Thanapon", Age: 15},
		Description: "Programmer",
	}

	fmt.Println("a = ", a)		// a =  {{Thanapon 15} Programmer}
	fmt.Println(a.Person.Name)	// Thanapon
	fmt.Println(a.Age)			// 15
}
