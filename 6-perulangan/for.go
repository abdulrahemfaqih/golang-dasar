package main

import (
	"fmt"
)

func main() {
	// perulangan keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("angka --> ", i+1)
	}
	fmt.Println()

	i := 0
	// perulangan keyword for dengan argumen hanya kondisi
	for i < 9 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println()
	// keyword for tanpa argumen
	i2 := 0
	for {
		fmt.Println("Angka", i2)
		i2++
		if i2 == 5 {
			break
		}
	}

	// perulangan keyword for dengan range
	for i := range 5 {
		fmt.Println("ANGKA --> ", i)
	}
	fmt.Println()

	// menampilkan char
	xs := "123"
	for i, v := range xs {
		fmt.Printf("Indeks = %d, Value = %d, Character = %c\n", i, v, v)
	}

	fmt.Println()
	ys := [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value = ", v)
	}

	fmt.Println()
	zs := ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value = ", v)
	}

	fmt.Println()
	kvs := map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Printf("keys = %c value = %d\n", k, v)
	}

	fmt.Println()
	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	fmt.Printf("\nperulangan bersarang\n")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}


	fmt.Println()
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("Matriks [", i, "][", j, "]", "\n")
		}
	}
}
