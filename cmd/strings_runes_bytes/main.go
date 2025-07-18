package main

import (
	"fmt"
	"strings"
)

func main(){
	//é -> Non ascii character
	var myString = "résumé" 
	var indexed = myString[1]

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i,v := range myString{
		fmt.Println(i,v)
	}

	fmt.Printf("\nThe length of my string is %v", len(myString))

	// --- Runes ---
	var myRune = []rune("résumé")
	fmt.Printf("\nThe length of my string (rune) is %v\n", len(myRune))

	for i,v := range myRune{
		fmt.Println(i,v)
	}

	//NOTE: You can declare a run type like this as well
	var anotherRune = 'a'
	fmt.Println(anotherRune)

	//NOTE: Strings are immutable
	//String Concatenation (2 ways -> Inefficient and efficient)
	//1. Inefficient (Because it makes a new string every time you concatenate)
	var strSlice = []string{"h","e", "l", "l", "o","g","o"}
	var catStr = ""

	for i := range strSlice{
		catStr += strSlice[i]
	}

	fmt.Printf("\n%v", catStr)

	//2. Using string builder -> Better
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}

	var catStr2 = strBuilder.String()
	fmt.Println(catStr2)

}