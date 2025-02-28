package main

import "fmt"

// Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain.
type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "Faqih"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("age : ", s1.person.age)
	fmt.Println("grade : ", s1.grade)
}
