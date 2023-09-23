package main

import (
	"fmt"
	"math"
)

// Struct Student yang memiliki properti Name dan Score.
type Student struct {
	Name  string
	Score int
}

func main() {
	// Membuat array dari objek Student dengan data yang diberikan.
	students := []Student{
		{"Rizky", 80},
		{"Kobar", 75},
		{"Ismail", 70},
		{"Umam", 75},
		{"Yopan", 60},
	}

	// Menghitung rata-rata skor siswa.
	averageScore := calculateAverageScore(students)

	// Mencari siswa dengan skor terendah dan tertinggi.
	minStudent, maxStudent := findMinAndMaxScoreStudent(students)

	// Menampilkan hasil.
	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Min Score of Students: %s(%d)\n", minStudent.Name, minStudent.Score)
	fmt.Printf("Max Score of Students: %s(%d)\n", maxStudent.Name, maxStudent.Score)
}

// Fungsi untuk menghitung rata-rata skor siswa.
func calculateAverageScore(students []Student) float64 {
	totalScore := 0
	for _, student := range students {
		totalScore += student.Score
	}
	return float64(totalScore) / float64(len(students))
}

// Fungsi untuk mencari siswa dengan skor terendah dan tertinggi.
func findMinAndMaxScoreStudent(students []Student) (minStudent Student, maxStudent Student) {
	minScore := math.MaxInt
	maxScore := math.MinInt

	for _, student := range students {
		if student.Score < minScore {
			minScore = student.Score
			minStudent = student
		}
		if student.Score > maxScore {
			maxScore = student.Score
			maxStudent = student
		}
	}

	return minStudent, maxStudent
}
