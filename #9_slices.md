# Slices

* Slices are similar to arrays but are flexible, meaning that their values can be added or reduced. Sliced doesn't have an explicit size and this is what makes it flexible and powerful, the capacity and length of a slice depends on the total number of data in the slice. 

## Declaring Slices
* This can be done using the var keyword

```go 
var users = []string{"John", "Peter"}
```

* Or by using the ":=" sign

```go
users := []string{"John", "Peter"}
```

## Create a Slice From an Array

This can be achieved by slicing an array

```go
num1 := [10]int{1, 2, 3, 4} // Array
num2 := num1[1:3]           // A slice created from an array.
```