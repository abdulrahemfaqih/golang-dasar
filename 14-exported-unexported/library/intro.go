package library

import "fmt"

/**
GO mengenal 2 jenis level akses atau hak akses:
	- hak akses exported atau public,menandakan bahwa komponen boleh untuk diakses dari package lain
	- hak akses unexported atau private, Berarti komponen hanya bisa diakses dari file yang package-nya sama, bisa dalam satu file yang sama atau di file berbeda yang masih 1 folder yang package-nya pastinya sama
**/

type Student struct {
	Name string
	Grade int
}

func SayHello(name string) {
	fmt.Println("helo")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya : ", name)
}
