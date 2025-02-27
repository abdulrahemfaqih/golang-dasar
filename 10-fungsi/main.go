package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"faqih", "lisa"}
	printMessage("halo", names)

}

func printMessage(message string, arr []string) {
	namesString := strings.Join(arr, " dan ")
	fmt.Println(message, namesString)
}



