package main

import "fmt"
import "unicode/utf8"

func main(){
	var intNum int = 3600
	fmt.Println(intNum)

	var intNum64 int64 = 4500
	fmt.Println(intNum64)

	var floatNum32 float32 = 12233.456
	fmt.Println(floatNum32)

	var floatNum64 float64 = 10.112
	var intNum16 int16 = 3000
	var result float64 = floatNum64 + float64(intNum16)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello" + " " + "World!"
	fmt.Println(myString)

	// * => len(): This doesn't tell you the number of characters in a string. It tells you the number of bytes!!
	var testString string = "test"
	fmt.Println(len(testString))  //Note: This will print 4, not to be confused. 

	//NOTE: Since Go uses utf8 encoding, characters outside the vanilla ascii character set are stored with more than a single byte.
	fmt.Println(len("A")) //Prints 1!
	fmt.Println(len("γ")) //Prints 2!

	//Gives number of characters
	fmt.Println(utf8.RuneCountInString("γ")) //This will print 1

	//Runes -> Interesting data type
	var myRune rune = 'a'
	fmt.Println(myRune) //This prints => 97

	var myBoolean bool = false
	fmt.Println(myBoolean)

	//Default values
	/* 
	int, float, rune (all the ints) => 0
	string -> ''
	bool -> false
	error -> nil
	*/

	//Other ways of declaring vars
	var myText = "text"  //You can omit the type => The type is inferred here! (Inferred as string)
	fmt.Println(myText)

	myOtherText := "boom, dropped the var keyword aswell"
	fmt.Println(myOtherText)

	//Initialise multiple vars at once
	var var1, var2 int = 2,3
	fmt.Println(var1, var2)

	var var3, var4 = 4,5
	fmt.Println(var3, var4)

	var5, var6 := 6,7
	fmt.Println(var5, var6)

	//Can't change the value of the constant once assigned. Also you have to assign a value to it, while defining
	const myConstant string = "constant"
	fmt.Println(myConstant)

	const pi float32 = 3.142
	fmt.Println(pi)
}