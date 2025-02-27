package main

import (
	"fmt"
)

func main() {
	var names [4]string
	names[0] = "abdul"
	names[1] = "faqih"
	names[2] = "rahem"
	names[3] = "sepp"

	fruits := [4]string{"apel", "mangga", "semangka", "pisang"}
	newFruits := [3]string{
		"pisang",
		"kiwi",
		"stawberry",
	}

	fmt.Println("Jumlah elemen \t\t:", len(fruits))
	fmt.Println("isi semua element \t:", fruits)
	fmt.Println("Jumlah elemen \t\t:", len(newFruits))
	fmt.Println("isi semua element \t:", newFruits)

	// inisialisasi dengan jumlah otomatis
	numbers := [...]int{2, 3, 4, 5}
	fmt.Println("Jumlah elemen \t\t:", len(numbers))
	fmt.Println("isi semua element \t:", numbers)

	// array multidimensi
	numbers1 := [2][3]int{{1, 2, 3}, {1, 2, 3}}
	numbers2 := [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println("numbers1\t\t:", numbers1)
	fmt.Println("numbers2\t\t:", numbers2)

	// perulangan
	fmt.Print("\nPERULANGAN ARRAY\n")
	fruits1 := [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("keyword for saja")
	for i := 0; i < len(fruits1); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits1[i])
	}

	fmt.Println("\nkeyword for - range")

	for i, fruit := range fruits1 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}


	// memakai _ untuk menampilkan valuenya saja, kadang kala index tidak perlu di tampilkan
	for _, fruit := range fruits1 {
		fmt.Println("nama buah : ", fruit)
	}

}
