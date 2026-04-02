// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Azi Bulus]
// Squad:  [Pointers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var smallWords = []string{"a", "an", "the", "and", "but",
		"or", "for", "nor", "on", "at", "to", "by",
		"in", "of", "up", "as", "is", "it"}
	fmt.Print("Input a text: ")
	text := bufio.NewScanner(os.Stdin)
	if text.Scan() {
		line := text.Text()
		words := strings.Fields(line)

		for i, word := range words {
			isSmall := slices.Contains(smallWords, strings.ToLower(word))
			if i == 0 || !isSmall {
				words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			} else {
				words[i] = strings.ToLower(word)
			}
		}
		fmt.Println(strings.Join(words, " "))
	}
}

func snake_case() {
	fmt.Print("Input a text: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)
		result := ""
		for i := 0; i < len(line); i++ {
			letter := line[i] >= 'a' && line[i] <= 'z'
			digit := line[i] >= '0' && line[i] <= '9'
			space := line[i] == ' '
			if letter || digit {
				result += string(line[i])
			} else if space {
				result += "_"
			}
		}
		for strings.Contains(result, "__") {
			result = strings.ReplaceAll(result, "__", "_")
		}
		result = strings.Trim(result, "_")
		fmt.Println(result)
	}
}

func reverse() {

}

func main() {
	snake_case()
	// lower()
}
