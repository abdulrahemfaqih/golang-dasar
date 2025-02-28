package main

// Teknik ini biasa digunakan ketika decoding data JSON yang struktur datanya cukup kompleks dengan proses decode hanya sekali.

type student struct {
	person struct {
		name string
		age int
	}
	grade int
	hobbies []string
}
