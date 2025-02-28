package main

import (
	"fmt"
)

func main() {
	// Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe data ketika deklarasi.
	var number *int
	fmt.Println(number)

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("NumberA (value) : ", numberA)
	fmt.Println("NumberA (address) : ", &numberA)
	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

	fmt.Println("")

	numberA = 5
	fmt.Println("NumberA (value) : ", numberA)
	fmt.Println("NumberA (address) : ", &numberA)
	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

	var numberC = 4
	fmt.Println("before :", numberC)
	change(&numberC, 10)
	fmt.Println("after : ", numberC)
}

func change(original *int, value int) {
	*original = value
}
