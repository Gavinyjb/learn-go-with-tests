// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, world")
// }
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case "french":
		prefix = FrenchHelloPrefix
	case "spanish":
		prefix = SpanishHelloPrefix
	}
	return prefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
	fmt.Println(Hello("world", "french"))
	fmt.Println(Hello("world", "spanish"))
}
