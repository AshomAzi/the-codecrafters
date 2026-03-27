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
	fmt.Print("Chose the first digit: ")
start0:
	fmt.Scanln(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Input the second digit: ")
	start1:
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v + %v = %v\n", val1, val2, val1+val2)
		} else {
			fmt.Print("Select a valid digit: ")
			goto start1
		}
	} else {
		fmt.Print("Select a valid digit: ")
		goto start0
	}

}

func subtraction() {
	fmt.Print("Chose the first digit: ")
start0:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Input the second digit: ")
	start1:
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v - %v = %v\n", val1, val2, val1-val2)
		} else {
			fmt.Print("Input a valid digit: ")
			goto start1
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start0
	}
}

func division() {
	fmt.Print("Chose a first digit: ")
start0:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Chose a second digit: ")
	start1:
		fmt.Scan(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil && val2 != 0 {
			fmt.Printf("%v / %v = %v\n", val1, val2, val1/val2)
		} else {
			fmt.Print("Input a valid digit, Zero is not allowed: ")
			goto start1
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start0
	}

}

func multiplication() {
	fmt.Print("Chose a first digit: ")
start1:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Chose a second digit: ")
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
	start1:
		fmt.Scan(&operation)
		if operation == 1 {
			addition()
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
			goto start1
		} else if operation == 6 {
			fmt.Println("Goodbye Codecrafter!")
			break
		} else {
			fmt.Println("Choose a valid operation!")
			goto start1
		}
	}
}
