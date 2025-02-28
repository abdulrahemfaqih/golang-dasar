package main

import (
	f "fmt"
	"golang-dasar/14-exported-unexported/library"
)

func main() {
	library.SayHello("faqih")
	var s1 = library.Student{Name: "faqih", Grade: 2}
	f.Println("name ", s1.Name)
	f.Println("grade ", s1.Grade)

}
