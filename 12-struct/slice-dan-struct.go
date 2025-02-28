package main

import "fmt"

// kombinasi slice dan struct

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []person{
		{name: "faqih", age: 21},
		{name: "lisa", age: 31},
		{name: "febi", age: 21},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}


	fmt.Println("")
	// inisialisasi slice anonymous struct

	var allMurid = []struct {
		person
		grade int
	}{
		{person: person{"fares", 21}, grade: 1},
		{person: person{"farish", 21}, grade: 2},
		{person: person{"yunus", 21}, grade: 2},
	}

	for _, murid := range allMurid {
		fmt.Println(murid)
	}


	// Deklarasi Anonymous Struct Menggunakan Keyword var
	var pengemudi struct {
		person
		lisensi string
	}

	pengemudi.person = person{"lisa", 21}
	pengemudi.lisensi = "C"
	
}
