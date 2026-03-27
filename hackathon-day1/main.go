package main

import (
	"fmt"
	"strconv"
)

var num1 string
var num2 string
var operation int
var exit int

func addition() {

	fmt.Print("Choose the first digit: ") // Here
start:
	fmt.Scanln(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Input the second digit: ")
	start2:
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v + %v = %v\n\n", val1, val2, val1+val2)
		} else {
			fmt.Println("Select a valid digit!")
			goto start2
		}
	} else {
		fmt.Println("Invalid input, Choose a digit!")
		goto start
	}

}

func subtraction() {
	fmt.Print("Choose the first digit: ")
start1:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Input the second digit: ")
	start2:
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v - %v = %v\n", val1, val2, val1-val2)
		} else {
			fmt.Print("Input a valid digit: ")
			goto start2
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start1
	}
}

func division() {
	fmt.Print("Choose a first digit: ")
start:
	fmt.Scan(&num1)
	val1, err := strconv.ParseFloat(num1, 64)
	if err == nil {
		fmt.Print("Choose a second digit: ")
	start1:
		fmt.Scan(&num2)
		val2, err := strconv.ParseFloat(num2, 64)
		if err == nil && val2 != 0 {
			fmt.Printf("%v / %v = %v\n", val1, val2, val1/val2)
		} else {
			fmt.Println("Input a valid digit, Zero is not allowed: ")
			goto start1
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start
	}

}

func multiplication() {
	fmt.Print("Choose a first digit: ")
start1:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Choose a second digit: ")
	start2:
		fmt.Scan(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v * %v = %v\n", val1, val2, val1*val2)
		} else {
			fmt.Print("Input a valid digit: ")
			goto start2
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start1
	}
}

func main() {
start0:
	fmt.Println("Choose an operation:\n1. addition\n2. subtraction\n3. multiplication\n4. division\n5. Help\n6. Exit")
	fmt.Println("Choose \"Help\" for suggestions! ")
	for i := 0; i <= 6; i++ {
		fmt.Scan(&operation)
		if operation == 1 {
			addition()
			goto start0
		} else if operation < 1 || operation > 6 {
			fmt.Print("Invalid Operation\n\n")
			goto start0
		} else if operation == 2 {
			subtraction()
			goto start0
		} else if operation == 3 {
			multiplication()
			goto start0
		} else if operation == 4 {
			division()
			goto start0
		} else if operation == 5 {
			fmt.Println("Choose any of the following for an opperation:\n1. addition\n2. subtraction\n3. multiplication\n4. division\n5. Help\n6. Exit")
			goto start0
		} else if operation == 6 {
			fmt.Println("Goodbye Codecrafter!")
			break
		}
	}
}
