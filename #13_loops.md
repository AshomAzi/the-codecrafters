# Loops

* Loops helps us to run a specific block of code of a given number of time and unlike other languages, GO only makes use of for loop. The for loop is made of 3 parts, the initializations, the condition and the counter(increment or decrement)

``` go
for i:= 0; i <= 3; i++ {
    fmt.Println(i)
}
```

## Continue
* This helps us to skip a point during a for loop and move onto the next one
``` go
for i:= 0; i <= 3; i++ {
    if i == 2{
        continue
    }
    fmt.Println(i)
}
```

## Break 
* This helps us to break out of a loop when a certain condition is met

``` go
for i:= 0; i <= 3; i++ {
    if i == 2{
        break
    }
    fmt.Println(i)
}
```