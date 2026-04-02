# Functions
* Functions are block of programs that can be called from anywhere in a program multiple times to perform a particular action.
* A function is created using the "func" keyword followed by the fuction name, a parenthesis and a curly braces. The paranthesis are used to collect parameters which are optional while the curly braces indicate the beginning the end of the block.

``` go
func add(){}
```

* Functions can not be executed unless being called.

## Parameters and Arguments
* Parameters are informations passed into a function when creating a function while arguments are the informations passed when calling the function.

```go 
//Example of parameters and arguments

func add(value1, value2 value type) { // Parameters when creating the function
    fmt.Print(value1+value2)
}

func main() {

    add(2, 4) // Arguments when calling the function
}
```

## Return 
* The return key word is used to return a particular value type like string, integer, bool etc. It is ususlly the last thing in a function block as anyother part of after that will not be executed.

example.
``` go
func add() int {
    new := 3 + 4
    return new // Return statement
}
```

## Recursion
* A recursive function is a function that calls itselt untill it reaches a stop.

exampls 
``` go
func add(x int) int {
    if == 0 {
        return x
    }
    fmt.Print(x)
    return add(x-1)
}
```