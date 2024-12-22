```go
package main

import (


    "fmt"
)

func main() {
    var m map[string]int
    m = make(map[string]int)
    value, ok := m["a"]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```