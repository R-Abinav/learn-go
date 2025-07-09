package main

import (
	"fmt"
	"errors"
)

//Note -> You have to put the starting braces '{' in the starting line, can't put it below or anywhere else
func main(){
	var getPrinted string = "Send as param to be printed"
	printSomething(getPrinted)

	var numerator int = 10
	var denominator int = 5
	var result, remainder, err = intDivision(numerator, denominator)

	//if else block:

	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result of integer division is %v", result)
	// }else{
	// 	fmt.Printf("The result of integer division is %v and its remainder is %v", result, remainder)
	// }

	//Switch statements
	//NOTE -> You don't have to write "break" after every case in GO. Here break is instilled in!

	switch{
		case err!=nil:
			fmt.Println(err.Error())
		case remainder==0:
			fmt.Printf("The result of integer division is %v", result)
		default:
			fmt.Printf("The result of integer division is %v and its remainder is %v", result, remainder)
	}

	// switch remainder{
	// 	case 0: 
	// 		fmt.Println("The division was exact")
	// 	case 1,2:
	// 		fmt.Println("The division was close")
	// 	default:
	// 		fmt.Println("The division was not close")
	// }
}

func intDivision(num1 int, num2 int) (int, int, error) {
	var err error  //Default value of error is nil
	if num2==0{
		err = errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result int = num1/num2
	var remainder int = num1%num2
	return result, remainder, err
}

func printSomething(getPrinted string){
	fmt.Println(getPrinted)
}