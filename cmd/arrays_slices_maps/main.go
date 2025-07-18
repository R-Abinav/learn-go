package main

import "fmt"

func main(){
	// ---- ARRAYS ----
	var intArr [3]int32 
	intArr[1] = 1

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) //Access element 1 and 2 like this

	//Print out the memory of each element using -> &
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//Another syntax to immediately initialise the array
	var intArr2 = [3]int32{1,2,3}
	fmt.Println(intArr2)

	//Another syntax
	intArr3 := [5]int32{1,2,3,4,5}
	fmt.Println(intArr3)

	//Can omit the "5" inside the [], as well using this syntax
	intArr4 := [...]int32{1,2,3,4}
	fmt.Println(intArr4)

	// ---- SLICES ----
	//Basically, slice is an array with extended functionality
	//By omiting the length value, you have a 'slice'
	var intSlice = []int32{4,5,6}
	fmt.Println(intSlice)

	//Length -> The number of items present in the slice
	//Capacity -> The number of items the slice can hold

	//Print out the length before
	fmt.Printf("The length of the array before appending is %v and capacity is %v", len(intSlice), cap(intSlice));

	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("\nThe length of the array after appending is %v and capacity is %v", len(intSlice), cap(intSlice));

	// fmt.Println(intSlice[4]) -> Index out of range error, even if capacity is 6.

	//Can insert multiple values
	var intSlice2 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//make function
	//make(data_type[], length, capcity (optional)) -> By default if not specified, the capacity is set to the length of the slice
	//NOTE -> IF you have a rough idea on how values you are going to store, always specify the capacity. This avoids the reallocation time when it doesnt have more space.
	//Reallocation takes a hit on performance
	var intSlice3 []int32 = make([]int32, 3, 8) 
	fmt.Println(intSlice3)

	// ---- MAPS ----
	//A key value pair. Can look up a value by its key.
	//Can create a map, with the same make syntax => map[key]value
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 16, "sarah": 24}
	fmt.Println(myMap2["Adam"])

	//NOTE: A map always returns something, so if you try search something using a key which doesn't exist in the map, then it returns the default value of the data type of the value
	//Example -> myMap2["Jason"] => Returns 0

	var age, ok = myMap2["Jason"]
	if ok{
		fmt.Printf("The age of Jason is %v", age)
	}else{
		fmt.Println("Invalid name!")
	}

	// ---- LOOPS ----
	//Iterate over maps
	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	//Iterate over arrays
	for index, value := range intArr3{
		fmt.Printf("Index: %v, Value: %v \n", index, value)
	}

	//No while loops in go, but you can write one like this
	var i int = 0
	for i<10{
		fmt.Println(i)
		i++
	}
	
	//You can totally omit the condition as well
	var j int = 0
	for{
		if j>10{
			break
		}
		fmt.Println(j)
		j++
	}

	//A c++ type syntax
	for k:=0; k<10; k++{
		fmt.Println(k)
	}

}