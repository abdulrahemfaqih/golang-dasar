package main

import "fmt"

func main() {
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	cfruits := append(fruits, "pepaya")
	fmt.Println(cfruits)

	bfruits := fruits[0:2]
	fmt.Println(len(bfruits))
	fmt.Println(cap(bfruits))

	// copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	dst1 := []string{"potato", "potato", "potato"}
	src1 := []string{"watermelon", "pinnaple"}
	n1 := copy(dst1, src1)


	fmt.Println(dst1)
	fmt.Println(src1)
	fmt.Println(n1)

}
