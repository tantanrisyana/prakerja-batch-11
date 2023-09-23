package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := append(arrayA, arrayB...)
	return RemoveDuplicates(mergedArray)
}

func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	//test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	//[king devil jin akuma eddie steve geese]
	fmt.Println(ArrayMerge([]string{"sergie", "jin"}, []string{"jin", "steve", "bryan"}))
	//[sergie jin steve bryan]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	//[alisa yoshimitsu devil jin law]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//[hwoarang]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	//[]

}
