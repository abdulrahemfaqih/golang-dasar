package main

import "fmt"

// go tidak mengadopsu class, tapi ada tipe data struktur Struct
// Struct adalah kumpulan definis variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu, property dalam struct, tipe datanya bisa bervariasi, mirip seperti map, hanya saya key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda

// deklarasi
type student struct {
	name string
	grade int
}

// tag property dalam struct
//Tag merupakan informasi opsional yang bisa ditambahkan pada property struct.
type person struct {
	name string `tag1`
	age int `tag2`
}


func main() {
	var s1 student
	s1.name = "Abdul Rahem Faqih"
	s1.grade = 2

	fmt.Println("Name", s1.name)
	fmt.Println("Grade", s1.grade)

	var s2 = student{}
	s2.name = "Faqih"

	var s3 = student{"ethan", 4}
	var s4 = student{name: "faqih"}

	fmt.Println("Student 2 : ", s2.name)
	fmt.Println("Student 3 : ", s3.name)
	fmt.Println("Student 3 grade : ", s3.grade)
	fmt.Println("Student 4 : ", s4.name)






}

