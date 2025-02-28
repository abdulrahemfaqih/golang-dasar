package main

import "fmt"

// Method pointer adalah method yang dimana variabel objek pemilik method tersebut adalah berbentuk pointer. Kelebihan method jenis ini adalah ketika kita melakukan manipulasi nilai pada property lain yang masih satu struct, nilai pada property tersebut bisa diubah di-level reference-nya
type studeng struct {
	name string
	age  int
}

func (s studeng) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to ", name)
	s.name = name
}

func (s *studeng) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to ", name)
	s.name = name
}

func main() {
	var s1 = studeng{"faqih", 10}
	fmt.Println("s1 before ", s1.name)

	s1.changeName1("lisa")
	fmt.Println("s1 after changeName1 ", s1.name)
	s1.changeName2("alksjdh")
	fmt.Println("s1 after changeName2 ", s1.name)
}
