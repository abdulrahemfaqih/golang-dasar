package main

import (
	"fmt"
	"strings"
)

func main() {
	avg := calculate(1,2,3,4,5,6,7,8,9)
	var msg = fmt.Sprintf("rata rata = %.2f", avg)
	fmt.Println(msg)

	// jika slice menajdi parameter variadic

	numbers := []int{1,2,3,4,5,6,7,8,9}
	avg2 := calculate(numbers...)
	fmt.Println(avg2)

	hobbies := []string{"berenang", "lari"}
	yourHobbies("faqih", hobbies...)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// variadic bisa dikombinasikan dengan paraneter baisa, tteapi diletakkan di belakang
func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")
	fmt.Printf("Hello my name :  %s  \n", name)
	fmt.Printf("My hobbies are : %s \n", hobbiesAsString)
}
