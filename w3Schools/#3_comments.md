# Comments

* A comment is a text that is ignored or overlooked during execution, it can also be used to give information on what the code is all about.

## There are 2 ways of writing a comment in go. 

1. Single-line comment
>* This can be achieved by using the double forward slash '//'. It is used to comment a single line.

2. Multi-line comment 
>* This can be achieved by using the forward slash '/' and the  asterisk symbol '*'. It is used to comment multiple lines at the same time.


```go 

package main // This is a single line comment

import ( 
    "fmt"
)

/*
This is
a multi-line
comment
*/

func main() {
    fmt.Println("Hello, World!")
}

```


