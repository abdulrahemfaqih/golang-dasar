package main

import "fmt"

func main() {
	var chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	chicken2 := map[string]int{
		"januari":  1,
		"februari": 2,
		"maret":    3,
	}
	fmt.Println(chicken2)

	// iterasi
	chicken3 := map[string]int{
		"januari":  50,
		"februari": 40,
	}

	for key, val := range chicken3 {
		fmt.Println(key, "\t:", val)
	}

	// deteksi keberadaan item dengan key tertentu
	value, isExist := chicken3["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item not Exist")
	}

	// hapus data
	delete(chicken3, "februari")
	for key, val := range chicken3 {
		fmt.Println(key, "\t:", val)
	}

	// kombinasi slice dan map
	mahasiswa := []map[string]string{
		{
			"nama":     "Abdul Rahem Faqih",
			"prodi":    "Teknik Informtika",
			"fakultas": "Teknik",
		},
		{
			"nama":     "Sitti Nur Khalisa",
			"prodi":    "Sistem Informasi",
			"fakultas": "Teknik",
		},
	}

	for _, m := range mahasiswa {
		fmt.Println(m["nama"], "-", m["prodi"], "-", m["fakultas"])
	}

}
