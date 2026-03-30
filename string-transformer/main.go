// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Azi Bulus]
// Squad:  [Pointers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var command string

func upper() {
	fmt.Print("Enter a text: ")
	line := bufio.NewScanner(os.Stdin)
	if line.Scan() {
		line := line.Text()
		if len(line) > 0 {
			fmt.Println(strings.ToUpper(line))
		} else {
			fmt.Println("Does not accept an empty input")
		}
	} else {
		fmt.Println("An error occured while reading input!")
	}
}

func lower() {
	fmt.Print("Enter a text: ")
	line := bufio.NewScanner(os.Stdin)
	if line.Scan() {
		line := line.Text()
		if len(line) > 0 {
			fmt.Println(strings.ToLower(line))
		} else {
			fmt.Println("Doesn't accept an empty input")
		}
	} else {
		fmt.Println("An error occured while reading input")
	}
}

func caps() {
	fmt.Print("Enter a text: ")
start1:
	line := bufio.NewScanner(os.Stdin)
	if line.Scan() {
		new := line.Text()
		nval := strings.Fields(new)
		if len(nval) > 0 {
			for i := 0; i <= len(nval)-1; i++ {
				low := strings.ToLower(nval[i])
				title := strings.ToUpper(string(low[0])) + string(low[1:])
				// title := strings.Title(low)
				fmt.Print(title, " ")
			}
			fmt.Println(" ")
		} else {
			fmt.Println("Doesn't accept an empty input")
			goto start1
		}
	} else {
		fmt.Println("Error Input! ")
	}
}

func title() {
	var small = []string{"a", "an", "the", "and", "but",
		"or", "for", "nor", "on", "at", "to", "by", "in",
		"of", "up", "as", "is", "it"}
	fmt.Print("Input a text: ")
	val := bufio.NewScanner(os.Stdin)
	if val.Scan() {
		line := val.Text()
		newLine := strings.Fields(line)
		for i := 0; i <= len(small)-1; i++ {
			for j := 0; j <= len(newLine)-1; j++ {
				if small[i] == newLine[j] {
					fmt.Println(small[j])
				}
				fmt.Print(small[i], " ")
			}
		}
		fmt.Println(" ")
	}

}

func snake_case() {

}

func reverse() {

}

func main() {
	title()
	// lower()
}
