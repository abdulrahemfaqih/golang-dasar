package main

import (
	"fmt"
	"math"
)

func main() {
	diameter := 15
	area1, circumference1 := calculate1(float64(diameter))
	fmt.Printf("Luas Lingkaran \t\t: %.2f \n", area1)
	fmt.Printf("Keliling Lingkaran \t: %.2f \n", circumference1)
	area2, circumference2 := calculate2(float64(diameter))
	fmt.Printf("Luas Lingkaran \t\t: %.2f \n", area2)
	fmt.Printf("Keliling Lingkaran \t: %.2f \n", circumference2)
}

func calculate1(d float64) (float64, float64) {
	// hitung luas
	area1 := math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	circumference1 := math.Pi * d

	return area1, circumference1
}

// Fungsi Dengan Predefined Return Value
func calculate2(d float64) (area2 float64, circumference2 float64) {
	// hitung luas
	area2 = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	circumference2 = math.Pi * d

	return
}
