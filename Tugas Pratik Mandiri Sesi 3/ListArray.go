package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	counts := make(map[int]int)

	// Menghitung berapa kali setiap angka muncul dalam string
	for _, char := range angka {
		num, _ := strconv.Atoi(string(char))
		counts[num]++
	}

	// Membuat slice untuk menyimpan angka yang hanya muncul sekali
	var result []int
	for _, char := range angka {
		num, _ := strconv.Atoi(string(char))
		if counts[num] == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 0 4]
}
