package main

import (
	"fmt"
	"strconv"
)

var value string
var operation string

func hexToDec() {
	fmt.Print("Input a value: ")
	start1:
	fmt.Scanln(&value)
	val1, err := strconv.ParseInt(value, 16, 64)
	if err == nil {
		fmt.Printf("%v = %v\n", value, val1)
	} else {
		fmt.Print("Input a valid hexadecimal value: ")
		goto start1
	}
}

func binToDec() {
	fmt.Print("Input a binary value: ")
	start1:
	fmt.Scanln(&value)
	val1, err := strconv.ParseInt(value, 2, 64)
	if err == nil {
		fmt.Printf("%v = %v\n", value, val1)
	} else {
		fmt.Print("Input a valid binary value:")
		goto start1
	}
}

func hexbinToDec() {

}

func main() {
	fmt.Println("Welcome to the Base Converter:")
	start1:
	fmt.Println("Choose an operation\n1. To convert from Hexadecimal to Binary\n2. To convert from Binary to Decimal\n3. To convert from Decimal Hexadecimal and Binary.")
	fmt.Print(">>> ")
	fmt.Scan(&operation)
	fmt.Println()

	if operation == "1" {
		hexToDec()
		goto start1
	} else if operation == "2" {
		binToDec()
		goto start1
	} else if operation == "3" {
		hexbinToDec()
	} else if operation == "" {
		fmt.Print("Empty input is not supported")
	} else {
		fmt.Println("Invalid Operation, Try Again: ")
		goto start1
	}

}
