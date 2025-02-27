package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Printf("original number: %v\n", numbers)
    fmt.Printf("filtered number: %v\n", newNumbers)
}
