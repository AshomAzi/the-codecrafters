# Data Types

* This is the classifications of data depending on what they are to represent. Go gives us 4 basic data types, integers, strings, floats and booleans.

1. Integers: These are data represented as digits without decimal points.
2. Strings: These are data which can be represented as text. Strings in golang is any word enclosed by double quotations marks e.g "Hello, World".

3. FLoats: This type of data are similar to integers but they come with decimal point, e.g 2.1
4. Booleans: This represents either true or false values.

```go 

package main
import (
    "fmt"
)


func main() {
    fmt.Println("Hello, World!")    // Printing the string Hello, World!
    fmt.Println(123)                // Printing the integer 123
    fmt.Println(2.1)                // Printint the float value 2.1
    fmt.Println(true)               // Printint the boolean value true
}
```