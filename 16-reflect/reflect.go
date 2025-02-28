package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

// pengaksesan informasi property variabel objek
func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe data	:", reflectType.Field(i).Type)
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println()
	}

}

func main() {
	// Mencari tipe data dan value menggunakan reflect
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	// pengaksesan nilai dalam bentuk interface()
	fmt.Println("tipe variabel : ", reflectValue.Type())
	fmt.Println("nilai variabel: ", reflectValue.Interface())

	var s1 = &student{Name: "faqih", Grade: 2}
	s1.getPropertyInfo()

}
