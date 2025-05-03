package main

import (
	"fmt"
	"math"
)

func volumeOfSphere(radius float64) float64 {
	return math.Pi * (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
}

func main() {
	sphereVolume := volumeOfSphere(5.0)
	fmt.Printf("Объём шара: %v m^3", sphereVolume)
}
