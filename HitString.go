package main

import "fmt"

func Mapping(slice []string) map[string]int {
	counts := make(map[string]int)

	for _, item := range slice {
		counts[item]++
	}

	return counts
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	// map[adi:1 asd:2 qwe:3]

	fmt.Println(Mapping([]string{}))
	// map[]
}
