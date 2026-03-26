package main

import (
	"fmt"
)


var num1 int
var num2 int
var operation string

func addition(){
	start0 :
	fmt.Print("Chose a first digit: ")
	fmt.Scanln(&num1)
	if num1 == int(num1) {
		fmt.Print("Choose a second digit: ")
		fmt.Scanln(&num2)
	} else {
		fmt.Println("Choose a second digit!")
		goto start0
	}



	// if num1 != int(num1) || num2 != int(num2){
	// 	fmt.Println("Choose only integers")
	// 	goto start0
	// } else {
	// 	fmt.Print(num1+num2)
	// }
	
}

func subtraction(){
	fmt.Print("Chose a first digit: ")
	fmt.Scan(&num1)
	fmt.Print("Choose a second digit: ")
	fmt.Scan(&num2)
	fmt.Print(num1-num2)
}

func division(){
	fmt.Print("Chose a first digit: ")
	fmt.Scan(&num1)
	start2 :
	fmt.Print("Choose a second digit: ")
	fmt.Scanln(&num2)
	if num2 == 0 {
		fmt.Println("Division by zero is not allowed!")
		goto start2
	} else {
		fmt.Println(float64(num1/num2))
	}
	
}

func multiplication(){
	fmt.Print("Chose a first digit: ")
	fmt.Scan(&num1)
	fmt.Print("Choose a second digit: ")
	fmt.Scanln(&num2)
	fmt.Println(num1*num2)
}


func main() {
	fmt.Println("Choose an operation: addition, division, multiplication or division: ")
	fmt.Println("Choose \"Help\" for suggestions! ")
	start1 :
	fmt.Scan(&operation)
	if operation == "addition" || operation == "+"{
		addition()
	} else if operation == "subtraction" || operation == "-"{
		subtraction()
	} else if operation == "multiplication" || operation == "*"{
		multiplication()
	} else if operation == "division" || operation == "/"{
		division()
	}else if operation == "Help" || operation == "help" {
		fmt.Println("Choose any of the following to choose an opperation: addition, division, multiplication or division:")
		goto start1
	} else {
		fmt.Println("Choose a valid operation!")
		goto start1
	}
}