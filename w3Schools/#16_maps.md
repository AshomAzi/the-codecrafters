# Maps
* Maps are used to store data in the key and value pair format, in an unordered manner and does not allow duplicates
* Maps can be created just like sice but using the "map" keyword, the key types and the type of value
```go 


func main() {
  var a = map[string]int{"one": 1, "two": 2, "three": 3}
  b := map[string]int{"one": 1, "two": 2, "three": 3}

  fmt.Printf("%v\n", a)
  fmt.Printf("%v\n", b)
}
```