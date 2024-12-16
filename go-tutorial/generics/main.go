package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

func SumInts(nums map[string]int64) int64 {
	var sum int64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumFloats(nums map[string]float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumIntsOrFloats[K comparable, V int64 | float64](nums map[K]V) V {
	var sum V

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumNumbers[K comparable, V Number](nums map[K]V) V {
	var sum V

	for _, num := range nums {
		sum += num
	}

	return sum
}
