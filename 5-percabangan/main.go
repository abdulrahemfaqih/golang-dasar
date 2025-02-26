package main

import "fmt"

func main() {
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulua")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)

	}

	// temporary
	var point2 = 8840.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch

	point3 := 3
	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// switch dengan gaya if
	point4 := 6
	switch {
	case point4 == 8:
		fmt.Println("Perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("haha gendeng")
			fmt.Println("Coba lagi ya!")
		}
	}

	// percabangan bersarang
	point5 := 10

	if point5 > 7 {
		switch point5 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")

		}
	} else {
		if point5 == 5 {
			fmt.Println("Not bad!")
		} else if point5 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point5 == 0 {
				fmt.Println("goblok")
			}
		}
	}

}
