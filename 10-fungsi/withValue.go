package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fungsi dengan return value

var randomMizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var randomValue int
	randomValue = randWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	fmt.Println("\nPembagian biasa\n")
	devideNumber(10,7)
	devideNumber(10,0)
}

func randWithRange(min, max int) int {
	value := randomMizer.Int()%(max-min+1) + min
	return value
}

func devideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider, %d cannot divided by %d\n", m, n)
		return
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
