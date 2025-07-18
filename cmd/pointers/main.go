package main

import (
	"fmt"
)

func main(){
	var myPtr *int32 
	fmt.Printf("\nThe value of the ptr is : %v", myPtr)

	var myPtr2 *int32 = new(int32)
	fmt.Printf("\nThe value of the ptr is : %v", myPtr2)

	var i int32 = 0
	fmt.Printf("\nThe value of i is: %v", i)

	var ptrToi *int32 = &i
	*ptrToi = 10

	fmt.Printf("\nThe value of the ptr to i: %v", ptrToi)
	fmt.Printf("\nThe value of i: %v", i)
	fmt.Printf("\nThe defereferenced value of ptr to i: %v", *ptrToi)

	// -- NOTE --
	//slices contain pointers to an underlying array
	var mySlice []int32 = []int32{1,2,3}
	fmt.Printf("\nThe slice before any operations: %v", mySlice)

	var mySliceCopy = mySlice

	mySliceCopy[1] = 10

	fmt.Println("\nOriginal slice: ", mySlice)
	fmt.Println("Copied Slice:", mySliceCopy)

	// -- Pass by reference --
	// You wanna pass by reference if your handling large data -> As pass by value creates copies
	var thing1 [5]int32 = [5]int32{1,2,3,4,5}
	var result [5]int32 = squareThing(&thing1)

	fmt.Printf("The original thing: %v", thing1)
	fmt.Printf("\nThe result is: %v", result)
}

func squareThing(thing2 *[5]int32) [5]int32{
	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i] 
	}
	return *thing2
}