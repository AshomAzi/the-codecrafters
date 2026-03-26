# Switch

Switch works similar to the  if statement. The value of a switch statement is compared to the case value and when it matches, it executes the block of code abd if non mathches, it executes the default case, and note, the default case is optional.

```go
func main() {
	day := 1

	switch day {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Value doesn't match")
  }
}
```