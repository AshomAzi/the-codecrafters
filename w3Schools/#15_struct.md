# Structs

* Structs are used to create a collection of  of different types into a single variable. To declare a struct, we use the keyword "type" followed by the name of the struct and the "struct" keyword

``` go 
type Person struct {
  name string
  age int
  job string
  salary int
}
```

## Accessing the members of a struct

This can be done by use the struct name followed by the "." operator and the struct member.