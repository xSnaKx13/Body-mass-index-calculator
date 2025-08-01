package main

import "math"

func calculation(weight, height float64) float64 {
	return weight / math.Pow(height/100, 2)
}
