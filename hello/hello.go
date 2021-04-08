// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, world")
// }
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return SpanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
