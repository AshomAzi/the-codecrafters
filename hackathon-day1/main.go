package main

import (
	"fmt"
	"strconv"
)

var num1 string
var num2 string
var operation string
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
	start1:
		fmt.Scanln(&operation)
		opps, err := strconv.Atoi(operation)
		if err == nil {
			if opps == 1 {
				addition()
				goto start0
			} else if opps < 1 || opps > 6 {
				fmt.Print("Invalid Operation\n\n")
				goto start0
			} else if opps == 2 {
				subtraction()
				goto start0
			} else if opps == 3 {
				multiplication()
				goto start0
			} else if opps == 4 {
				division()
				goto start0
			} else if opps == 5 {
				fmt.Print("Choose any of the following for an operation:\n1. To perform an addition\n2. To perform a subtraction\n3. To perform a multiplication\n4. To perform a division\n5. Help\n6. Exit\n")
				goto start1
			} else if opps == 6 {
				fmt.Println("Goodbye Codecrafter!")
				break
			}
		} else {
			fmt.Print("Input a valid operation, or choose \"Help\" for assistance: ")
			goto start1
		}

	}
}
