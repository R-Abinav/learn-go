package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type example struct{
	d1 uint8
	d2 string
	gasEngine
}

type example2 struct{
	d1 uint8
	d2 string
	gasEngineInfo gasEngine
}

func (e gasEngine) milesLeft() uint8{
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8{
	return e.mpkwh * e.kwh
}

func canMakeItOrNot(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("\nCan make it!")
	}else{
		fmt.Println("\nCannot make it!")
	}
}
type engine interface{
	milesLeft() uint8
}

func main(){
	var varStruct example = example{d1: 1, d2: "Hello", gasEngine: gasEngine{mpg: 12}}
	fmt.Printf("\n%v", varStruct)
	
	var varStruct2 example2 = example2{d1: 1, d2: "Hello", gasEngineInfo: gasEngine{mpg: 12}}
	fmt.Printf("\n%v", varStruct2)

	var myEngine gasEngine = gasEngine{mpg: 2, gallons: 3}
	fmt.Printf("\nTotal miles left in tank: %v", myEngine.milesLeft())

	var myEngine2 electricEngine = electricEngine{mpkwh: 5, kwh: 3}
	canMakeItOrNot(myEngine2, 20)
}