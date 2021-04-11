package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	//fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

//func main() {
//	Greet(os.Stdout, "Elodie")
//}
