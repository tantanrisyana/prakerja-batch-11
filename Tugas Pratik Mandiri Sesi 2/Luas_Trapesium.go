package main

import (
	"fmt"
)

func main() {
	var tinggi, alas1, alas2 float64

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)

	fmt.Print("Masukkan panjang alas atas: ")
	fmt.Scan(&alas1)

	fmt.Print("Masukkan panjang alas bawah: ")
	fmt.Scan(&alas2)

	luas := 0.5 * (alas1 + alas2) * tinggi

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
