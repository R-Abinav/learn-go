package main

import (
	"fmt"
)

func main(){
	var intSlice []int = []int{1,2,3,4,5}
	fmt.Println(sumSlice[int](intSlice))

	var floatSlice []float32 = []float32{2,3,4}
	fmt.Println(sumSlice[float32](floatSlice))

	var nonEmptySlice []string = []string{"hello", "bigman"}
	isEmpty[string](nonEmptySlice)

	var emptySlice []int = []int{}
	isEmpty[int](emptySlice)
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T){
	if len(slice) == 0{
		fmt.Println("True, slice is empty")
	}else{
		fmt.Println("False, slice is not empty")
	}
}