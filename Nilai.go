package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// Inisialisasi min dan max dengan nilai pertama dalam slice.
	min = *numbers[0]
	max = *numbers[0]

	// Iterasi melalui angka-angka input menggunakan pointer.
	for _, numPtr := range numbers {
		num := *numPtr
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Print("Input angka ke-1: ")
	fmt.Scan(&a1)
	fmt.Print("Input angka ke-2: ")
	fmt.Scan(&a2)
	fmt.Print("Input angka ke-3: ")
	fmt.Scan(&a3)
	fmt.Print("Input angka ke-4: ")
	fmt.Scan(&a4)
	fmt.Print("Input angka ke-5: ")
	fmt.Scan(&a5)
	fmt.Print("Input angka ke-6: ")
	fmt.Scan(&a6)

	// Menggunakan pointer untuk mengirim angka ke fungsi getMinMax.
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("Nilai min: %d\n", min)
	fmt.Printf("Nilai max: %d\n", max)
}
