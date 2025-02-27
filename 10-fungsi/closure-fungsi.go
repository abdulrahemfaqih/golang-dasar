package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
	// digunakan untuk menampilkan data tanpa melihat tipe datanya. Jadi bisa digunakan untuk menampilkan data array, int, float, bool, dan lainnya. Bisa dilihat di contoh statement, data bertipe array dan numerik ditampilkan menggunakan %v
}
