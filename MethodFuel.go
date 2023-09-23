package main

import "fmt"

type Car struct {
	Type   string
	fuelIn float64
}

func (c Car) CalculateDistance() float64 {
	JmlBensin := 1.5 // Efisiensi bahan bakar dalam L
	return c.fuelIn / JmlBensin
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		fuelIn: 10.0, // Jumlah bahan bakar dalam liter
	}

	distance := myCar.CalculateDistance()
	fmt.Printf("Perkiraan jarak yang bisa ditempuh oleh %s adalah %.2f mil\n", myCar.Type, distance)
}
