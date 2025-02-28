package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAll(i int) string {
	return strings.Split(s.name, " ")[i-1]
}


func main() {
	var s1 = student{"abdul rahem faqih", 21}
	s1.sayHello()

	var name = s1.getNameAll(2)
	fmt.Println("potongan nama : ", name)
}
