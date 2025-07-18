package main

import (
	"fmt"
)

func main(){
	var myPtr *int32 
	fmt.Printf("\n The value of the ptr is : %v", myPtr)

	var myPtr2 *int32 = new(int32)
	fmt.Printf("\n The value of the ptr is : %v", myPtr2)

	var i int32 = 0
	fmt.Printf("\n The value of i is: %v", i)

	var ptrToi *int32 = &i
	*ptrToi = 10

	fmt.Printf("\n The value of the ptr to i: %v", ptrToi)
	fmt.Printf("\n The value of i: %v", i)
	fmt.Printf("\n The defereferenced value of ptr to i: %v", *ptrToi)

	// -- NOTE --
	//slices contain pointers to an underlying array
	var mySlice []int32 = []int32{1,2,3}
	var mySliceCopy = mySlice

	mySliceCopy[1] = 10

	fmt.Println("\nOriginal slice: ", mySlice)
	fmt.Println("\nCopied Slice:", mySliceCopy)

	// -- Pass by reference --
	// You wanna pass by reference if your handling large data -> As pass by value creates copies
	var thing1 [5]int32 = [5]int32{1,2,3,4,5}
	var result [5]int32 = squareThing(&thing1)

	fmt.Printf("\nThe original thing: %v", thing1)
	fmt.Printf("\nThe result is: %v", result)
}

func squareThing(thing2 *[5]int32) [5]int32{
	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i] 
	}
	return *thing2
}