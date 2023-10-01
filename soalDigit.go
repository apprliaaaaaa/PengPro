package main

import (
	"fmt"
	"strconv"
)

func main() {
	var hasil, d1, d2, d3 int
	fmt.Println()
	fmt.Print("Masukan angka = ")
	fmt.Scan(&hasil)
	d1 = (hasil % 1000) / 100
	d2 = (hasil % 100) / 10
	d3 = hasil % 10

	fmt.Println("d1 = " + strconv.Itoa(d1) + ", d2 = " + strconv.Itoa(d2) + ", d3 = " + strconv.Itoa(d3))
}
