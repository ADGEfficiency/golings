// generics2
// Make me compile!

package main

import "fmt"

type Number interface {
	int
}

func main() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(addNumbers(1.0, 2.3))
}

func addNumbers(n1 float32, n2 float32) float32 {
	return n1 + n2
}
