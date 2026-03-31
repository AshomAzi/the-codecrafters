# Constants

* Constants are variables whose data should not be changed, it is decleard usimg the "const" keyword. It is used when you do not want the data of a particulare to be accidentally changed in the code, exampls when using for mathematical computations and u don't want a constant like "pi" to change.

```go 

package main

import (
    "fmt"
)

const pi float64 = 3.142 // This is a constant variable

func main() {
    fmt.Print(pi)
}

```