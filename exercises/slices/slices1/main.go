// slices1
// Make me compile!

package main

import "fmt"

func main() {
	a := make([]float32, 4, 4) // play with length and capacity
	fmt.Println("length of 'a':", len(a))
	fmt.Println("capacity of 'a':", cap(a))
	fmt.Println(a)
}
