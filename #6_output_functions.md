# Output Functions

* This output function is used to print results out to the console. This functions are gotten from the format package in go which we have to import before we can be able to use it. The most widely used output functions in go are Print(), Println(), and Print(f).

1. fmt.Print(): This prints the output as being givent to it.
2. fmt.Println(): This is similar to fmt.Print(), but this returns a new line after every print.
3. fmt.Printf(): Unlike the previous two print functions, this uses special formatting verbs such as %v, %s to return the desired output.

```go 

package main
import (
    "fmt"
)

const pi float64 = 3.142 // This is a constant variable

func main() {
    fmt.Print(pi)
    fmt.Println(pi)
    fmt.Printf("%v%", pi)
}
```