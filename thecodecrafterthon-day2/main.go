package main

import (
	"fmt"
	"strconv"
	"strings"
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
	fmt.Print("Input a valid decimal value: ")
	start1:
	fmt.Scan(&value)
	val1, err := strconv.Atoi(value)
	if err == nil {
		res1 := strconv.FormatInt(int64(val1), 2)
		res2 := strconv.FormatInt(int64(val1), 16)
		fmt.Printf("Binary value of %v is %v\nHexadecimal value of %v is %v\n", value, res1, value, strings.ToUpper(res2))
	} else {
		fmt.Print("Input a valid decimal value: ")
		goto start1
	}
}

func main() {
	fmt.Println("Welcome to the Base Converter:")
start1:
	fmt.Println("Choose an operation\n1. To convert from Hexadecimal to Binary\n2. To convert from Binary to Decimal\n3. To convert from Decimal Hexadecimal and Binary.")
	fmt.Print(">>> ")
	fmt.Scan(&operation)

	if operation == "1" {
		hexToDec()
		goto start1
	} else if operation == "2" {
		binToDec()
		goto start1
	} else if operation == "3" {
		hexbinToDec()
		goto start1
	// } else if operation  {
	// 	fmt.Print("Empty input is not supported")
	} else {
		fmt.Print("Invalid Operation, Try Again: ")
		goto start1
	}

}
