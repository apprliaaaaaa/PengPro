package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	var hasil float64

	fmt.Println("Soal No 3")
	fmt.Print("Masukan nilai X = ")
	fmt.Scan(&X)
	hasil = (math.Pow(X, 3) + (3 * X)) / (math.Pow(X, 4) - (3 * math.Pow(X, 2)) + 4)
	fmt.Println(hasil)
}
