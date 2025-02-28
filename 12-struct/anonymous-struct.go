package main

import "fmt"

// Anonymous struct adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan objek. Teknik ini cukup efisien digunakan pada use case pembuatan variabel objek yang struct-nya hanya dipakai sekali.

type person struct {
	name string
	age  int
}

func main() {
	// anonymous struct tanpa pengisian property
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name : ", s1.person.name)
	fmt.Println("age : ", s1.person.age)
	fmt.Println("grade : ", s1.grade)

	// anonympus struct dengan opengisian property
	var s2 = struct {
		person
		grade int
	}{
		person: person{"lisa", 21},
		grade: 2,
	}

	fmt.Println("name s2", s2.person.name)




}
