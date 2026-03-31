# Arrays

* Arrays are used to store multiple values of same type into a single variable. Arrays have a defined fixed size and length and it cannot be changed. Arrays can be declared in two ways
1. Using the var keyword 

```go
var users = [3]string{"john", "victor", "Philip"}
```

2. Using the ":=" sign

```go
users := [3]string{"john", "victor", "Philip"}
```

## Accessing Elements of an array

* This can be done using the index value. Indexing in GO starts from 0, therefors to access the first element you would use the number 0

```go
fmt.Println(users[0]) // This accesses the first value in the array.
```

## Change Elements of an Array

* This can be achieved by re-assigning a new value to the index of that array, while ensuring that they are of same type.

```go
users[0] = "Peter"
```

## Initializing an array

* When initializing an array, when values are not assigned or partially assigned, go explicitly assigns the rest of the value with and empty string for strings ans zero for integers.

```go
num1 := [10]int{1, 2, 3, 4}
```
