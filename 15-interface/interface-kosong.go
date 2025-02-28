package main

import (
	"fmt"
	"strings"
)

// Interface kosong atau empty interface yang dinotasikan dengan interface{} atau any, merupakan tipe data yang sangat spesial karena variabel bertipe ini bisa menampung segala jenis data, baik itu numerik, string, bahkan array, pointer, apapun.

func main() {
	var secret interface{}

	secret = "abdul rahem faqih"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data = map[string]interface{}{
		"name": "faqih",
		"grade": 2,
		"breakfast" : []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)

	var secret2 interface{}

	secret2 = 2
	var number = secret2.(int) * 10
	fmt.Println(secret2, "multiplied by 10 is :", number)

	secret2 = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret2.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")


}
