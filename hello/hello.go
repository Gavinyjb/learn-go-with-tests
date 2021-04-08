// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, world")
// }
package main

import "fmt"

func Hello(name string) string {
    return "Hello, "+name
}

func main() {
    fmt.Println(Hello("world"))
}