package main

import (
	"fmt"
)

func main(){
	var intSlice []int = []int{1,2,3,4,5}
	fmt.Println(sumSlice[int](intSlice))

	var floatSlice []float32 = []float32{2,3,4}
	fmt.Println(sumSlice[float32](floatSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}